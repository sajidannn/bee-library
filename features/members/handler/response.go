package handler

import "bee-library/features/members/entity"

type MemberResponse struct {
	ID      	uint   `json:"id"`
	Name    	string `json:"name"`
	Email   	string `json:"email"`
	Phone   	string `json:"phone"`
	Address 	string `json:"address"`
}

type MemberDetailResponse struct {
	ID      	uint   `json:"id"`
	Name    	string `json:"name"`
	Email   	string `json:"email"`
	Phone   	string `json:"phone"`
	Address 	string `json:"address"`
	Photo   	string `json:"photo"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func ToMemberResponse(member entity.Member) MemberResponse {
	return MemberResponse{
		ID:      member.ID,
		Name:    member.Name,
		Email:   member.Email,
		Phone:   member.Phone,
		Address: member.Address,
	}
}

func ToMemberDetailResponse(member entity.Member) MemberDetailResponse {
	return MemberDetailResponse{
		ID:      member.ID,
		Name:    member.Name,
		Email:   member.Email,
		Phone:   member.Phone,
		Address: member.Address,
		Photo:   member.Photo,
		CreatedAt: member.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: member.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func ToMemberResponseList(members []entity.Member) []MemberResponse {
	var responseList []MemberResponse
	for _, member := range members {
		responseList = append(responseList, MemberResponse{
			ID:      member.ID,
			Name:    member.Name,
			Email:   member.Email,
			Phone:   member.Phone,
			Address: member.Address,
		})
	}
	return responseList
}
