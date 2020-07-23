package dao

import (
	"claps-test/model"
	"claps-test/util"
)

//根据project获取所有的仓库信息
func ListRepositoriesByProjectId(projectId uint32)(repositoriesDto *[]model.RepositoryDto,err error){
	repositoriesDto = &[]model.RepositoryDto{}
	err = util.DB.Debug().Table("repository").Where("project_id=?",projectId).Scan(repositoriesDto).Error
	return
}
