package service

import (
	"claps-test/dao"
	"claps-test/model"
	"claps-test/util"
)

//通过projectName查询,查询某个项目的详情
func GetProjectByName(name string) (projectDetailInfo *map[string]interface{},err *util.Err){

	project,err1 := dao.GetProjectByName(name)
	if err1 != nil {
		err = util.NewErr(err1,util.ErrDataBase,"获取项目信息失败")
		return
	}

	repositories,err := ListRepositoriesByProjectId(project.Id)
	if err != nil {
		return
	}

	members,err := ListMembersByProjectName(name)
	if err != nil {
		return
	}

	botIds,err := ListBotDtosByProjectId(project.Id)
	if err != nil {
		return
	}

	projectDetailInfo = &map[string]interface{}{
		"project":project,
		"repositories":repositories,
		"members":members,
		"botIds":botIds,
	}
	return
}

//获取数据库中所有project
func ListProjectsAll() (projects *[]model.Project,err *util.Err){
	projects,err1 := dao.ListProjectsAll()
	if err1 != nil {
		err = util.NewErr(err1,util.ErrDataBase,"获取所有项目失败")
		return
	}

	return
}


//查询某用户的所有项目,获取数据库中所有project
func ListProjectsByUserId(userId int64) (projects *[]model.Project,err *util.Err){
	projects,err1 := dao.ListProjectsByUserId(userId)
	if err1 != nil {
		err = util.NewErr(err1,util.ErrDataBase,"获取项目机器人失败")
		return
	}

	return
}

func ListTransactionsByProjectNameAndAssetId(name string,assetId string)(transactions *[]model.Transaction,err *util.Err){
	transactions,err1 := dao.ListTransactionsByProjectNameAndAssetId(name,assetId)
	if err1 != nil {
		err = util.NewErr(err1,util.ErrDataBase,"获取项目获取捐赠记录失败")
	}
	return
}

func ListRepositoriesByProjectId(projectId uint32)(repositories *[]model.Repository,err *util.Err){
	repositories,err1 := dao.ListRepositoriesByProjectId(projectId)
	if err1 != nil {
		err = util.NewErr(err1,util.ErrDataBase,"获取项目仓库失败")
	}
	return
}

func ListMembersByProjectName(name string)(members *[]model.User,err *util.Err){
	//mambers格式不同,删除project_id和userid字段
	members,err1 := dao.ListMembersByProjectName(name)
	if err1 != nil {
		err = util.NewErr(err1,util.ErrDataBase,"获取项目成员失败")
	}
	return
}

func ListBotDtosByProjectId(projectId uint32)(botDtos *[]model.BotDto,err *util.Err){
	botDtos,err1 := dao.ListBotDtosByProjectId(projectId)
	if err1 != nil {
		err = util.NewErr(err1,util.ErrDataBase,"获取项目机器人失败")
	}
	return
}