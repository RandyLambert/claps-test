package dao

import (
	"claps-test/model"
	"claps-test/util"
)

//根据project获取所有的仓库信息
func ListRepositoriesByProjectId(projectId uint32)(repositories *[]model.Repository,err error){
	repositories = &[]model.Repository{}
	err = util.DB.Debug().Where("project_id=?",projectId).Find(repositories).Error
	return
}
