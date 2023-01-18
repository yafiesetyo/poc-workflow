package repositories

import (
	"github.com/yafiesetyo/poc-workflow/app/model"
	"gorm.io/gorm"
)

type KtpRepoImpl struct {
	DB *gorm.DB
}

func NewKtpRepo(db *gorm.DB) KtpRepoImpl {
	return KtpRepoImpl{
		DB: db,
	}
}

func (r *KtpRepoImpl) Register(req model.KTP) error {
	return r.DB.Table("ktp_register").Create(&req).Error
}

func (r *KtpRepoImpl) FindByID(id uint64) (res model.KTP, err error) {
	return res, r.DB.Table("ktp_register").
		Where(`"deletedAt" isnull and "id" = ?`, id).
		Scan(&res).Error
}

func (r *KtpRepoImpl) Updates(id uint64, field map[string]interface{}) error {
	return r.DB.Table("ktp_register").
		Where(`"deletedAt" isnull and id=?`, id).
		Updates(field).Error
}
