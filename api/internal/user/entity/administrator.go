package entity

import (
	"time"
)

// Administrator - システム管理者情報
type Administrator struct {
	Admin       `gorm:"-"`
	AdminID     string    `gorm:"primaryKey;<-:create"` // 管理者ID
	PhoneNumber string    `gorm:""`                     // 電話番号
	CreatedAt   time.Time `gorm:"<-:create"`            // 登録日時
	UpdatedAt   time.Time `gorm:""`                     // 更新日時
}

type Administrators []*Administrator

type NewAdministratorParams struct {
	Admin       *Admin
	PhoneNumber string
}

func NewAdministrator(params *NewAdministratorParams) *Administrator {
	return &Administrator{
		AdminID:     params.Admin.ID,
		PhoneNumber: params.PhoneNumber,
		Admin:       *params.Admin,
	}
}

func (a *Administrator) Fill(admin *Admin) {
	a.Admin = *admin
}

func (as Administrators) IDs() []string {
	res := make([]string, len(as))
	for i := range as {
		res[i] = as[i].AdminID
	}
	return res
}

func (as Administrators) Fill(admins map[string]*Admin) {
	for _, a := range as {
		admin, ok := admins[a.AdminID]
		if !ok {
			admin = &Admin{ID: a.AdminID, Role: AdminRoleAdministrator}
		}
		a.Fill(admin)
	}
}
