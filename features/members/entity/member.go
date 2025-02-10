package entity

import "time"

type Member struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	Photo     string `json:"photo"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type MemberRepository interface {
	GetAll() ([]Member, error)
	GetByID(id uint) (*Member, error)
	Create(member *Member) error
	// login
	Update(id uint, updatedMember *Member) error
	// update password
	Delete(id uint) error
	IsEmailExist(email string) (bool, error)
}

type MemberService interface {
	GetAllMembers() ([]Member, error)
	GetMemberByID(id uint) (*Member, error)
	CreateMember(member *Member) error
	// login
	UpdateMember(id uint, updatedMember *Member) error
	// update password
	DeleteMember(id uint) error
}