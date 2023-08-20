package skill

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type DataStore interface {
	FindByName(ctx context.Context, name string) (*Skill, error)
	FindByID(ctx context.Context, skillID string) (*Skill, error)
	FindAll(ctx context.Context) ([]Skill, error)
	FindAllByPageAndSort(ctx context.Context, pageNo int, pageSize int) (Page, error)
	UpdateOne(ctx context.Context, skill *Skill) error
}

type Page struct {
	Data       any `json:"data"`
	PageNo     int `json:"pageNo"`
	PageSize   int `json:"pageSize"`
	TotalPage  int `json:"totalPage"`
	TotalCount int `json:"totalCount"`
}

func NewPage(data any, no int, size int, totalCount int) Page {
	return Page{
		Data:       data,
		PageNo:     no,
		PageSize:   size,
		TotalCount: totalCount,
		TotalPage:  TotalPage(totalCount, size),
	}
}

func TotalPage(totalCount int, size int) int {
	carry := 0
	if totalCount%size != 0 {
		carry = 1
	}
	pageCal := totalCount / size
	return pageCal + carry

}

type Collection struct {
	*mongo.Collection
}

func (c *Collection) FindByName(ctx context.Context, name string) (*Skill, error) {
	return nil, nil
}

func (c *Collection) FindByID(ctx context.Context, skillID string) (*Skill, error) {
	return nil, nil
}

func (c *Collection) FindAllByPageAndSort(ctx context.Context, pageNo int, pageSize int) (Page, error) {
	var limit int64
	if pageNo < 1 {
		limit = int64(pageSize)
	} else {
		limit = int64(pageSize * pageNo)
	}

	type Data struct {
		ID primitive.ObjectID `bson:"_id"`
	}
	cursor, err := c.Find(ctx, bson.D{}, options.Find().SetProjection(bson.M{"_id": int32(1)}).SetLimit(limit))
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

	totalSkill, err := c.CountDocuments(ctx, bson.D{})
	if err != nil {
		log.Println(fmt.Errorf("can not count skill: %w", err))
	}

	lastedID := datas[len(datas)-1].ID
	query := bson.M{"_id": bson.M{"$gt": lastedID}}
	cursor, err = c.Find(ctx, query, options.Find().SetLimit(int64(pageSize)))
	if err != nil {
		log.Println(fmt.Errorf("can not find skill: %w", err))
	}

	var result []Skill
	for cursor.Next(ctx) {
		var skill Skill
		if err := cursor.Decode(&result); err != nil {
			log.Println(fmt.Errorf("can not decode skill: %w", err))
		}
		result = append(result, skill)
	}

	page := NewPage(result, pageNo, pageSize, int(totalSkill))

	return page, nil
}

func (c *Collection) FindAll(ctx context.Context) ([]Skill, error) {
	return nil, nil
}

func (c *Collection) InsertOne(ctx context.Context, skill *Skill) error {
	return nil
}

func (c *Collection) UpdateOne(ctx context.Context, skill *Skill) error {
	return nil
}

func NewCollection(db *mongo.Database) DataStore {
	col := db.Collection("skills")
	return &Collection{col}
}
