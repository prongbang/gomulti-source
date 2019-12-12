package post

import "github.com/prongbang/gomulti-source/pkg/db"

type DataSource interface {
	GetAll() []Post
}

type dataSource struct {
	DbSource db.DataSource
}

func (d *dataSource) GetAll() []Post {
	post := []Post{}
	err := d.DbSource.GetMariaDB().Select(&post, "SELECT * FROM post")
	if err != nil {
		return []Post{}
	}
	return post
}

func NewDataSource(dbSource db.DataSource) DataSource {
	return &dataSource{
		DbSource: dbSource,
	}
}
