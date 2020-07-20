package service

import (
	"claps-test/dao"
	"claps-test/model"
	"claps-test/util"
	"context"
	"errors"
	"github.com/fox-one/mixin-sdk-go"
	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"time"
)

func ListAssets()(assets []*mixin.Asset,err error){

	assets,err = util.MixinClient.ReadAssets(context.Background())
	return
}

func GetAssetById(assetID string) (asset *mixin.Asset,err error){
	asset,err =  util.MixinClient.ReadAsset(context.Background(),assetID)
	if asset != nil{
		if assetID != asset.AssetID {
			log.Error("asset should be %s but get %s\n", assetID, asset.AssetID)
			return nil,errors.New("GetAsset error")
		}
	}
	return
}

func CreateMixinClient(botId string)(client *mixin.Client,err error){
	bot,err := dao.GetBotById(botId)
	if err != nil {
		return
	}
	s := &mixin.Keystore{
		ClientID:   bot.Id,
		SessionID:  bot.SessionId,
		PrivateKey: bot.PrivateKey,
		PinToken: bot.PinToken,
	}

	client, err = mixin.NewFromKeystore(s)
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

func DoTransfer(ctx context.Context, botId, assetID, opponentID string, amount decimal.Decimal, memo, pin string,userId uint32) (err error) {


	//opponentid是转给谁
	// Transfer transfer to account
	//	asset_id, opponent_id, amount, traceID, memo
	// 把该user的钱转账到该账户返回快照
	user,err := CreateMixinClient(botId)
	if err != nil {
		return
	}
	//traceid暂时不应该这样
	snapshot, err := user.Transfer(ctx, &mixin.TransferInput{
		TraceID:    uuid.Must(uuid.NewV4()).String(),
		AssetID:    assetID,
		OpponentID: opponentID,
		Amount:     amount,
		Memo:       memo,
	}, pin)

	if err != nil {
		return
	}

	//这里的memberwallet,是通过外部获取的,业务逻辑不是这样,暂时这么写,这里的wallet里面的balance功能存疑
	memberWallet,err := dao.GetMemberWalletByProjectIdAndUserIdAndBotIdAndAssetId(1,1,botId,assetID)
	if err != nil {
		return
	}
	memberWallet.Balance = memberWallet.Balance.Sub(amount)

	err = dao.UpdateMemberWallet(memberWallet)
	if err != nil {
		return
	}

	wallet,err := dao.GetWalletByBotIdAndAssetId(botId,assetID)
	if err != nil {
		return
	}

	wallet.Balance = wallet.Balance.Sub(amount)

	err = dao.UpdateWallet(wallet)
	if err != nil {
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
	err = dao.InsertTransfer(transfer)

	return
}
