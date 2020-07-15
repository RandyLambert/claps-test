package service

import (
	"claps-test/dao"
	"claps-test/model"
	"github.com/fox-one/mixin-sdk-go"
	"github.com/shopspring/decimal"
)

//通过projectName查询,查询某个项目的详情
func GetProject(name string) (projectDetailInfo *map[string]interface{},err error){

	project,err := dao.GetProject(name)
	if err != nil {
		return
	}

	projectInfo,err := GetProjectInfo(project)
	if err != nil {
		return
	}
	repositories,err := GetProjectRepositories(project.Id)
	if err!=nil {
		return
	}
	members,err := GetProjectMembers(name)
	if err!=nil {
		return
	}
	botIds,err := GetProjectBotIds(project.Id)
	if err != nil {
		return
	}
	projectDetailInfo = &map[string]interface{}{
		"project":projectInfo,
		"repositories":repositories,
		"members":members,
		"botIds":botIds,
	}
	return
}

//获取数据库中所有project
func GetProjects() (projectInfos []*map[string]interface{},err error){
	projects,err := dao.GetProjects()
	if err != nil{
		return
	}

	for i :=range *projects {
		var projectInfo *map[string]interface{}
		projectInfo,err = GetProjectInfo(&(*projects)[i])
		if err != nil {
			return
		}
		projectInfos = append(projectInfos,projectInfo)
	}
	return
}


//查询某用户的所有项目,获取数据库中所有project
func GetProjectsByUserId(userId int64) (){
	/*
	projects,err := dao.GetProjectsByUserId(userId)
	if err != nil{
		return
	}

	for i :=range *projects {
		var projectInfo *map[string]interface{}
		//projectInfo,err = GetProjectInfo(&(*projects)[i])
		if err != nil {
			return
		}
		projectInfos = append(projectInfos,projectInfo)
	}
	 */
	return
}

func GetProjectInfo(project *model.Project)(projectInfo *map[string]interface{},err error){

	projectTotal,err := dao.GetProjectTotal(project.Id)
	if err != nil {
		return
	}
	total := decimal.Zero
	for i := range *projectTotal {
		var assert *mixin.Asset
		assert,err = GetAsset((*projectTotal)[i].AssetId)
		if err != nil {
			return
		}
		assertTotal := decimal.NewFromFloat((*projectTotal)[i].Total)
		total.Add(assert.PriceUSD.Mul(assertTotal))
	}

	patrons,err := dao.GetProjectPatrons(project.Id)
	projectInfo = &map[string]interface{}{
		"patrons":patrons,
		"total":total,
		"project":project,
	}
	return
}

func GetProjectTransactions(name string,assetId string)(transactions *[]model.Transaction,err error){
	transactions,err = dao.GetProjectTransactions(name,assetId)
	return
}

func GetProjectRepositories(projectId uint32)(repositoriesInfo []*map[string]interface{},err error){
	repositories,err := dao.GetProjectRepositories(projectId)
	if err!=nil{
		return
	}
	for i := range *repositories{
		//todo github star
		var stars uint32
		//(*repositories)[i].Id
		//model.Repository{
		//	Id:          0,
		//	ProjectId:   0,
		//	Type:        "",
		//	Name:        "",
		//	Slug:        "",
		//	Description: "",
		//	CreatedAt:   time.Time{},
		//	UpdatedAt:   time.Time{},
		//}
		repositoriesInfo = append(repositoriesInfo,&map[string]interface{}{
			"stars":stars,
			"repository":(*repositories)[i],
			//"id":(*repositories)[i].Id,
			//"project_id":(*repositories)[i].ProjectId,
			//"type":(*repositories)[i].Type,
			//"project_id":(*repositories)[i].Id,
			//"project_id":(*repositories)[i].Id,
			//"project_id":(*repositories)[i].Id,
		})
	}

	return
}

func GetProjectMembers(name string)(members *[]model.User,err error){
	//mambers格式不同,删除project_id和userid字段
	members,err = dao.GetProjectMembers(name)
	return
}

func GetProjectBotIds(projectId uint32)(botids *[]model.BotId,err error){
	botids,err = dao.GetProjectBotIds(projectId)
	return
}