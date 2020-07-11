package dao

import (
	"claps-test/common"
	"claps-test/models"
)

func GetProjects() (projects *[]models.Project){
	db := common.GetDB()
	projects = &[]models.Project{}
	db.Debug().Find(&projects)
	return
}

func GetProject(name string) (project *models.Project){
	db := common.GetDB()
	project = &models.Project{}
	db.Debug().Where("name=?",name).Find(&project)

	return
}

func GetProjectMembers(name string)(users *[]models.User){
	db := common.GetDB()
	db.Debug().Joins("INNER JOIN member ON member.user_id = user.id").Joins("INNER JOIN project ON project.name=?",name).Where("project.id=?","member.project_id").Find(&users)
	return
}

func GetProjectTransactions(name string,assetId string)(transactions *[]models.Transaction){
	db := common.GetDB()
	transactions = &[]models.Transaction{}
	db.Debug().Joins("INNER JOIN project ON project.name=?",name).Where("asset_id=?",assetId).Find(&transactions)
	return
}