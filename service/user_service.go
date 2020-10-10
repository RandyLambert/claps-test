package service

import (
	"claps-test/model"
	"claps-test/util"
	"context"
	"github.com/google/go-github/v32/github"
	"github.com/shopspring/decimal"
	"golang.org/x/oauth2"
	//log "github.com/sirupsen/logrus"
)

//从github服务器请求获取用户的邮箱信息
func ListEmailsByToken(githubToken string) (emails []*github.UserEmail, err *util.Err) {

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	tc := oauth2.NewClient(context.Background(), ts)

	client := github.NewClient(tc)
	emails, _, err2 := client.Users.ListEmails(context.Background(), nil)

	if err2 != nil {
		err = util.NewErr(err2, util.ErrThirdParty, "从github获取Email错误")
	}

	return
}

func GetBalanceAndTotalToUSDByUserId(userId int64, assets *[]model.Asset) (err *util.Err, total decimal.Decimal, balance decimal.Decimal) {

	//遍历assets数组获取所有的币种
	var assetMap map[string]decimal.Decimal
	assetMap = make(map[string]decimal.Decimal)
	//生成币种对应ｍａｐ方便后面调用
	for _, asset := range *assets {
		assetMap[asset.AssetId] = asset.PriceUsd
	}

	memberWalletDtos, err1 := model.MEMBERWALLET.GetMemberWalletByUserId(userId)
	if err1 != nil {
		err = util.NewErr(err1, util.ErrDataBase, "查询数据库的用户钱包出错")
		return
	}

	//把balance相加到tmp里面
	if memberWalletDtos != nil {
		for _, value := range *memberWalletDtos {
			balance = (value.Balance.Mul(assetMap[value.AssetId])).Add(balance)
			total = (value.Total.Mul(assetMap[value.AssetId])).Add(total)
		}
	}
	total = total.Truncate(4)
	balance = balance.Truncate(4)
	return
}

//获取用户的所有币种的余额
func GetBalanceAndTotalByUserIdAndAssets(userId int64, assets *[]model.Asset) (err *util.Err, dto *[]model.MemberWalletDto) {

	//遍历assets数组获取所有的币种
	var memberWalletMap map[string]*model.MemberWalletDto
	memberWalletMap = make(map[string]*model.MemberWalletDto)

	//生成币种对应ｍａｐ方便后面调用
	for _, asset := range *assets {
		memberWalletMap[asset.AssetId] = &model.MemberWalletDto{AssetId: asset.AssetId}
	}

	memberWalletDtos, err1 := model.MEMBERWALLET.GetMemberWalletByUserId(userId)
	if err1 != nil {
		err = util.NewErr(err1, util.ErrDataBase, "查询数据库的用户钱包出错")
		return
	}

	dto = &[]model.MemberWalletDto{}
	//把balance相加到tmp里面
	if memberWalletDtos != nil {
		for _, value := range *memberWalletDtos {
			memberWalletMap[value.AssetId].Balance = value.Balance.Add(memberWalletMap[value.AssetId].Balance)
			memberWalletMap[value.AssetId].Total = value.Total.Add(memberWalletMap[value.AssetId].Total)
		}

		for _, memberWallet := range memberWalletMap {
			memberWallet.Balance = memberWallet.Balance.Truncate(8)
			memberWallet.Total = memberWallet.Total.Truncate(8)
			*dto = append(*dto, *memberWallet)
		}
	}

	return
}

func ListTransfersByMixinId(mixinId string) (transfers *[]model.Transfer, err *util.Err) {
	transfers, err1 := model.TRANSFER.ListTransferByMixinId(mixinId)
	if err1 != nil {
		err = util.NewErr(err1, util.ErrDataBase, "数据库查询transfer出错")
	}
	return
}

func ListTransfersByProjectIdAndQuery(mixinId string,q *model.PaginationQ) (transfers *[]model.Transfer, number int, err *util.Err) {

	transfers, number, err1 := model.TRANSFER.ListTransfersByMixinIdAndQuery(mixinId,q)
	if err1 != nil {
		err = util.NewErr(err1, util.ErrDataBase, "获取项目获取捐赠记录失败")
	}
	return
}

func SumProjectDonationsByUserId(userId int64) (donations int64, err *util.Err) {
	donations, err1 := model.PROJECT.SumProjectDonationsByUserId(userId)
	if err1 != nil {
		err = util.NewErr(err, util.ErrDataBase, "数据库获取用户项目信息和出错")
	}
	return
}

//更新user表中的mixin_id字段
func UpdateUserMixinId(userId int64, mixinId string) (err *util.Err) {
	err1 := model.USER.UpdateUserMixinId(userId, mixinId)
	if err1 != nil {
		err = util.NewErr(err1, util.ErrDataBase, "更新数据库mixin_id错误")
	}
	return
}

func GetMixinIdByUserId(userId int64) (mixinId string, err *util.Err) {
	user, err1 := model.USER.GetUserById(userId)
	if err1 != nil {
		err = util.NewErr(err1, util.ErrDataBase, "从数据库查询user信息错误")
		return
	}
	mixinId = user.MixinId
	return
}

func UpdateUserWithdrawalWay(userId int64,withdrawWal string)(err *util.Err){
	err1 := model.USER.UpdateUserWithdrawalWay(userId,withdrawWal)
	if err1 != nil {
		err = util.NewErr(err1, util.ErrDataBase, "更新数据库withdrawalWay信息错误")
	}
	return
}