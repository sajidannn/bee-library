package service

import (
	"bee-library/features/members/entity"
	"bee-library/features/members/repository"
	"bee-library/helper"
)

type MemberService interface {
	GetAllMembers() ([]entity.Member, error)
	GetMemberByID(id uint) (*entity.Member, error)
	CreateMember(member *entity.Member) error
	// login
	UpdateMember(id uint, updatedMember *entity.Member) error
	// update password
	DeleteMember(id uint) error
}

type memberService struct {
	repo repository.MemberRepository
}

func NewMemberService(repo repository.MemberRepository) MemberService {
	return &memberService{repo: repo}
}

func (s *memberService) GetAllMembers() ([]entity.Member, error) {
	return s.repo.GetAll()
}

func (s *memberService) GetMemberByID(id uint) (*entity.Member, error) {
	member, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return member, nil
}

func (s *memberService) CreateMember(newMember *entity.Member) error {
	exists, err := s.repo.IsEmailExist(newMember.Email)
	if err != nil {
		return helper.ErrInternalServer
	}
	if exists {
		return helper.ErrEmailExists
	}

	newMember.Password, _ = helper.HashPassword(newMember.Password)
	if err := s.repo.Create(newMember); err != nil {
		return helper.ErrInternalServer
	}

	return nil
}

func (s *memberService) UpdateMember(id uint, updatedMember *entity.Member) error {
	_, err := s.GetMemberByID(id)
	if err != nil {
		return err
	}
	if err := s.repo.Update(id, updatedMember); err != nil {
		return helper.ErrInternalServer
	}
	return nil
}


func (s *memberService) DeleteMember(id uint) error {
	_, err := s.GetMemberByID(id)
	if err != nil {
		return err
	}
	if err := s.repo.Delete(id); err != nil {
		return helper.ErrInternalServer
	}
	return nil
}