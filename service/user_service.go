package service

import (
	"claps-test/dao"
	"claps-test/model"
	"claps-test/util"
	"context"
	"github.com/google/go-github/v32/github"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

//从gitub服务器请求获取用户的邮箱信息
func ListEmailsByToken(githubToken string) (emails []*github.UserEmail,err *util.Err) {

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	tc := oauth2.NewClient(context.Background(), ts)

	client := github.NewClient(tc)
	emails, _, err2 := client.Users.ListEmails(context.Background(), nil)

	if err2 != nil{
		err = util.NewErr(err2,util.ErrThirdParty,"从github获取Email错误")
	} else{
		log.Debug("获取的email是",emails)
	}

	return
}

//获取用户的所有币种的余额
func GetUserBalanceByAllAssets(userId uint32,assets *[]model.Asset)(err *util.Err,dto *[]model.MemberWalletDto){

	dto = &[]model.MemberWalletDto{}
	//遍历assets数组获取所有的币种
	for i := range *assets{
		tmp := model.MemberWalletDto{}
		tmp.AssetId = (*assets)[i].AssetId

		memberWalletDtos,err1 := dao.GetMemeberWalletByUserIdAndAssetId(userId,(*assets)[i].AssetId)
		if err1 != nil {
			err = util.NewErr(err1,util.ErrDataBase,"查询数据库的用户钱包出错")
			return
		}
		//把balance相加到tmp里面
		if memberWalletDtos != nil{
			log.Debug(*memberWalletDtos)
			for j := range *memberWalletDtos{
				tmp.Balance = ((*memberWalletDtos)[j].Balance.Mul((*assets)[i].PriceUsd) ).Add(tmp.Balance)
			}
		}
		*dto = append(*dto, tmp)
	}
	return
}

func GetTransferByUserIdAndAssetId(userId uint32,assetId string)(transfers *[]model.Transfer,err *util.Err) {
	transfers,err1 := dao.GetTransferByUserIdAndAssetId(userId,assetId)
	log.Debug("service transfers = ",transfers)
	if err1 != nil{
		err = util.NewErr(err1,util.ErrDataBase,"数据库查询transfer出错")
	}
	return
}

func GetPatronsByUserId(userId uint32) (patrons uint32,err *util.Err) {
	projects,err1 := dao.ListProjectsByUserId(userId)
	if err1 != nil{
		err = util.NewErr(err,util.ErrDataBase,"数据库获取所有project信息出错")
		return
	}

	//遍历projects,把patrons相加
	for i := range *projects{
		patrons = patrons + (*projects)[i].Donations
	}
	return
}

//更新user表中的mixin_id字段
func UpdateUserMixinId(user_id uint32,mixinId string) (err *util.Err) {
	err1 := dao.UpdateUserMixinId(user_id,mixinId)
	if err1 != nil{
		err = util.NewErr(err1,util.ErrDataBase,"更新数据库mixin_id错误")
	}
	return
}