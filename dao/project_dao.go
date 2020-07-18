package dao

import (
	"claps-test/util"
	"claps-test/model"
)

//通过项目Id获取项目的Total?
func GetProjectTotalById(projectId uint32)(total *[]model.ProjectTotal,err error){
	total = &[]model.ProjectTotal{}
	err = util.DB.Debug().Table("wallet").Select("asset_id,total").Where("project_id=?",projectId).Scan(total).Error
	return
}

//获取所有项目
func ListProjectsAll() (projects *[]model.Project,err error){

	projects = &[]model.Project{}
	err = util.DB.Debug().Find(projects).Error
	return
}

//通过项目名字获取项目
func GetProjectByName(name string) (project *model.Project,err error){

	project = &model.Project{}
	err = util.DB.Debug().Where("name=?",name).Find(&project).Error
	return
}

//根据userid获取所有项目
func ListProjectsByUserId(userId int64)(projects *[]model.Project,err error){
	projects = &[]model.Project{}
	err = util.DB.Debug().Where("id IN(?)",
		util.DB.Debug().Table("member").Select("project_id").Where("user_id=?",userId).SubQuery()).Find(projects).Error
	return
}

func GetProjectByBotId(BotId string)(projectTotal *model.ProjectTotals,err error){
	projectTotal = &model.ProjectTotals{}
	err = util.DB.Debug().Table("project").Where("project_id=?",
		util.DB.Debug().Table("bot").Select("project_id").Where("id=?",BotId).SubQuery()).Find(projectTotal).Error
	return
}

func UpdateProject(projectTotal *model.ProjectTotals)(err error){
	err = util.DB.Debug().Table("project").Save(projectTotal).Error
	return
}