package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	USER_ID   uint32    `gorm:"primary_key;auto_increment" json:"user_id"`
	Username  string    `gorm:"size:255;not null;unique" json:"username"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpadtedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func Hash(pass string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
}

func VerifyPassword(hashPass, pass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(pass))
}

func (u *User) BeforeSave() error {
	hashPass, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashPass)
	return nil
}

func (u *User) Prepare() {
	u.USER_ID = 0
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpadtedAt = time.Now()
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Username == "" {
			return errors.New("Required Username")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	case "login":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	default:
		if u.Username == "" {
			return errors.New("Required Username")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) FindAllUser(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

func (u *User) FindByUserID(db *gorm.DB, uid uint32) (*User, error) {
	var err error
	err = db.Debug().Model(User{}).Where("user_id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return u, err
}

func (u *User) UpdateUser(db *gorm.DB, uid uint32) (*User, error) {
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&User{}).Where("user_id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"password":  u.Password,
			"username":  u.Username,
			"email":     u.Email,
			"update_at": time.Now(),
		},
	)

	if db.Error != nil {
		return &User{}, db.Error
	}

	err = db.Debug().Model(&User{}).Where("user_id = ?", uid).Take(&u).Error

	if err != nil {
		return &User{}, err
	}
	return u, nil

}

func (u *User) DeleteUser(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&User{}).Where("user_id = ?", uid).Take(&User{}).Delete(&User{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
