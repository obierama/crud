package model

import (
	"errors"
	"strings"

	"github.com/jinzhu/gorm"
)

type Profile struct {
	gorm.Model
	Nama          string `gorm:"size:50;not null" json:"nama"`
	Alamat        string `gorm:"size:100;" json:"alamat"`
	Jenis_kelamin string `gorm:"size:9;not null;" json:"jenis_kelamin"`
	Ijazah        string `gorm:"size:100;not null;" json:"ijazah"`
	Nomer_tlp     string `gorm:"size:100;not null;" json:"nomer_tlp"`
	Tgl_lahir     string `gorm:"size:10;not null;"" json:"tgl_lahir"`
}
type StukturProfile struct {
	Id            int64  `json:"id"`
	Nama          string `json:"nama"`
	Alamat        string `json:"alamat"`
	Jenis_kelamin string `json:"nama_lengkap"`
	Ijazah        string `json:"ijazah"`
	Nomer_tlp     string `json:"nomer_tlp"`
	Tgl_lahir     string `json:"tgl_lahir"`
}

func (v *Profile) Update(id uint, db *gorm.DB) (*Profile, error) {
	if err := db.Debug().Table("profiles").Where("id = ?", id).Updates(Profile{
		Nama:          v.Nama,
		Alamat:        v.Alamat,
		Jenis_kelamin: v.Jenis_kelamin,
		Ijazah:        v.Ijazah,
		Nomer_tlp:     v.Nomer_tlp,
		Tgl_lahir:     v.Tgl_lahir,
	}).Error; err != nil {
		return &Profile{}, err
	}
	return v, nil
}

func (u *Profile) Validate(action string) error {
	switch strings.ToLower(action) {

	case "update":

		if u.ID == 0 {
			return errors.New("id is required")
		}
		if u.Nama == "" {
			return errors.New("nama is required")
		}
		if u.Alamat == "" {
			return errors.New("alamat is required")
		}

		if u.Jenis_kelamin == "" {
			return errors.New("jenis_kelamin is required")
		}

		if u.Ijazah == "" {
			return errors.New("ijazah is required")
		}

		if u.Nomer_tlp == "" {
			return errors.New("nomer_tlp is required")
		}
		if u.Tgl_lahir == "" {
			return errors.New("tgl_lahir is DD-MM-YYYY")
		}

		return nil

	default:
		if u.Nama == "" {
			return errors.New("nama is required")
		}
		if u.Alamat == "" {
			return errors.New("alamat is required")
		}

		if u.Jenis_kelamin == "" {
			return errors.New("jenis_kelamin is required")
		}

		if u.Ijazah == "" {
			return errors.New("ijazah is required")
		}

		if u.Nomer_tlp == "" {
			return errors.New("nomer_tlp is required")
		}

		if u.Tgl_lahir == "" {
			return errors.New("tgl_lahir is DD-MM-YYYY")
		}

		return nil

	}

}

func (u *Profile) SaveProfile(db *gorm.DB) (*Profile, error) {
	var err error

	err = db.Debug().Create(&u).Error
	if err != nil {
		return &Profile{}, err
	}
	return u, nil
}

func (u *Profile) GetProfile(db *gorm.DB, parameter string, data string) (*Profile, error) {
	account := &Profile{}
	if err := db.Debug().Table("profiles").Where(parameter, data).First(account).Error; err != nil {
		return nil, err
	}
	return account, nil
}

func (u *Profile) GetProfileInt(db *gorm.DB, parameter string, data int) (*Profile, error) {
	account := &Profile{}
	if err := db.Debug().Table("profiles").Where(parameter, data).First(account).Error; err != nil {
		return nil, err
	}
	return account, nil
}

//untuk fungsi delete/hapus
func (u *Profile) Delete(db *gorm.DB, data string) (*Profile, error) {
	account := &Profile{}
	if err := db.Unscoped().Table("profiles").Where("id = ?", data).Delete(account).Error; err != nil {
		return nil, err
	}
	return account, nil
}
