package repository

import (
	"bee-library/features/members/entity"
	"errors"

	"gorm.io/gorm"
)

type MemberRepository interface {
	GetAll() ([]entity.Member, error)
	GetByID(id uint) (*entity.Member, error)
	Create(member *entity.Member) error
	// login
	Update(id uint, updatedFields map[string]interface{}) error
	// update password
	Delete(id uint) error
	IsEmailExist(email string) (bool, error)
}

type memberRepository struct {
	db *gorm.DB
}

func NewMemberRepository(db *gorm.DB) MemberRepository {
	return &memberRepository{db: db}
}

func (r *memberRepository) GetAll() ([]entity.Member, error) {
	var members []entity.Member
	err := r.db.Find(&members).Error
	return members, err
}

func (r *memberRepository) GetByID(id uint) (*entity.Member, error) {
	var member entity.Member
	err := r.db.First(&member, id).Error
	return &member, err
}

func (r *memberRepository) Create(member *entity.Member) error {
	return r.db.Create(member).Error
}

func (r *memberRepository) Update(id uint, updatedFields map[string]interface{}) error {
	return r.db.Model(&entity.Member{}).Where("id = ?", id).Updates(updatedFields).Error
}


func (r *memberRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Member{}, id).Error
}

func (r *memberRepository) IsEmailExist(email string) (bool, error) {
	var member entity.Member
	err := r.db.Where("email = ?", email).First(&member).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

