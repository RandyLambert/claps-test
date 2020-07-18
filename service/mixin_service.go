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
	since := time.Now()
	for {
		property,_ := dao.GetPropertyByKey("last_snapshot_id")
		lastSnapshotID := property.Value
		snapshots, err := util.MixinClient.ReadNetworkSnapshots(ctx, "", since, "ASC", 100)
		if len(snapshots) == 0 {
			time.Sleep(time.Second)
			continue
		}

		if err != nil {
			log.Error(err.Error())
			continue
		}

		for i,_ := range snapshots {
			if lastSnapshotID == snapshots[i].SnapshotID {
				continue
			}

			if snapshots[i].UserID != "" && snapshots[i].Amount.Cmp(decimal.Zero) > 0 {
				project,err := dao.GetProjectByBotId(snapshots[i].OpponentID)
				if err != nil {
					log.Error(err.Error())
				}

				transaction := &model.Transaction{
					Id:        snapshots[i].SnapshotID,
					ProjectId: project.Id,
					BotId:     snapshots[i].UserID,
					AssetId:   snapshots[i].AssetID,
					Amount:    snapshots[i].Amount,
					CreatedAt: snapshots[i].CreatedAt,
					Sender:    snapshots[i].UserID,
					Receiver:  snapshots[i].OpponentID,
				}
				err = dao.InsertTransaction(transaction)
				if err != nil {
					log.Error(err.Error())
				}

				asset,err := GetAssetById(snapshots[i].AssetID)
				if err != nil {
					log.Error(err.Error())
				}
				project.Total = project.Total.Add(asset.PriceUSD.Mul(snapshots[i].Amount))
				count,err := dao.CountPatronByProjectIdAndSender(project.Id,snapshots[i].UserID)
				if count == 0 {
					project.Patrons += 1
				}

				err = dao.UpdateProject(project)
				if err != nil {
					log.Error(err.Error())
				}

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
