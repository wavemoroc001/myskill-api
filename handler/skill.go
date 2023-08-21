package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"myskill-api/model"
	"myskill-api/service"
	"net/http"
)

type SkillHandler interface {
	GetSkillByID(gctx *gin.Context)
	GetBulkSkill(gctx *gin.Context)
	GetSkillByKeyword(gctx *gin.Context)
	UpdateSkillByID(gctx *gin.Context)
	InsertSkill(gctx *gin.Context)
}

type Skill struct {
	service service.SkillService
}

func (s *Skill) InsertSkill(gctx *gin.Context) {
	var req AddSkillRequest
	if err := gctx.ShouldBindJSON(&req); err != nil {
		gctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	skill := &model.Skill{
		Name:        req.Name,
		Description: req.Description,
		Logo:        req.Logo,
		Kind:        req.Kind,
	}

	err := s.service.InsertSkill(gctx, skill)
	if err != nil {
		gctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	gctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (s *Skill) UpdateSkillByID(gctx *gin.Context) {
	skillID := gctx.Param("id")
	var req UpdateSkillRequest
	if err := gctx.ShouldBindJSON(&req); err != nil {
		gctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	skill, err := s.service.GetSkillByID(gctx, skillID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			gctx.JSON(http.StatusNotFound, gin.H{
				"message": fmt.Errorf("skill not found: %w", skillID),
			})
		}
		gctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	if req.Name != "" {
		skill.Name = req.Name
	}

	if req.Description != "" {
		skill.Description = req.Description
	}

	if req.Logo != "" {
		skill.Logo = req.Logo
	}

	if err := s.service.UpdateSkillByID(gctx, skill); err != nil {
		gctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	gctx.JSON(http.StatusAccepted, gin.H{"message": "skill updated"})
}

func (s *Skill) GetSkillByKeyword(gctx *gin.Context) {
	keyword := gctx.Query("keyword")
	skills, err := s.service.GetSkillByKeyword(gctx, keyword)
	if err != nil {
		gctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	gctx.JSON(http.StatusOK, skills)
}

func (s *Skill) GetSkillByID(gctx *gin.Context) {
	id := gctx.Param("id")
	skill, err := s.service.GetSkillByID(gctx, id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			gctx.JSON(404, gin.H{
				"message": err.Error(),
			})
			return
		}
		gctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	gctx.JSON(200, skill)
}

func (s *Skill) GetBulkSkill(gctx *gin.Context) {
	skills, err := s.service.GetBulkSkill(gctx)
	if err != nil {
		gctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	gctx.JSON(http.StatusOK, skills)
}

func NewSkillHandler(service service.SkillService) SkillHandler {
	return &Skill{
		service: service,
	}
}
