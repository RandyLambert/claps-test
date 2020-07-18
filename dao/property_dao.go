package dao

import (
	"claps-test/model"
	"claps-test/util"
)

func GetPropertyByKey(Key string)(property *model.Property,err error){
	property = &model.Property{
		Key:Key,
	}
	err = util.DB.First(property).Error
	return
}

func UpdateProperty(property *model.Property)(err error){
	err = util.DB.Save(property).Error
	return
}

func InsertProperty(property *model.Property)(err error){
	err = util.DB.Create(property).Error
	return
}
