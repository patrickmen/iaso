package dao

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Interface interface {
	Get()       *Base
	Set(Base)
}

type Base struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Success bool   `json:"success"`
}

func (b *Base) Get() *Base {
	return b
}

func (b *Base) Set(v Base) {
	*b = v
}

type AboutUsResponse struct {
	Base
	Data      []AboutUsData           `json:"data"`
}

type PartneringResponse struct {
	Base
	Data      []PartneringData        `json:"data"`
}

type UserResponse struct {
	Base
	Id        int64
	Data      UserData        `json:"data"`
}

type ContactUsResponse struct {
	Base
	Id        int64
	Data      ContactUsData   `json:"data"`
}

type PipelineResponse struct {
	Base
	Data      []PipelineData          `json:"data"`
}

type NewsResponse struct {
	Base
	Data      []NewsData              `json:"data"`
}

type ProductsResponse struct {
	Base
	Data      []ProductData           `json:"data"`
}

type TechnologyResponse struct {
	Base
	Data      []TechnologyData        `json:"data"`
}

type CareersResponse struct {
	Base
	Data      []CareerData            `json:"data"`
}

type UserData struct {
	User       string          `json:"user"`
	Token      string          `json:"token"`
	ExpireAt   time.Time       `json:"expireAt"`
}

type AboutUsData struct {
	Id                string   `json:"id"`
	Content           string   `json:"content"`
	CreatedAt         string   `json:"createdAt"`
	UpdatedAt         string   `json:"updatedAt"`
}

type TechnologyData struct {
	Id                string   `json:"id"`
	Content           string   `json:"content"`
	CreatedAt         string   `json:"createdAt"`
	UpdatedAt         string   `json:"updatedAt"`
}

type PartneringData struct{
	Id                string   `json:"id"`
	Content           string   `json:"content"`
	CreatedAt         string   `json:"createdAt"`
	UpdatedAt         string   `json:"updatedAt"`
}

type PipelineData struct {
	Id                string  `json:"id"`
	Content           string  `json:"content"`
	CreatedAt         string  `json:"createdAt"`
	UpdatedAt         string  `json:"updatedAt"`
}

type NewsData struct {
	Id                string  `json:"id"`
	Title             string  `json:"title"`
	Description       string  `json:"description"`
	Content           string  `json:"content"`
	CreatedAt         string  `json:"createdAt"`
	UpdatedAt         string  `json:"updatedAt"`
}

type ProductData struct {
	Id                string  `json:"id"`
	Title             string  `json:"title"`
	Cover             string  `json:"cover"`
	Description       string  `json:"description"`
	Content           string  `json:"content"`
	CreatedAt         string  `json:"createdAt"`
	UpdatedAt         string  `json:"updatedAt"`
}

type CareerData struct{
	Id                string  `json:"id"`
	Content           string  `json:"content"`
	CreatedAt         string  `json:"createdAt"`
	UpdatedAt         string  `json:"updatedAt"`
}

type ContactUsData struct {
	Token      string         `json:"token"`
	ExpireAt   time.Time      `json:"expireAt"`
}

func Success(g *gin.Context, i Interface, code int) {
	i.Set(Base{Code: code, Message: http.StatusText(code), Success: true})
	g.JSON(code, i)
}

func FailWithMessage(g *gin.Context, i Interface, code int, format string, args ...interface{}) {
	i.Set(Base{Code: code, Message: fmt.Sprintf(format, args...), Success: false})
	g.JSON(code, i)
}

