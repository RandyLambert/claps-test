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
	err = util.DB.Debug().Table("bot").Where("project_id=?",projectId).Scan(botDto).Error
	return
}

func GetBotDtoById(botId string)(botDto *model.BotDto,err error){
	botDto = &model.BotDto{}
	err = util.DB.Table("bot").Where("bot_id=?",botId).Scan(botDto).Error
	return
}