package dao

import (
	"claps-test/model"
	"github.com/jinzhu/gorm"
)

func init() {
	RegisterMigrateHandler(func(db *gorm.DB) error {

		if err := db.AutoMigrate(&model.User{}).Error; err != nil {
			return err
		}
		return nil
	})
}

func init() {
	RegisterMigrateHandler(func(db *gorm.DB) error {

		if err := db.AutoMigrate(&model.Member{}).Error; err != nil {
			return err
		}
		return nil
	})
}

//从数据库中通过ID获取user信息,存储在user中,引用传值
func GetUserById(id int64) (user *model.UserMixinId, err error) {
	user = &model.UserMixinId{}
	err = db.Debug().Table("user").Where("id = ?", id).Scan(user).Error
	return
}

//不管记录是否找到，都将参数赋值给 struct 并保存至数据库
func InsertOrUpdateUser(user *model.User) (err error) {

	var cnt int64
	db.Debug().Table("user").Where("id = ?", user.Id).Count(&cnt)
	if cnt == 0 {
		err = db.Debug().Create(user).Error
		return
	} else {
		db.Debug().Model(&user).Omit("mixin_id").Updates(user)
	}
	return
}

//通过projectId获取一个项目的所有成员信息
func ListMembersByProjectId(projectId int64) (users *[]model.User, err error) {

	users = &[]model.User{}
	err = db.Debug().Where("id IN (?)",
		db.Debug().Table("member").Select("user_id").Where("project_id=?", projectId).SubQuery()).Find(users).Error
	return
}

//根据user_id更新表中的mixin_id信息
func UpdateUserMixinId(userId int64, mixinId string) (err error) {

	err = db.Debug().Model(&model.User{}).Where("id = ?", userId).Update("mixin_id", mixinId).Error
	return
}
