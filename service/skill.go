package service

import (
	"context"
	"myskill-api/datastore"
	"myskill-api/model"
)

type SkillService interface {
	GetSkillByID(ctx context.Context, id string) (*model.Skill, error)
	GetSkillByKeyword(ctx context.Context, keyword string) ([]model.Skill, error)
	UpdateSkillByID(ctx context.Context, skill *model.Skill) error
	InsertSkill(ctx context.Context, skill *model.Skill) error
	GetBulkSkill(ctx context.Context) ([]model.Skill, error)
}

type Skill struct {
	skillDataStore datastore.SkillCollection
}

func (s *Skill) GetBulkSkill(ctx context.Context) ([]model.Skill, error) {
	skills, err := s.skillDataStore.FindAllBulkSkill(ctx)
	if err != nil {
		return nil, err
	}
	return skills, nil
}

func (s *Skill) InsertSkill(ctx context.Context, skill *model.Skill) error {
	if err := s.skillDataStore.Insert(ctx, skill); err != nil {
		return err
	}
	return nil
}

func (s *Skill) UpdateSkillByID(ctx context.Context, skill *model.Skill) error {
	if err := s.skillDataStore.Update(ctx, skill); err != nil {
		return err
	}
	return nil
}

func (s *Skill) GetSkillByKeyword(ctx context.Context, keyword string) ([]model.Skill, error) {
	result, err := s.skillDataStore.FindByKeyword(ctx, keyword)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Skill) GetSkillByID(ctx context.Context, id string) (*model.Skill, error) {
	result, err := s.skillDataStore.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func NewSkillService(skillDataStore datastore.SkillCollection) SkillService {
	return &Skill{
		skillDataStore: skillDataStore,
	}
}
