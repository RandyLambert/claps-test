package model

import "github.com/jinzhu/gorm"

func init() {
	RegisterMigrateHandler(func(db *gorm.DB) error {

		if err := db.AutoMigrate(&Bot{}).Error; err != nil {
			return err
		}
		return nil
	})
}

type Bot struct {
	Id           string `gorm:"type:varchar(50);primary_key;not null"`
	ProjectId    int64  `gorm:"type:bigint;primary_key;not null"`
	Distribution string `gorm:"type:char;primary_key;not null"`
	SessionId    string `gorm:"type:varchar(50);not null;unique_index:session_id_UNIQUE"`
	Pin          string `gorm:"type:varchar(6);not null"`
	PinToken     string `gorm:"type:varchar(200);not null;unique_index:pin_token_UNIQUE"`
	PrivateKey   string `gorm:"type:text;not null"`
}

type BotDto struct {
	Id           string `json:"id,omitempty" gorm:"type:varchar(50);primary_key;not null"`
	Distribution string `json:"distribution,omitempty" gorm:"type:char;primary_key;not null"`
}

const (
	MericoAlgorithm = "0" //Merico算法
	Commits         = "1" //commit数量
	ChangedLines    = "2" //代码行数
	IdenticalAmount = "3" //平均分配
)


var BOT *Bot

func (bot *Bot) GetBotById(botId string) (botData *Bot, err error) {
	botData = &Bot{}
	err = db.Debug().Where("id=?", botId).Find(botData).Error
	return
}

//根据projectId获取所有的机器人Id
func (bot *Bot) ListBotDtosByProjectId(projectId int64) (botDto *[]BotDto, err error) {
	botDto = &[]BotDto{}
	err = db.Debug().Table("bot").Where("project_id=?", projectId).Scan(botDto).Error
	return
}

func (bot *Bot) GetBotDtoById(botId string) (botDto *BotDto, err error) {
	botDto = &BotDto{}
	err = db.Table("bot").Where("id=?", botId).Scan(botDto).Error
	return
}
