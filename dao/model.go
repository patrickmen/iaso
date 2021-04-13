package dao

import (
	"time"
)

type BAuthUsers struct {
	Id                int64
	Username          string `xorm:"username"`
	Password          string `xorm:"password"`
	Phone             string `xorm:"phone"`
	Mail              string `xorm:"mail"`
	Super             bool   `xorm:"is_super"`
	Active            bool   `xorm:"is_active"`
}

type BAboutUs struct {
	Id                int64
	Content           string    `xorm:"content"`
	CreatedAt         time.Time `xorm:"created_time"`
	UpdatedAt         time.Time `xorm:"updated_time"`
}

type BProducts struct {
	Id                int64
	Title             string  `xorm:"title"`
	Cover             string  `xorm:"cover"`
	Description       string  `xorm:"description"`
	Content           string  `xorm:"content"`
	CreatedAt         time.Time `xorm:"created_time"`
	UpdatedAt         time.Time `xorm:"updated_time"`
}

type BResources struct {
	Id                int64
	Title             string  `xorm:"title"`
	Cover             string  `xorm:"cover"`
	Description       string  `xorm:"description"`
	Content           string  `xorm:"content"`
	CreatedAt         time.Time `xorm:"created_time"`
	UpdatedAt         time.Time `xorm:"updated_time"`
}

type BNews struct {
	Id                int64
	Title             string    `xorm:"title"`
	Description       string    `xorm:"description"`
	Content           string    `xorm:"content"`
	CreatedAt         time.Time `xorm:"created_time"`
	UpdatedAt         time.Time `xorm:"updated_time"`
}

type BPartnering struct {
	Id                int64
	Content           string    `xorm:"content"`
	CreatedAt         time.Time `xorm:"created_time"`
	UpdatedAt         time.Time `xorm:"updated_time"`
}

type BCareers struct {
	Id                int64
	Content           string    `xorm:"content"`
	CreatedAt         time.Time `xorm:"created_time"`
	UpdatedAt         time.Time `xorm:"updated_time"`
}

type BContactUs struct {
	Name              string    `xorm:"name"`
	Phone             string    `xorm:"phone"`
	Email             string    `xorm:"email"`
	Message           string    `xorm:"message"`
}



