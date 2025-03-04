package po

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"column:id; not null; primarykey; autoincrement;"`
	UserName string    `gorm:"column:user_name; type:varchar(255); not null; unique"`
	IsActive bool      `gorm:"column:is_active; type:boolean; default:true;"`
	Roles    []Role    `gorm:"many2many:go_user_roles"`
}

func (u *User) TableName() string {
	return "go_db_user"
}
