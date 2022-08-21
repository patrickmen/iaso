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
	Image             string    `xorm:"image"`
	Align             string    `xorm:"align"`
	Lang              string    `xorm:"lang"`
	CreatedAt         time.Time `xorm:"created_time"`
	UpdatedAt         time.Time `xorm:"updated_time"`
}

type BTargetValidation struct {
	Id                int64
	Content           string    `xorm:"content"`
	Image             string    `xorm:"image"`
	Align             string    `xorm:"align"`
	Lang              string    `xorm:"lang"`
	CreatedAt         time.Time `xorm:"created_time"`
	UpdatedAt         time.Time `xorm:"updated_time"`
}

type BSbdd struct {
	Id                int64
	Content           string    `xorm:"content"`
	Image             string    `xorm:"image"`
	Align             string    `xorm:"align"`
	Lang              string    `xorm:"lang"`
	CreatedAt         time.Time `xorm:"created_time"`
	UpdatedAt         time.Time `xorm:"updated_time"`
}

type BBiomarker struct {
	Id                int64
	Content           string    `xorm:"content"`
	Image             string    `xorm:"image"`
	Align             string    `xorm:"align"`
	Lang              string    `xorm:"lang"`
	CreatedAt         time.Time `xorm:"created_time"`
	UpdatedAt         time.Time `xorm:"updated_time"`
}

type BProducts struct {
	Id                int64
	Title             string    `xorm:"title"`
	Cover             string  `xorm:"cover"`
	Description       string    `xorm:"description"`
	Content           string    `xorm:"content"`
	Lang              string    `xorm:"lang"`
	CreatedAt         time.Time `xorm:"created_time"`
	UpdatedAt         time.Time `xorm:"updated_time"`
}

type BPipeline struct {
	Id                int64
	Content           string    `xorm:"content"`
	Image             string    `xorm:"image"`
	Align             string    `xorm:"align"`
	Lang              string    `xorm:"lang"`
	CreatedAt         time.Time `xorm:"created_time"`
	UpdatedAt         time.Time `xorm:"updated_time"`
}

type BNews struct {
	Id                int64
	Title             string    `xorm:"title"`
	Description       string    `xorm:"description"`
	Content           string    `xorm:"content"`
	Image             string    `xorm:"image"`
	Align             string    `xorm:"align"`
	Lang              string    `xorm:"lang"`
	CreatedAt         time.Time `xorm:"created_time"`
	UpdatedAt         time.Time `xorm:"updated_time"`
}

type BIndustrialInstitution struct {
	Id                int64
	Content           string    `xorm:"content"`
	Image             string    `xorm:"image"`
	Align             string    `xorm:"align"`
	Lang              string    `xorm:"lang"`
	CreatedAt         time.Time `xorm:"created_time"`
	UpdatedAt         time.Time `xorm:"updated_time"`
}

type BAcademicInstitution struct {
	Id                int64
	Content           string    `xorm:"content"`
	Image             string    `xorm:"image"`
	Align             string    `xorm:"align"`
	Lang              string    `xorm:"lang"`
	CreatedAt         time.Time `xorm:"created_time"`
	UpdatedAt         time.Time `xorm:"updated_time"`
}

type BCareers struct {
	Id                int64
	Content           string    `xorm:"content"`
	Image             string    `xorm:"image"`
	Align             string    `xorm:"align"`
	Lang              string    `xorm:"lang"`
	CreatedAt         time.Time `xorm:"created_time"`
	UpdatedAt         time.Time `xorm:"updated_time"`
}

type BContactUs struct {
	Name              string    `xorm:"name"`
	Phone             string    `xorm:"phone"`
	Email             string    `xorm:"email"`
	Message           string    `xorm:"message"`
}



