package dao

import (
	"claps-test/util"
	"claps-test/model"
)

func GetBotById(botId string)(bot *model.Bot,err error){
	bot = &model.Bot{}
	err = util.DB.Debug().Where("id=?",botId).Find(bot).Error
	return
}

//根据projectId获取所有的机器人Id
func ListBotDtosByProjectId(projectId uint32)(botDto *[]model.BotDto,err error){
	botDto = &[]model.BotDto{}
	err = util.DB.Debug().Table("bot").Where("project_id=?",projectId).Find(botDto).Error
	return
}