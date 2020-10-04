package model

type RepositoryToMericoProject struct {
	RepositoryId    int64  `gorm:"type:bigint;primary_key;not null"`
	MericoProjectId string `gorm:"type:varchar(50);primary_key;not null"`
}
