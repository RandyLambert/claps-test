package service

import (
	"claps-test/dao"
	"claps-test/model"
	"claps-test/util"
	"context"
	"github.com/fox-one/mixin-sdk-go"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

//获取所有币的信息
func ListAssetsAllByMixinClient(client *mixin.Client)(assets []*mixin.Asset,err *util.Err){

	assets,err1 := client.ReadAssets(context.Background())
	if err1 != nil {
		err = util.NewErr(err1,util.ErrThirdParty,"通过相应bot获取mixin全部asset信息错误")
	}
	return

}

func CreateMixinClient(bot *model.Bot)(client *mixin.Client,err *util.Err){

	s := &mixin.Keystore{
		ClientID:   bot.Id,
		SessionID:  bot.SessionId,
		PrivateKey: bot.PrivateKey,
		PinToken: bot.PinToken,
	}

	client, err1 := mixin.NewFromKeystore(s)
	if err1 != nil {
		err = util.NewErr(err1,util.ErrThirdParty,"创建mixinclient失败")
	}
	return
}

func SyncAssets() {

	ctx := context.TODO()
	for {
		assetsInfo,err := util.MixinClient.ReadAssets(ctx)
		//错误处理
		if err != nil {
			log.Error(err.Error())
			continue
		}

		for i := range assetsInfo {
			//log.Debug(assetsInfo[i].Name)
			if assetsInfo[i].AssetID == util.BTC ||
				assetsInfo[i].AssetID == util.BCH ||
				assetsInfo[i].AssetID == util.ETC ||
				assetsInfo[i].AssetID == util.EOS ||
				assetsInfo[i].AssetID == util.XRP ||
				assetsInfo[i].AssetID == util.XEM ||
				assetsInfo[i].AssetID == util.USDT {

				asset := &model.Asset{
					AssetId:  assetsInfo[i].AssetID,
					Symbol:   assetsInfo[i].Symbol,
					Name:     assetsInfo[i].Name,
					IconUrl:  assetsInfo[i].IconURL,
					PriceBtc: assetsInfo[i].PriceBTC,
					PriceUsd: assetsInfo[i].PriceUSD,
				}
				//第一次使用前,如果数据库没有信息,需要先创建几条记录,之后使用就每次更新即可
				//err = dao.InsertAsset(asset)
				err = dao.UpdateAsset(asset)

				if err != nil {
					log.Error(err.Error())
				}
			}
		}
		time.Sleep(time.Minute*5)
	}
}

func SyncTransfer() {

	ctx := context.TODO()
	for {
		transfers,err := dao.ListTransfersByStatus('1')
		if err != nil {
			log.Error(err.Error())
			continue
		}

		//说明当前时间没有提现记录
		if len(*transfers) == 0 {
			time.Sleep(1*time.Second)
			continue
		}

		for i := range *transfers {
			//opponentid是转给谁
			// Transfer transfer to account
			//	asset_id, opponent_id, amount, traceID, memo
			// 把该user的钱转账到该账户返回快照
			bot,err := dao.GetBotById((*transfers)[i].BotId)
			if err != nil {
				log.Error(err.Error())
				continue
			}

			user,err := CreateMixinClient(bot)
			if err != nil {
				log.Error(err.Error())
				continue
			}
			//traceid暂时不应该这样
			snapshot, err := user.Transfer(ctx, &mixin.TransferInput{
				TraceID:    uuid.Must(uuid.NewV4()).String(),
				AssetID:    (*transfers)[i].AssetId,
				OpponentID: (*transfers)[i].BotId,
				Amount:     (*transfers)[i].Amount,
				Memo:       (*transfers)[i].Memo,
			}, bot.Pin)

			if err != nil {
				log.Error(err.Error())
				continue
			}

			(*transfers)[i].Status = '1'
			(*transfers)[i].TraceId = snapshot.TraceID
			(*transfers)[i].SnapshotId = snapshot.SnapshotID
			(*transfers)[i].CreatedAt = snapshot.CreatedAt

			err = dao.UpdateTransfer(&(*transfers)[i])
			if err != nil {
				log.Error(err.Error())
			}
		}

		time.Sleep(300*time.Millisecond)
	}
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

		//从mixin获取当前时间之后的snapshots
		snapshots, err := util.MixinClient.ReadNetworkSnapshots(ctx, "", since, "ASC", 100)

		//错误处理
		if err != nil {
			log.Error(err.Error())
			continue
		}

		//这个时间端没有交易记录
		if len(snapshots) == 0 {
			time.Sleep(time.Second)
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
				projectTotal,err := dao.GetProjectTotalByBotId(snapshots[i].OpponentID)
				//错误处理有问题
				if err != nil {
					log.Error(err.Error())
					continue
				}

				transaction := &model.Transaction{
					Id:        snapshots[i].SnapshotID,
					ProjectId: projectTotal.Id,
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
					continue
				}

				//查找汇率等详细信息
				asset,err := dao.GetPriceUsdByAssetId(snapshots[i].AssetID)
				if err != nil {
					log.Error(err.Error())
					continue
				}

				//更新Total字段
				projectTotal.Total = projectTotal.Total.Add(asset.PriceUsd.Mul(snapshots[i].Amount))
				count,err := dao.CountPatronByProjectIdAndSender(projectTotal.Id,snapshots[i].UserID)
				if count == 0 {
					projectTotal.Patrons += 1
				}

				err = dao.UpdateProjectTotal(projectTotal)
				if err != nil {
					log.Error(err.Error())
					continue
				}

				//更新项目钱包
				walletTotal,err := dao.GetWalletTotalByBotIdAndAssetId(snapshots[i].OpponentID,snapshots[i].AssetID)
				if err != nil {
					log.Error(err.Error())
					continue
				}
				walletTotal.Total = walletTotal.Total.Add(snapshots[i].Amount)
				err = dao.UpdateWalletTotal(walletTotal)
				if err != nil {
					log.Error(err.Error())
					continue
				}
				//根据不同的分配算法进行配置
				bot,err := dao.GetBotDtoById(snapshots[i].OpponentID)

				switch bot.Distribution {
				case model.PersperAlgorithm:
					go distributionByPersperAlgorithm(transaction)
				case model.Commits:
					go distributionByCommits(transaction)
				case model.ChangedLines:
					go distributionByChangedLines(transaction)
				case model.IdenticalAmount:
					go distributionByIdenticalAmount(transaction)
				}
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

//获取认证之后的客户端
func GetMixinAuthorizedClient(ctx *gin.Context,code string)(client *mixin.Client,err *util.Err) {
	//从配置文件中读取Id和密码
	clientId := viper.GetString("MIXIN_CLIENT_ID")
	clientSecret := viper.GetString("MIXIN_CLIENT_SECRET")

	//生成Key
	key := mixin.GenerateEd25519Key()

	//code换Token
	store,err1 := mixin.AuthorizeEd25519(ctx,clientId,clientSecret,code,"",key)
	if err1 != nil {
		err = util.NewErr(err1,util.ErrThirdParty,"mixin Ed25519出错")
		return
	}

	//换取
	client, err2 := mixin.NewFromOauthKeystore(store)
	if err2 != nil {
		err = util.NewErr(err2,util.ErrThirdParty,"mixin store to client error")
		return
	}

	/*
	//将client存入session
	session := sessions.Default(ctx)
	session.Set("mixinClient",client)
	err3 := session.Save()
	if err3 != nil{
		err = util.NewErr(err3,util.ErrInternalServer,"设置mixin client session 出错")
		return
	}
	 */

	return
}

func GetMixinUserInfo(ctx *gin.Context,client *mixin.Client) (user *mixin.User, err *util.Err) {

	user, err1 := client.UserMe(ctx)
	if err1 != nil {
		err = util.NewErr(err,util.ErrDataBase,"获取mixin用户的信息出错")
		return
	}
	return
}


/*
func DoTransfer(botId, assetID ,opponentID, memo string, amount decimal.Decimal, userId uint32) (err *util.Err) {

	transfer := &model.Transfer{
		BotId:      botId,
		UserId:     userId,
		TraceId:    string(userId)+assetID+botId,
		//OpponentId: opponentID,
		AssetId:    assetID,
		Amount:     amount,
		Memo:       memo,
		Status:    '0',
	}

	err1 := dao.InsertTransfer(transfer)

	if err1 != nil {
		err = util.NewErr(err,util.ErrDataBase,"提现记录首次写入数据库失败导致提现失败,确认当时是否有一笔提现记录未被确认?")
		return
	}

	//这里的memberwallet,是通过外部获取的,业务逻辑不是这样,暂时这么写
	memberWallet,err1 := dao.GetMemberWalletByProjectIdAndUserIdAndBotIdAndAssetId(1,1,botId,assetID)
	if err1 != nil {
		err = util.NewErr(err,util.ErrDataBase,"获取用户钱包失败导致提现失败")
		return
	}

	memberWallet.Balance = memberWallet.Balance.Sub(amount)

	err1 = dao.UpdateMemberWallet(memberWallet)
	if err1 != nil {
		err = util.NewErr(err,util.ErrDataBase,"更新用户钱包可提现值导致提现失败")
	}

	return
}
 */
