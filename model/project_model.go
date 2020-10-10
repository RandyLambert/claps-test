package model

import (
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
	"time"
)

func init() {
	RegisterMigrateHandler(func(db *gorm.DB) error {

		if err := db.AutoMigrate(&Project{}).Error; err != nil {
			return err
		}
		return nil
	})
}

type Project struct {
	Id          int64           `json:"id,omitempty" gorm:"type:bigserial;auto_increment;primary_key"`
	Name        string          `json:"name,omitempty" gorm:"type:varchar(50);not null"`
	DisplayName string          `json:"display_name,omitempty" gorm:"type:varchar(50);default:null"`
	Description string          `json:"description,omitempty" gorm:"type:varchar(120);default:null"`
	AvatarUrl   string          `json:"avatar_url,omitempty" gorm:"type:varchar(100);default:null"`
	Donations   int64           `json:"donations,omitempty" gorm:"type:bigint;default:0"`
	Total       decimal.Decimal `json:"total,omitempty" gorm:"type:varchar(128);default:null"`
	CreatedAt   time.Time       `json:"createdAt,omitempty" gorm:"type:timestamp with time zone"`
	UpdatedAt   time.Time       `json:"updatedAt,omitempty" gorm:"type:timestamp with time zone"`
}

type ProjectTotal struct {
	Id        int64           `json:"id,omitempty" gorm:"type:bigserial;auto_increment;primary_key"`
	Donations int64           `json:"donations,omitempty" gorm:"type:bigint;default:0"`
	Total     decimal.Decimal `json:"total,omitempty" gorm:"type:varchar(128);default:null"`
}

type Badge struct {
	Code    string `form:"code" json:"code" binding:"required"`
	Color   string `form:"color" json:"color" binding:"required"`
	BgColor string `form:"bg_color" json:"bg_color" binding:"required"`
	Size    string `form:"size" json:"size" binding:"required"`
}
var PROJECT *Project
//获取所有项目
func (proj *Project) ListProjectsAll() (projects *[]Project, err error) {

	projects = &[]Project{}
	err = db.Debug().Find(projects).Error
	return
}


func (proj *Project) getProjectsNumbers() (number int,err error) {

	err = db.Debug().Table("project").Count(&number).Error
	return
}

func (proj * Project) ListProjectsByQuery(q *PaginationQ) (projects *[]Project,number int,err error) {
	projects = &[]Project{}
	number,err = proj.getProjectsNumbers()
	if err != nil {
		return
	}

	tx := db.Debug().Table("project")
	if q.Limit < 0{
		q.Limit = 20
	}

	if q.Offset < 0{
		q.Offset = 0
	}

	if q.Q != ""{
		tx = tx.Where("name Like ?","%" + q.Q + "%")
	}

	err = tx.Limit(q.Limit).Offset(q.Offset).Find(projects).Error
	return
}

//通过项目id获取项目
func (proj *Project) GetProjectById(projectId int64) (project *Project, err error) {

	project = &Project{}
	err = db.Debug().Where("id=?", projectId).Find(project).Error
	return
}

//根据userId获取所有项目
func (proj *Project) ListProjectsByUserId(userId int64) (projects *[]Project, err error) {
	projects = &[]Project{}
	err = db.Debug().Where("id IN(?)",
		db.Debug().Table("member").Select("project_id").Where("user_id=?", userId).SubQuery()).Find(projects).Error
	return
}

func (proj *Project) GetProjectTotalByBotId(BotId string) (projectTotal *ProjectTotal, err error) {
	projectTotal = &ProjectTotal{}
	err = db.Debug().Table("project").Select("id,donations,total").Where("id=?",
		db.Debug().Table("bot").Select("project_id").Where("id=?", BotId).SubQuery()).Scan(projectTotal).Error
	return
}

func (proj *Project) UpdateProjectTotal(projectTotal *ProjectTotal) (err error) {
	err = db.Debug().Table("project").Save(projectTotal).Error
	return
}

func (proj *Project) SumProjectDonationsByUserId(userId int64) (donations int64, err error) {
	type Result struct {
		Total int64
	}
	var result Result
	err = db.Debug().Table("project").Select("sum(donations) as total").Where("id IN(?)",
		db.Debug().Table("member").Select("project_id").Where("user_id=?", userId).SubQuery()).Scan(&result).Error
	donations = result.Total
	return
}

func (proj *Project) GetProjectsTotal()(total uint32,err error) {
	return
}