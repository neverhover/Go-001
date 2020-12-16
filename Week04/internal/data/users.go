package data

import (
	"github.com/neverhover/Go-001/tree/main/Week04/internal/pkg/err"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Users struct {
	ID          string `gorm:"Column:id;type:varchar;primary_key" json:"id" form:"id"`
	Password    string `gorm:"Column:password;type:varchar;" json:"password" form:"password"`
	Domain      string `gorm:"Column:domain;type:varchar;" json:"domain" form:"domain"`
	NumberAlias string `gorm:"Column:number_alias;type:varchar;" json:"number_alias" form:"number_alias"`
	Mailbox     string `gorm:"Column:mailbox;type:varchar;" json:"mailbox" form:"mailbox"`
	DialString  string `gorm:"Column:dial_string;type:varchar;" json:"dial_string" form:"dial_string"`
	UserContext string `gorm:"Column:user_context;type:varchar;" json:"user_context" form:"user_context"`
}

func (u *Users) Create(db *gorm.DB) (*Users, error) {
	result := db.Create(u)
	return u, result.Error
}

func (u *Users) Get(db *gorm.DB, id string) error {
	result := db.Where("id = ?", id).First(u)
	return result.Error
}

func (u *Users) Delete(db *gorm.DB) error {
	if u.ID == "" {
		return errors.Wrapf(err.ErrInconsistentIDs, "Remove user by id %s", u.ID)
	}

	var result *gorm.DB

	result = db.Unscoped().Delete(u)
	if result.Error != nil {
		return errors.Wrapf(err.ErrInconsistentIDs, "Remove user by id %s", u.ID)
	}
	return nil
}

func (u *Users) Update(db *gorm.DB) (*Users, error) {
	//result := db.Save(p)
	obj := Users{}
	obj.ID = u.ID
	result := db.Updates(u)
	return u, result.Error
}

