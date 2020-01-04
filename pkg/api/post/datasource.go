package post

import (
	"context"
	"time"

	"github.com/prongbang/gomulti-source/pkg/db"
	"go.mongodb.org/mongo-driver/bson"
)

type DataSource interface {
	GetList() []Post
	GetById(id string) Post
}

type dataSource struct {
	DbSource db.DataSource
}

func (d *dataSource) GetList() []Post {
	post := []Post{}
	err := d.DbSource.GetMariaDB().Select(&post, "SELECT * FROM post")
	if err != nil {
		return []Post{}
	}
	return post
}

func (d *dataSource) GetById(id string) Post {
	post := Post{}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	collection := d.DbSource.GetMongoDB().Collection("post")
	cur, err := collection.Find(ctx, bson.M{"id": id})
	if err != nil {
		return Post{}
	}
	err = cur.All(ctx, &post)
	defer cur.Close(ctx)
	if err != nil {
		return Post{}
	}
	if err := cur.Err(); err != nil {
		return Post{}
	}
	return post
}

func NewDataSource(dbSource db.DataSource) DataSource {
	return &dataSource{
		DbSource: dbSource,
	}
}
