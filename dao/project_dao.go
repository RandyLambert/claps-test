package dao

import (
	"claps-test/common"
	"claps-test/models"
)

func GetProjectTotal(projectId uint32)(total *[]models.ProjectTotal,err error){
	total = &[]models.ProjectTotal{}
	err = common.DB.Debug().Table("wallet").Select("asset_id,total").Where("project_id=?",projectId).Scan(total).Error
	return
}

func GetProjectPatrons(projectId uint32)(count int,err error){
	//db.Table("deleted_users").Select("count(distinct(name))").Count(&count)
	err = common.DB.Debug().Table("transaction").Select("count(distinct(sender))").Where("project_id=?",projectId).Count(&count).Error
	return
}

func GetProjects() (projects *[]models.Project,err error){

	projects = &[]models.Project{}
	err = common.DB.Debug().Find(projects).Error
	return
}

func GetProject(name string) (project *models.Project,err error){

	project = &models.Project{}
	err = common.DB.Debug().Where("name=?",name).Find(&project).Error
	return
}

func GetProjectMembers(name string)(users *[]models.User,err error){

	//db.Where("amount > ?", db.Table("orders").Select("AVG(amount)").Where("state = ?", "paid").SubQuery()).Find(&orders)
	// SELECT * FROM "orders"  WHERE "orders"."deleted_at" IS NULL AND (amount > (SELECT AVG(amount) FROM "orders"  WHERE (state = 'paid')));
	users = &[]models.User{}
	err = common.DB.Debug().Where("user.id IN (?)",
		common.DB.Debug().Table("member").Select("project_id=?",
			common.DB.Debug().Table("project").Select("project.id").Where("project.name=?",name).SubQuery()).SubQuery()).Find(&users).Error
	// IN
	//db.Where("name IN (?)", []string{"jinzhu", "jinzhu 2"}).Find(&users)
	//// SELECT * FROM users WHERE name in ('jinzhu','jinzhu 2');

	return
}

func GetProjectTransactions(name string,assetId string)(transactions *[]models.Transaction,err error){

	transactions = &[]models.Transaction{}
	err = common.DB.Debug().Joins("INNER JOIN project ON project.name=?",name).Where("asset_id=?",assetId).Find(transactions).Error
	return
}

func GetProjectRepositories(projectId uint32)(repositories *[]models.Repository,err error){
	repositories = &[]models.Repository{}
	err = common.DB.Debug().Where("project_id=?",projectId).Find(repositories).Error
	return
}

func GetProjectBotIds(projectId uint32)(botId *[]models.BotId,err error){
	botId = &[]models.BotId{}
	err = common.DB.Debug().Table("bot").Where("project_id=?",projectId).Find(botId).Error
	return
}