package service

import (
	"goAPISample/app/db"
	"goAPISample/app/model"

	"time"
)

type post model.Post

type PostService struct{}

func (ps PostService) GetPosts() ([]post, error) {

	// DB接続
	db := db.GetDB()

	var p []post

	// DB接続確認
	if err := db.Take(&p).Error; err != nil {
		return nil, err
	}

	db.Find(&p)

	return p, nil
}

func (ps PostService) GetPostById(id int) (post, error) {

	db := db.GetDB()

	var p post

	result := db.First(&p, id)
	if err := result.Error; err != nil {
		return p, err
	}

	return p, nil
}

func (ps PostService) CreatePost(content string) (post, error) {

	db := db.GetDB()

	var p post

	// DB接続確認
	if err := db.Take(&p).Error; err != nil {
		return p, err
	}

	p = post{Content: content, CreatedAt: time.Now()}

	db.Create(&p)

	return p, nil
}

func (ps PostService) UpdatePostById(id int, content string) (post, error) {

	// DB接続
	db := db.GetDB()

	var p post

	result := db.First(&p, id)

	if err := result.Error; err != nil {
		return p, err
	}
	p.Content = content
	db.Save(&p)

	db.Create(&p)

	return p, nil
}

func (ps PostService) DeletePostById(id int) (post, error) {

	// DB接続
	db := db.GetDB()

	var p post

	db.Delete(&p, id)

	return p, nil
}
