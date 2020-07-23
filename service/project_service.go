package service

import (
	"claps-test/dao"
	"claps-test/model"
	"claps-test/util"
	"errors"
)

//通过projectName查询,查询某个项目的详情
func GetProjectByName(name string) (projectDetailInfo *map[string]interface{},err *util.Err){

	project,err1 := dao.GetProjectByName(name)
	if err1 != nil {
		err = util.NewErr(err1,util.ErrDataBase,"获取项目信息失败")
		return
	}

	repositories,err1 := dao.ListRepositoriesByProjectId(project.Id)
	if err1 != nil {
		err = util.NewErr(err1,util.ErrDataBase,"获取项目仓库失败")
		return
	}


	//mambers格式不同,删除project_id和userid字段
	members,err1 := dao.ListMembersByProjectName(project.Name)
	if err1 != nil {
		err = util.NewErr(err1,util.ErrDataBase,"获取项目成员失败")
		return
	}


	botDtos,err1 := dao.ListBotDtosByProjectId(project.Id)
	if err1 != nil {
		err = util.NewErr(err1,util.ErrDataBase,"获取项目机器人失败")
		return
	}


	projectDetailInfo = &map[string]interface{}{
		"project":project,
		"repositories":repositories,
		"members":members,
		"botIds":botDtos,
	}
	return
}

//获取数据库中所有project
func ListProjectsAll() (projects *[]model.Project,err *util.Err){
	projects,err1 := dao.ListProjectsAll()
	if err1 != nil {
		err = util.NewErr(err1,util.ErrDataBase,"获取所有项目失败")
	}
	return
}


//查询某用户的所有项目,获取数据库中所有project
func ListProjectsByUserId(userId uint32) (projects *[]model.Project,err *util.Err){
	projects,err1 := dao.ListProjectsByUserId(userId)
	if err1 != nil {
		err = util.NewErr(err1,util.ErrDataBase,"获取项目机器人失败")
	}

	return
}

func ListTransactionsByProjectNameAndAssetId(name string,assetId string)(transactions *[]model.Transaction,err *util.Err){
	if assetId == ""{
		err = util.NewErr(errors.New("没有QUERY值"),util.ErrUnauthorized,"没有QUERY值无法请求成功")
		return
	}

	transactions,err1 := dao.ListTransactionsByProjectNameAndAssetId(name,assetId)
	if err1 != nil {
		err = util.NewErr(err1,util.ErrDataBase,"获取项目获取捐赠记录失败")
	}
	return
}

func ListMembersByProjectName(projectName string)(members *[]model.User,err *util.Err){
	members,err1 := dao.ListMembersByProjectName(projectName)
	if err1 != nil {
		err = util.NewErr(err1,util.ErrDataBase,"获取项目成员失败")
	}
	return
}
