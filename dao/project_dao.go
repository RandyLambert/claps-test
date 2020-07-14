package dao

import (
	"claps-test/util"
	"claps-test/model"
)

func GetProjectTotal(projectId uint32)(total *[]model.ProjectTotal,err error){
	total = &[]model.ProjectTotal{}
	err = util.DB.Debug().Table("wallet").Select("asset_id,total").Where("project_id=?",projectId).Scan(total).Error
	return
}

func GetProjectPatrons(projectId uint32)(count int,err error){
	//db.Table("deleted_users").Select("count(distinct(name))").Count(&count)
	err = util.DB.Debug().Table("transaction").Select("count(distinct(sender))").Where("project_id=?",projectId).Count(&count).Error
	return
}

func GetProjects() (projects *[]model.Project,err error){

	projects = &[]model.Project{}
	err = util.DB.Debug().Find(projects).Error
	return
}

func GetProject(name string) (project *model.Project,err error){

	project = &model.Project{}
	err = util.DB.Debug().Where("name=?",name).Find(&project).Error
	return
}

func GetProjectMembers(name string)(users *[]model.User,err error){

	//db.Where("amount > ?", db.Table("orders").Select("AVG(amount)").Where("state = ?", "paid").SubQuery()).Find(&orders)
	// SELECT * FROM "orders"  WHERE "orders"."deleted_at" IS NULL AND (amount > (SELECT AVG(amount) FROM "orders"  WHERE (state = 'paid')));
	users = &[]model.User{}
	err = util.DB.Debug().Where("user.id IN (?)",
		util.DB.Debug().Table("member").Select("user_id").Where("project_id=?",
			util.DB.Debug().Table("project").Select("project.id").Where("project.name=?",name).SubQuery()).SubQuery()).Find(&users).Error
	// IN
	//db.Where("name IN (?)", []string{"jinzhu", "jinzhu 2"}).Find(&users)
	//// SELECT * FROM users WHERE name in ('jinzhu','jinzhu 2');

	return
}

func GetProjectTransactions(name string,assetId string)(transactions *[]model.Transaction,err error){

	transactions = &[]model.Transaction{}
	err = util.DB.Debug().Where("asset_id=?",assetId).Where("project.id=?",
		util.DB.Debug().Table("project").Select("project.id").Where("project.name=?",name).SubQuery()).Find(transactions).Error
	//err = util.DB.Debug().Joins("INNER JOIN project ON project.name=?",name).Where("asset_id=?",assetId).Find(transactions).Error
	return
}

func GetProjectRepositories(projectId uint32)(repositories *[]model.Repository,err error){
	repositories = &[]model.Repository{}
	err = util.DB.Debug().Where("project_id=?",projectId).Find(repositories).Error
	return
}

func GetProjectBotIds(projectId uint32)(botId *[]model.BotId,err error){
	botId = &[]model.BotId{}
	err = util.DB.Debug().Table("bot").Where("project_id=?",projectId).Find(botId).Error
	return
}