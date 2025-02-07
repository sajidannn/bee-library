package handler

import "bee-library/features/members/entity"

type MemberResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Photo   string `json:"photo,omitempty"`
}

func ToMemberResponse(member entity.Member) MemberResponse {
	return MemberResponse{
		ID:      member.ID,
		Name:    member.Name,
		Email:   member.Email,
		Phone:   member.Phone,
		Address: member.Address,
		Photo:   member.Photo,
	}
}

func ToMemberResponseList(members []entity.Member) []MemberResponse {
	var responseList []MemberResponse
	for _, member := range members {
		responseList = append(responseList, ToMemberResponse(member))
	}
	return responseList
}
