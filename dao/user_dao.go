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


//从数据库中通过ID获取user信息,存储在user中,引用传值
func GetUserByUserId(id int64) (user *model.UserMixinId, err error) {
	user = &model.UserMixinId{}
	err = db.Debug().Table("user").Where("id = ?", id).Scan(user).Error
	return
}

//不管记录是否找到，都将参数赋值给 struct 并保存至数据库
func InsertOrUpdateUser(user *model.User) (err error) {
	err = db.Debug().Save(user).Error
	/*
		var cnt int64
		util.DB.Debug().Table("user").Where("id = ?",user.Id).Count(&cnt)
		if cnt == 0{
			err = util.DB.Debug().Create(user).Error
			return
		}
	*/
	//util.DB.Save(user)
	return
}

//通过projectName获取一个项目的所有成员信息
func ListMembersByProjectName(projectName string) (users *[]model.User, err error) {

	//db.Where("amount > ?", db.Table("orders").Select("AVG(amount)").Where("state = ?", "paid").SubQuery()).Find(&orders)
	// SELECT * FROM "orders"  WHERE "orders"."deleted_at" IS NULL AND (amount > (SELECT AVG(amount) FROM "orders"  WHERE (state = 'paid')));
	users = &[]model.User{}
	err = db.Debug().Where("id IN (?)",
		db.Debug().Table("member").Select("user_id").Where("project_id=?",
			db.Debug().Table("project").Select("project.id").Where("project.name=?", projectName).SubQuery()).SubQuery()).Find(users).Error
	// IN
	//db.Where("name IN (?)", []string{"jinzhu", "jinzhu 2"}).Find(&users)
	//// SELECT * FROM users WHERE name in ('jinzhu','jinzhu 2');

	return
}

func ListMembersByProjectId(projectId int64) (users *[]model.User, err error) {

	//db.Where("amount > ?", db.Table("orders").Select("AVG(amount)").Where("state = ?", "paid").SubQuery()).Find(&orders)
	// SELECT * FROM "orders"  WHERE "orders"."deleted_at" IS NULL AND (amount > (SELECT AVG(amount) FROM "orders"  WHERE (state = 'paid')));
	users = &[]model.User{}
	err = db.Debug().Where("id IN (?)",
		db.Debug().Table("member").Select("user_id").Where("project_id=?", projectId).SubQuery()).Find(users).Error
	// IN
	//db.Where("name IN (?)", []string{"jinzhu", "jinzhu 2"}).Find(&users)
	//// SELECT * FROM users WHERE name in ('jinzhu','jinzhu 2');

	return
}

//根据user_id更新表中的mixin_id信息
func UpdateUserMixinId(userId int64, mixinId string) (err error) {
	err = db.Debug().Model(&model.User{}).Where("id = ?", userId).Update("mixin_id", mixinId).Error
	return
}
