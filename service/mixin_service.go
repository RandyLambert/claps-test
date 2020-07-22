package service

import (
	"claps-test/dao"
	"claps-test/model"
	"claps-test/util"
	"context"
	"github.com/fox-one/mixin-sdk-go"
	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"time"
)

//获取所有币的信息
func ListAssets()(assets []*mixin.Asset,err *util.Err){

	assets,err1 := util.MixinClient.ReadAssets(context.Background())
	if err1 != nil {
		err = util.NewErr(err1,util.ErrThirdParty,"获取全部asset信息错误")
	}
	return
}

func GetAssetById(assetID string) (asset *mixin.Asset,err *util.Err){
	asset,err1 := util.MixinClient.ReadAsset(context.Background(),assetID)
	if err1 != nil {
		return nil,util.NewErr(err1,util.ErrThirdParty,"获取单个asset信息错误")
	}

	if assetID != asset.AssetID {
		return nil,util.NewErr(err1,util.ErrThirdParty,"获取信息Id不一致")
	}
	return
}

func CreateMixinClient(botId string)(client *mixin.Client,err *util.Err){
	bot,err1 := dao.GetBotById(botId)
	if err1 != nil {
		err = util.NewErr(err1,util.ErrDataBase,"获取bot信息失败")
		return
	}
	s := &mixin.Keystore{
		ClientID:   bot.Id,
		SessionID:  bot.SessionId,
		PrivateKey: bot.PrivateKey,
		PinToken: bot.PinToken,
	}

	client, err1 = mixin.NewFromKeystore(s)
	if err1 != nil {
		err = util.NewErr(err1,util.ErrThirdParty,"创建mixinclient失败")
	}
	return
}


func SyncSnapshots() {

	ctx := context.TODO()

	//获取当前时间
	since := time.Now()

	//死循环,读到上次最后一条就break
	for {
		//获取最后一次跟新记录
		property,_ := dao.GetPropertyByKey("last_snapshot_id")
		lastSnapshotID := property.Value

		//从mixin获取当前时间之前的
		snapshots, err := util.MixinClient.ReadNetworkSnapshots(ctx, "", since, "ASC", 100)
		//这个时间端没有交易记录
		if len(snapshots) == 0 {
			time.Sleep(time.Second)
			continue
		}

		//错误处理
		if err != nil {
			log.Error(err.Error())
			continue
		}

		//遍历100记录
		for i,_ := range snapshots {
			/*
			log.Debug(*snapshots[i])
			log.Debug("\n")
			 */
			if lastSnapshotID == snapshots[i].SnapshotID {
				continue
			}

			//筛选自己的转入
			if snapshots[i].UserID != "" && snapshots[i].Amount.Cmp(decimal.Zero) > 0 {
				//根据机器人从数据库里找到项目
				project,err := dao.GetProjectByBotId(snapshots[i].OpponentID)
				//错误处理有问题
				if err != nil {
					log.Error(err.Error())
				}

				transaction := &model.Transaction{
					Id:        snapshots[i].SnapshotID,
					ProjectId: project.Id,
					BotId:     snapshots[i].OpponentID,
					AssetId:   snapshots[i].AssetID,
					Amount:    snapshots[i].Amount,
					CreatedAt: snapshots[i].CreatedAt,
					Sender:    snapshots[i].UserID,
					Receiver:  snapshots[i].OpponentID,
				}
				//插入捐赠记录
				err = dao.InsertTransaction(transaction)
				if err != nil {
					log.Error(err.Error())
				}

				//查找汇率等详细信息,目前还是及时获取的
				//可以建一张表,来缓存AssetId信息
				asset,err := GetAssetById(snapshots[i].AssetID)
				if err != nil {
					log.Error(err.Error())
				}

				//更新Total字段
				project.Total = project.Total.Add(asset.PriceUSD.Mul(snapshots[i].Amount))
				count,err := dao.CountPatronByProjectIdAndSender(project.Id,snapshots[i].UserID)
				if count == 0 {
					project.Patrons += 1
				}

				err = dao.UpdateProject(project)
				if err != nil {
					log.Error(err.Error())
				}

				//更新项目钱包

				//根据不同的分配算法进行配置
			}
		}

		lastSnapshotID = snapshots[len(snapshots)-1].SnapshotID
		property = &model.Property{
			Key:   "last_snapshot_id",
			Value: lastSnapshotID,
		}
		err = dao.UpdateProperty(property)
		if err != nil {
			log.Error(err.Error())
		}
		since = snapshots[len(snapshots)-1].CreatedAt
		time.Sleep(100 * time.Millisecond)
	}
}

func DoTransfer(ctx context.Context, botId, assetID, opponentID string, amount decimal.Decimal, memo, pin string,userId uint32) (err *util.Err) {


	//opponentid是转给谁
	// Transfer transfer to account
	//	asset_id, opponent_id, amount, traceID, memo
	// 把该user的钱转账到该账户返回快照
	user,err := CreateMixinClient(botId)
	if err != nil {
		return
	}
	//traceid暂时不应该这样
	snapshot, err1 := user.Transfer(ctx, &mixin.TransferInput{
		TraceID:    uuid.Must(uuid.NewV4()).String(),
		AssetID:    assetID,
		OpponentID: opponentID,
		Amount:     amount,
		Memo:       memo,
	}, pin)

	if err1 != nil {
		err = util.NewErr(err,util.ErrThirdParty,"转账失败")
		return
	}

	//这里的memberwallet,是通过外部获取的,业务逻辑不是这样,暂时这么写,这里的wallet里面的balance功能存疑
	memberWallet,err1 := dao.GetMemberWalletByProjectIdAndUserIdAndBotIdAndAssetId(1,1,botId,assetID)
	if err1 != nil {
		err = util.NewErr(err,util.ErrDataBase,"获取用户钱包失败")
		return
	}
	memberWallet.Balance = memberWallet.Balance.Sub(amount)

	err1 = dao.UpdateMemberWallet(memberWallet)
	if err1 != nil {
		err = util.NewErr(err,util.ErrDataBase,"更新用户钱包失败")
		return
	}

	wallet,err1 := dao.GetWalletByBotIdAndAssetId(botId,assetID)
	if err1 != nil {
		err = util.NewErr(err,util.ErrDataBase,"获取项目钱包失败")
		return
	}

	wallet.Balance = wallet.Balance.Sub(amount)

	err1 = dao.UpdateWallet(wallet)
	if err1 != nil {
		err = util.NewErr(err,util.ErrDataBase,"更新项目钱包失败")
		return
	}

	transfer := &model.Transfer{
		BotId:      user.ClientID,
		SnapshotId: snapshot.SnapshotID,
		UserId:     userId,
		TraceId:    snapshot.TraceID,
		OpponentId: opponentID,
		AssetId:    assetID,
		Amount:     decimal.Decimal{},
		Memo:       memo,
		CreatedAt:  snapshot.CreatedAt,
	}

	err1 = dao.InsertTransfer(transfer)
	if err1 != nil {
		err = util.NewErr(err,util.ErrDataBase,"交易记录插入数据库失败")
	}
	return
}
