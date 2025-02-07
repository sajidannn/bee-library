package handler

import (
	"bee-library/features/members/entity"
	"bee-library/features/members/service"
	"bee-library/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MemberHandler struct {
	service service.MemberService
}

func NewMemberHandler(service service.MemberService) *MemberHandler {
	return &MemberHandler{service: service}
}

func (h *MemberHandler) GetAllMembers(c *gin.Context) {
	members, err := h.service.GetAllMembers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseError{
			Status:  "error",
			Message: "Failed to fetch members",
		})
		return
	}
	c.JSON(http.StatusOK, helper.Response{
		Status:  "success",
		Message: "Members retrieved successfully",
		Data:    ToMemberResponseList(members),
	})
}

func (h *MemberHandler) GetMemberByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	member, err := h.service.GetMemberByID(uint(id))
	if err != nil {
		c.JSON(helper.MapErrorCode(err), helper.ResponseError{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, helper.Response{
		Status:  "success",
		Message: "Member retrieved successfully",
		Data:    ToMemberResponse(*member),
	})
}

func (h *MemberHandler) CreateMember(c *gin.Context) {
	var req MemberCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseError{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	member := entity.Member{
		Name:    	req.Name,
		Email:   	req.Email,
		Password: req.Password,
		Phone:   	req.Phone,
		Address: 	req.Address,
		Photo:   	req.Photo,
	}

	err := h.service.CreateMember(&member)
	if err != nil {
		c.JSON(helper.MapErrorCode(err), helper.ResponseError{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, helper.Response{
		Status:  "success",
		Message: "Member created successfully",
		Data:    ToMemberResponse(member),
	})
}

// login

func (h *MemberHandler) UpdateMember(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req MemberUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseError{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	updatedMember := entity.Member{}
	if req.Name != nil {
		updatedMember.Name = *req.Name
	}
	if req.Email != nil {
		updatedMember.Email = *req.Email
	}
	if req.Phone != nil {
		updatedMember.Phone = *req.Phone
	}
	if req.Address != nil {
		updatedMember.Address = *req.Address
	}
	if req.Photo != nil {
		updatedMember.Photo = *req.Photo
	}

	err := h.service.UpdateMember(uint(id), &updatedMember)
	if err != nil {
		c.JSON(helper.MapErrorCode(err), helper.ResponseError{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, helper.Response{
		Status:  "success",
		Message: "Member updated successfully",
		Data:    ToMemberResponse(updatedMember),
	})
}

// update password

func (h *MemberHandler) DeleteMember(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.service.DeleteMember(uint(id))
	if err != nil {
		c.JSON(helper.MapErrorCode(err), helper.ResponseError{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, helper.Response{
		Status:  "success",
		Message: "Member deleted successfully",
	})
}
