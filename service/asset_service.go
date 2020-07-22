package service

import (
	"claps-test/dao"
	"claps-test/model"
	"claps-test/util"
)

func GetAssetById(assetID string) (asset *model.Asset,err *util.Err){
	asset,err1 := dao.GetAssetById(assetID)
	if err1 != nil {
		err = util.NewErr(err1,util.ErrDataBase,"通过id获取asset信息失败")
	}
	return
}

func ListAssetsAllByDB()(assets *[]model.Asset,err *util.Err){
	assets,err1 := dao.ListAssetsAllByDB()
	if err1 != nil {
		err = util.NewErr(err1,util.ErrDataBase,"获取全部assets信息失败")
	}
	return

}
