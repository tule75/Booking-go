package po

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID       uuid.UUID `gorm:"column:id; not null; primarykey, autoincrement;"`
	RoleName string    `gorm:"column:role_name; type:varchar(255); not null;"`
	RoleNote string    `gorm:"column:role_note; type:text; not null;"`
}

func (u *Role) TableName() string {
	return "go_db_role"
}
