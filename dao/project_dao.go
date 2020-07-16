package dao

import (
	"claps-test/util"
	"claps-test/model"
)

//通过项目Id获取项目的Total?
func GetProjectTotal(projectId uint32)(total *[]model.ProjectTotal,err error){
	total = &[]model.ProjectTotal{}
	err = util.DB.Debug().Table("wallet").Select("asset_id,total").Where("project_id=?",projectId).Scan(total).Error
	return
}

//根据项目Id获取项目的所有捐赠
func GetProjectPatrons(projectId uint32)(count int,err error){
	//db.Table("deleted_users").Select("count(distinct(name))").Count(&count)
	err = util.DB.Debug().Table("transaction").Select("count(distinct(sender))").Where("project_id=?",projectId).Count(&count).Error
	return
}

//获取所有项目
func GetProjects() (projects *[]model.Project,err error){

	projects = &[]model.Project{}
	err = util.DB.Debug().Find(projects).Error
	return
}

//通过项目名字获取项目
func GetProject(name string) (project *model.Project,err error){

	project = &model.Project{}
	err = util.DB.Debug().Where("name=?",name).Find(&project).Error
	return
}

//通过projectName获取一个项目的所有成员信息
func GetProjectMembers(name string)(users *[]model.User,err error){

	//db.Where("amount > ?", db.Table("orders").Select("AVG(amount)").Where("state = ?", "paid").SubQuery()).Find(&orders)
	// SELECT * FROM "orders"  WHERE "orders"."deleted_at" IS NULL AND (amount > (SELECT AVG(amount) FROM "orders"  WHERE (state = 'paid')));
	users = &[]model.User{}
	err = util.DB.Debug().Where("user.id IN (?)",
		util.DB.Debug().Table("member").Select("user_id").Where("project_id=?",
			util.DB.Debug().Table("project").Select("project.id").Where("project.name=?",name).SubQuery()).SubQuery()).Find(users).Error
	// IN
	//db.Where("name IN (?)", []string{"jinzhu", "jinzhu 2"}).Find(&users)
	//// SELECT * FROM users WHERE name in ('jinzhu','jinzhu 2');

	return
}


//获取捐赠记录:assetId货币种类,name??
func GetProjectTransactions(name string,assetId string)(transactions *[]model.Transaction,err error){

	transactions = &[]model.Transaction{}
	err = util.DB.Debug().Where("asset_id=?",assetId).Where("project.id=?",
		util.DB.Debug().Table("project").Select("project.id").Where("project.name=?",name).SubQuery()).Find(transactions).Error
	//err = util.DB.Debug().Joins("INNER JOIN project ON project.name=?",name).Where("asset_id=?",assetId).Find(transactions).Error
	return
}

//根据project获取所有的仓库信息
func GetProjectRepositories(projectId uint32)(repositories *[]model.Repository,err error){
	repositories = &[]model.Repository{}
	err = util.DB.Debug().Where("project_id=?",projectId).Find(repositories).Error
	return
}

//根据projectId获取所有的机器人Id
func GetProjectBotIds(projectId uint32)(botId *[]model.BotId,err error){
	botId = &[]model.BotId{}
	err = util.DB.Debug().Table("bot").Where("project_id=?",projectId).Find(botId).Error
	return
}

//根据userid获取所有项目
func GetProjectsByUserId(userId int64)(projects *[]model.Project,err error){
	projects = &[]model.Project{}
	err = util.DB.Debug().Where("id IN(?)",
		util.DB.Debug().Table("member").Select("project_id").Where("user_id=?",userId).SubQuery()).Find(projects).Error
	return
}