package datastore

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"myskill-api/model"
)

type SkillCollection interface {
	FindByKeyword(ctx context.Context, keyword string) ([]model.Skill, error)
	FindByID(ctx context.Context, skillID string) (*model.Skill, error)
	FindAll(ctx context.Context) ([]model.Skill, error)
	FindAllByPageAndSort(ctx context.Context, pageNo int, pageSize int) (Page, error)
	Update(ctx context.Context, skill *model.Skill) error
	Insert(ctx context.Context, skill *model.Skill) error
}

type Skill struct {
	*mongo.Collection
}

func (s *Skill) FindByKeyword(ctx context.Context, keyword string) ([]model.Skill, error) {
	query := bson.M{
		"name": primitive.Regex{
			Pattern: fmt.Sprintf(".*%v.*", keyword),
			Options: "i",
		},
	}

	cursor, err := s.Find(ctx, query)
	if err != nil {
		return []model.Skill{}, err
	}

	skills := make([]model.Skill, 0)
	for cursor.Next(ctx) {
		var skill model.Skill
		if err := cursor.Decode(&skill); err != nil {
			return nil, err
		}
		skills = append(skills, skill)
	}

	return skills, nil
}

func (s *Skill) FindByID(ctx context.Context, skillID string) (*model.Skill, error) {
	objectID, err := primitive.ObjectIDFromHex(skillID)
	if err != nil {
		return nil, fmt.Errorf("can not convert skillID to objectID: %w", err)
	}

	var skill *model.Skill
	if err := s.FindOne(ctx, bson.M{"_id": objectID}).Decode(&skill); err != nil {
		return nil, err
	}
	return skill, nil
}

func (s *Skill) FindAllByPageAndSort(ctx context.Context, pageNo int, pageSize int) (Page, error) {
	var limit int64
	if pageNo < 1 {
		limit = int64(pageSize)
	} else {
		limit = int64(pageSize * pageNo)
	}

	type Data struct {
		ID primitive.ObjectID `bson:"_id"`
	}
	cursor, err := s.Find(ctx, bson.D{}, options.Find().SetProjection(bson.M{"_id": int32(1)}).SetLimit(limit))
	if err != nil {
		log.Println(fmt.Errorf("can not find lastID skill: %w", err))
	}

	// get lastID
	var datas []Data
	for cursor.Next(ctx) {
		var data Data
		if err := cursor.Decode(&data); err != nil {
			log.Println(fmt.Errorf("can not decode skill: %w", err))
		}
		log.Printf("%#v", data.ID.String())
		datas = append(datas, data)
	}

	totalSkill, err := s.CountDocuments(ctx, bson.D{})
	if err != nil {
		log.Println(fmt.Errorf("can not count skill: %w", err))
	}

	lastedID := datas[len(datas)-1].ID
	query := bson.M{"_id": bson.M{"$gt": lastedID}}
	cursor, err = s.Find(ctx, query, options.Find().SetLimit(int64(pageSize)))
	if err != nil {
		log.Println(fmt.Errorf("can not find skill: %w", err))
	}

	var result []model.Skill
	for cursor.Next(ctx) {
		var skill model.Skill
		if err := cursor.Decode(&result); err != nil {
			log.Println(fmt.Errorf("can not decode skill: %w", err))
		}
		result = append(result, skill)
	}

	page := NewPage(result, pageNo, pageSize, int(totalSkill))

	return page, nil
}

func (s *Skill) FindAll(ctx context.Context) ([]model.Skill, error) {
	var result []model.Skill
	cur, err := s.Find(ctx, bson.D{})
	if err != nil {
		return result, err
	}
	for cur.Next(ctx) {
		var skill model.Skill
		if err := cur.Decode(&skill); err != nil {
			return result, err
		}
		result = append(result, skill)
	}

	return result, nil
}

func (s *Skill) Insert(ctx context.Context, skill *model.Skill) error {
	_, err := s.InsertOne(ctx, skill)
	if err != nil {
		return err
	}
	return nil
}

func (s *Skill) Update(ctx context.Context, skill *model.Skill) error {
	filter := bson.D{
		{"_id", skill.ID},
	}

	field := bson.A{}

	if skill.Name != "" {
		field = append(field, bson.D{{"name", skill.Name}})
	}
	if skill.Description != "" {
		field = append(field, bson.D{{"description", skill.Description}})
	}
	if skill.Logo != "" {
		field = append(field, bson.D{{"logo", skill.Logo}})
	}

	updateCMD := bson.D{
		{"$set", field},
	}

	_, err := s.UpdateOne(ctx, filter, updateCMD)
	if err != nil {
		return err
	}
	return nil
}

func NewSkillCollection(db *mongo.Database) SkillCollection {
	col := db.Collection("skills")
	return &Skill{col}
}
