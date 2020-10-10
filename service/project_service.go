package service

import (
	"claps-test/model"
	"claps-test/util"
	log "github.com/sirupsen/logrus"
)

//通过projectid查询,查询某个项目的详情
func GetProjectById(projectId int64) (projectDetailInfo *map[string]interface{}, err *util.Err) {

	project, err1 := model.PROJECT.GetProjectById(projectId)
	if err1 != nil {
		err = util.NewErr(err1, util.ErrDataBase, "获取项目信息失败")
		return
	}

	repositoryDtos, err1 := model.REPOSITORY.ListRepositoriesByProjectId(projectId)
	if err1 != nil {
		err = util.NewErr(err1, util.ErrDataBase, "获取项目仓库失败")
		return
	}

	//mambers格式不同,删除project_id和userid字段
	members, err1 := model.USER.ListMembersByProjectId(projectId)
	if err1 != nil {
		err = util.NewErr(err1, util.ErrDataBase, "获取项目成员失败")
		return
	}

	botDtos, err1 := model.BOT.ListBotDtosByProjectId(projectId)
	if err1 != nil {
		err = util.NewErr(err1, util.ErrDataBase, "获取项目机器人失败")
		return
	}

	projectDetailInfo = &map[string]interface{}{
		"project":      project,
		"repositories": repositoryDtos,
		"members":      members,
		"botIds":       botDtos,
	}
	return
}

//获取数据库中所有project
func ListProjectsAll() (projects *[]model.Project, err *util.Err) {
	projects, err1 := model.PROJECT.ListProjectsAll()
	if err1 != nil {
		err = util.NewErr(err1, util.ErrDataBase, "获取所有项目失败")
	}
	return
}

func ListProjectsByQuery(q *model.PaginationQ) (projects *[]model.Project,number int, err *util.Err){
	projects,number,err1 := model.PROJECT.ListProjectsByQuery(q)
	if err1 != nil {
		err = util.NewErr(err1, util.ErrDataBase, "根据query获取项目失败失败")
	}
	return
}

//查询某用户的所有项目,获取数据库中所有project
func ListProjectsByUserId(userId int64) (projects *[]model.Project, err *util.Err) {
	projects, err1 := model.PROJECT.ListProjectsByUserId(userId)
	if err1 != nil {
		err = util.NewErr(err1, util.ErrDataBase, "获取项目机器人失败")
	}

	return
}

func ListTransactionsByProjectId(projectId int64) (transactions *[]model.Transaction, err *util.Err) {

	transactions, err1 := model.TRANSACTION.ListTransactionsByProjectId(projectId)
	if err1 != nil {
		err = util.NewErr(err1, util.ErrDataBase, "获取项目获取捐赠记录失败")
	}
	return
}

func ListTransactionsByProjectIdAndQuery(projectId int64,q *model.PaginationQ) (transactions *[]model.Transaction, number int, err *util.Err) {

	transactions, number, err1 := model.TRANSACTION.ListTransactionsByProjectIdAndQuery(projectId,q)
	if err1 != nil {
		err = util.NewErr(err1, util.ErrDataBase, "获取项目获取捐赠记录失败")
	}
	return
}

func ListMembersByProjectId(projectId int64) (members *[]model.User, err *util.Err) {
	members, err1 := model.USER.ListMembersByProjectId(projectId)
	if err1 != nil {
		err = util.NewErr(err1, util.ErrDataBase, "获取项目成员失败")
	}
	return
}

func GetProjectBadge(badge *model.Badge) (err *util.Err) {
	//compact
	//full

	fiat, err1 := model.FIAT.GetFiatByCode(badge.Code)
	if err1 != nil {
		err = util.NewErr(err1, util.ErrDataBase, "获取fiat失败")
	}
	log.Debug(fiat)
	return
}
