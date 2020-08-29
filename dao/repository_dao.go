package dao

import (
	"claps-test/model"
	"claps-test/util"
)

//根据project获取所有的仓库信息
func ListRepositoriesByProjectId(projectId int64) (repositories *[]model.Repository, err error) {
	repositories = &[]model.Repository{}
	err = util.DB.Debug().Table("repository").Where("project_id=?", projectId).Scan(repositories).Error
	return
}
