package dao

import (
	"claps-test/model"
	"claps-test/util"
)

//从数据库中通过ID获取user信息,存储在user中,引用传值
func SelectUserById(user *model.User,Id int64) {
	util.DB.Debug().First(user,Id)
	return
}

//不管记录是否找到，都将参数赋值给 struct 并保存至数据库
func CreateOrUpdateUser(user *model.User) {
	util.DB.Debug().FirstOrCreate(user)
	return
}