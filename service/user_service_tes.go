package service

import (
	"fmt"
)
func TetGetUserBalanceByAllAssets() {
	assets, err := ListAssetsAllByDB()
	if err != nil {
		return
	}
	//查询用户钱包,获得相应的余额,添加到币信息的后面
	err2, dto := GetBalanceAndTotalByUserIdAndAssets(46085959, assets)
	if err2 != nil {
		return
	}
	fmt.Println(dto)
	err2,total,balance := GetBalanceAndTotalToUSDByUserId(46085959,assets)
	if err2 != nil {
		return
	}
	fmt.Println("total",total,"balance",balance)
}
