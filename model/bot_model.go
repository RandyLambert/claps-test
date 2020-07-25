package model

type Bot struct {
	Id           string `gorm:"type:varchar(50);primary_key;unique_index:id_UNIQUE;not null"`
	ProjectId    uint32 `gorm:"type:int unsigned;primary_key;not null"`
	Distribution rune   `gorm:"type:enum('0','1','2','3');primary_key;not null"`
	SessionId    string `gorm:"type:varchar(50);not null;unique_index:session_id_UNIQUE"`
	Pin          string `gorm:"type:varchar(6);not null"`
	PinToken     string `gorm:"type:varchar(200);not null;unique_index:pin_token_UNIQUE"`
	PrivateKey   string `gorm:"type:text;not null"`
}

type BotDto struct {
	Id           string `json:"id" gorm:"type:varchar(50);primary_key;unique_index:id_UNIQUE;not null"`
	Distribution string `json:"distribution" gorm:"type:enum('0','1','2','3');primary_key;not null"`
}

const (
	PersperAlgorithm = "0" //persper算法
	Commits          = "1" //commit数量
	ChangedLines     = "2" //代码行数
	IdenticalAmount  = "3"
)
