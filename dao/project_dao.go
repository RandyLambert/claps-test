package dao

import (
	"claps-test/model"
	"claps-test/util"
)

//获取所有项目
func ListProjectsAll() (projects *[]model.Project, err error) {

	projects = &[]model.Project{}
	err = util.DB.Debug().Find(projects).Error
	return
}

//通过项目名字获取项目
func GetProjectByName(name string) (project *model.Project, err error) {

	project = &model.Project{}
	err = util.DB.Debug().Where("name=?", name).Find(&project).Error
	return
}

//根据userid获取所有项目
func ListProjectsByUserId(userId uint32) (projects *[]model.Project, err error) {
	projects = &[]model.Project{}
	err = util.DB.Debug().Where("id IN(?)",
		util.DB.Debug().Table("member").Select("project_id").Where("user_id=?", userId).SubQuery()).Find(projects).Error
	return
}

func GetProjectTotalByBotId(BotId string) (projectTotal *model.ProjectTotal, err error) {
	projectTotal = &model.ProjectTotal{}
	err = util.DB.Debug().Table("project").Select("id,donations,total").Where("id=?",
		util.DB.Debug().Table("bot").Select("project_id").Where("id=?", BotId).SubQuery()).Scan(projectTotal).Error
	return
}

func UpdateProjectTotal(projectTotal *model.ProjectTotal) (err error) {
	err = util.DB.Debug().Table("project").Save(projectTotal).Error
	return
}
