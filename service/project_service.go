package service

import (
	"claps-test/dao"
	"claps-test/model"
	//"github.com/fox-one/mixin-sdk-go"
	//"github.com/shopspring/decimal"
)

//通过projectName查询,查询某个项目的详情
func GetProjectByName(name string) (projectDetailInfo *map[string]interface{},err error){

	project,err := dao.GetProjectByName(name)
	if err != nil {
		return
	}

	//projectInfo,err := GetProjectInfo(project)
	//if err != nil {
	//	return
	//}
	repositories,err := ListRepositoriesByProjectId(project.Id)
	if err!=nil {
		return
	}
	members,err := ListMembersByProjectName(name)
	if err!=nil {
		return
	}
	botIds,err := ListBotIdsByProjectId(project.Id)
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
func ListProjectsAll() (projects *[]model.Project,err error){
	projects,err = dao.ListProjectsAll()
	if err != nil{
		return
	}

	//for i :=range *projects {
	//	var projectInfo *map[string]interface{}
		//projectInfo,err = GetProjectInfo(&(*projects)[i])
	//	if err != nil {
	//		return
	//	}
	//	projectInfos = append(projectInfos,projectInfo)
	//}
	return
}


//查询某用户的所有项目,获取数据库中所有project
func ListProjectsByUserId(userId int64) (projects *[]model.Project,err error){
	projects,err = dao.ListProjectsByUserId(userId)
	//if err != nil{
	//	return
	//}

	//以后定期获取金额的时候可能使用
	//for i :=range *projects {
	//	var projectInfo *map[string]interface{}
	//	projectInfo,err = GetProjectInfo(&(*projects)[i])
	//	if err != nil {
	//		return
	//	}
	//	projectInfos = append(projectInfos,projectInfo)
	//}
	return
}

//以后定时获取金额时使用
//func GetProjectInfo(project *model.Project)(projectInfo *map[string]interface{},err error){
//
//	projectTotal,err := dao.GetProjectTotal(project.Id)
//	if err != nil {
//		return
//	}
//	total := decimal.Zero
//	for i := range *projectTotal {
//		var assert *mixin.Asset
//		assert,err = GetAsset((*projectTotal)[i].AssetId)
//		if err != nil {
//			return
//		}
//		assertTotal := decimal.NewFromFloat((*projectTotal)[i].Total)
//		total.Add(assert.PriceUSD.Mul(assertTotal))
//	}
//
//	patrons,err := dao.GetProjectPatrons(project.Id)
//	projectInfo = &map[string]interface{}{
//		"patrons":patrons,
//		"total":total,
//		"project":project,
//	}
//	return
//}

func ListTransactionsByProjectNameAndAssetId(name string,assetId string)(transactions *[]model.Transaction,err error){
	transactions,err = dao.ListTransactionsByProjectNameAndAssetId(name,assetId)
	return
}

func ListRepositoriesByProjectId(projectId uint32)(repositories *[]model.Repository,err error){
	repositories,err = dao.ListRepositoriesByProjectId(projectId)
	return
}

func ListMembersByProjectName(name string)(members *[]model.User,err error){
	//mambers格式不同,删除project_id和userid字段
	members,err = dao.ListMembersByProjectName(name)
	return
}

func ListBotIdsByProjectId(projectId uint32)(botids *[]model.BotId,err error){
	botids,err = dao.ListBotIdsByProjectId(projectId)
	return
}