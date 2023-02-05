package models

import "gorm.io/gorm"

type RepositoryModel struct {
	gorm.Model
	Uid  string
	Md5  string
	Name string
	Ext  string
	Size int64
	Path string
}
