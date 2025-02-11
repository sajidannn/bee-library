package handler

import (
	"bee-library/features/members/entity"
	"bee-library/helper"
	"bee-library/utils"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type MemberHandler struct {
	service entity.MemberService
}

func NewMemberHandler(service entity.MemberService) *MemberHandler {
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
	if len(members) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "There's no data"})
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
		Data:    ToMemberDetailResponse(*member),
	})
}

func (h *MemberHandler) CreateMember(c *gin.Context) {
	var req MemberCreateRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseError{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	var photoURL string
	file, fileExists := c.Get("photo_file")
	fileName, nameExists := c.Get("photo_fileName")

	if fileExists && nameExists {
		uploadedURL, err := utils.UploadToCloudinary(file.(multipart.File), fileName.(string), "members-photo")
		if err != nil {
			c.JSON(http.StatusInternalServerError, helper.ResponseError{
				Status:  "error",
				Message: "Failed to upload photo",
			})
			return
		}
		photoURL = uploadedURL
	}

	member := entity.Member{
		Name:      req.Name,
		Email:     req.Email,
		Phone:     req.Phone,
		Address:   req.Address,
		Photo:     photoURL,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
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
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseError{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	existingMember, err := h.service.GetMemberByID(uint(id))
	if err != nil {
		c.JSON(helper.MapErrorCode(err), helper.ResponseError{
			Status:  "error",
			Message: "Member not found",
		})
		return
	}

	updatedMember := entity.Member{
		UpdatedAt: time.Now(),
	}
	if req.Name != nil {
		updatedMember.Name = *req.Name
	}
	if req.Phone != nil {
		updatedMember.Phone = *req.Phone
	}
	if req.Address != nil {
		updatedMember.Address = *req.Address
	}

	file, fileExists := c.Get("photo_file")
	fileName, nameExists := c.Get("photo_fileName")

	if fileExists && nameExists {
		if existingMember.Photo != "" {
			err := utils.DeleteFromCloudinary(existingMember.Photo)
			if err != nil {
				c.JSON(http.StatusInternalServerError, helper.ResponseError{
					Status:  "error",
					Message: "Failed to delete old photo",
				})
				return
			}
		}
		uploadedURL, err := utils.UploadToCloudinary(file.(multipart.File), fileName.(string), "members-photo")
		if err != nil {
			c.JSON(http.StatusInternalServerError, helper.ResponseError{
				Status:  "error",
				Message: "Failed to upload photo",
			})
			return
		}
		updatedMember.Photo = uploadedURL
	}
	
	err = h.service.UpdateMember(uint(id), &updatedMember)
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
