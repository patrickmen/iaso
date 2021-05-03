package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"
	"xorm.io/xorm"

	"iaso/dao"
	"iaso/mysql"
)

type AboutUs interface {
	List() gin.HandlerFunc
	Create() gin.HandlerFunc
	Update() gin.HandlerFunc
	Delete() gin.HandlerFunc
}

type aboutUs struct {
	Logger            *zap.SugaredLogger
	MysqlClient       xorm.Interface
}

func NewAboutUs(logger *zap.SugaredLogger) AboutUs {
	return &aboutUs{
		Logger:       logger.Named("about-us"),
		MysqlClient:  mysql.Client,
	}
}

func (a *aboutUs) List() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			aboutUsData    dao.AboutUsData
			response       dao.AboutUsResponse
		)

		logger := a.Logger.Named("GetAboutUsInfo")

		aboutUsDataList := make([]dao.AboutUsData, 0)
		sql := "select * from b_about_us;"
		records, _ := a.MysqlClient.Query(sql)
		for _, record := range records {
			aboutUsData = dao.AboutUsData{
				Id:          string(record["id"]),
				Content:     string(record["content"]),
			}
			aboutUsDataList = append(aboutUsDataList, aboutUsData)
		}

		logger.Infof("Succeeded to get the record from database.")
		response.Data = aboutUsDataList
		dao.Success(c, &response, http.StatusOK)
		return
	}
}

func (a *aboutUs) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			aboutUsData     dao.AboutUsData
			response        dao.AboutUsResponse
		)

		logger := a.Logger.Named("CreateAboutUsInfo")

		aboutUsDataList := make([]dao.AboutUsData, 0)
		sql := "select * from b_about_us;"
		records, _ := a.MysqlClient.Query(sql)
		for _, record := range records {
			aboutUsData = dao.AboutUsData{
				Id:          string(record["id"]),
				Content:     string(record["content"]),
			}
			aboutUsDataList = append(aboutUsDataList, aboutUsData)
		}

		if err := c.ShouldBindBodyWith(&aboutUsData, binding.JSON); err != nil {
			response.Data = aboutUsDataList
			logger.Errorf("Faliled to bind request: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}

		record := &dao.BAboutUs{
			Content:     aboutUsData.Content,
		}
		_, err := a.MysqlClient.Omit("created_time", "updated_time").InsertOne(record)
		if err != nil {
			response.Data = aboutUsDataList
			logger.Errorf("Fail to add the aboutUs record: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError,
				fmt.Sprintf("Failed to add the aboutUs record!"))
			return
		}
		aboutUsData.Id = strconv.FormatInt(record.Id, 10)
		aboutUsDataList = append(aboutUsDataList, aboutUsData)
		response.Data = aboutUsDataList
		logger.Debugf("Add a record id: %s into database.", record.Id)
		dao.Success(c, &response, http.StatusCreated)
		return
	}
}

func (a *aboutUs) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			aboutUsData     dao.AboutUsData
			response        dao.AboutUsResponse
		)

		logger := a.Logger.Named("UpdateAboutUsInfo")

		aboutUsDataList := make([]dao.AboutUsData, 0)
		sql := "select * from b_about_us;"
		records, _ := a.MysqlClient.Query(sql)
		for _, record := range records {
			aboutUsData = dao.AboutUsData{
				Id:          string(record["id"]),
				Content:     string(record["content"]),
			}
			aboutUsDataList = append(aboutUsDataList, aboutUsData)
		}

		if err := c.ShouldBindBodyWith(&aboutUsData, binding.JSON); err != nil {
			response.Data = aboutUsDataList
			logger.Errorf("Failed to bind request: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}

		aboutUsId, _ := strconv.ParseInt(c.Param("id"), 10, 64)

		record := &dao.BAboutUs{
			Content:      aboutUsData.Content,
		}

		_, err := a.MysqlClient.Omit("created_time", "updated_time").Where(
			"id = ?", aboutUsId).Update(record)
		if err != nil {
			response.Data = aboutUsDataList
			logger.Errorf("Failed to update the aboutUs record: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}

		for index, data := range aboutUsDataList {
			if data.Id == c.Param("id") {
				aboutUsDataList[index] = dao.AboutUsData{
					Id:          data.Id,
					Content:     record.Content,
				}
				break
			}
		}
		response.Data = aboutUsDataList
		logger.Debugf("Update a record id: %s into database.", record.Id)
		dao.Success(c, &response, http.StatusOK)
		return
	}
}

func (a *aboutUs) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			indexDeleted       int
			aboutUsData        dao.AboutUsData
			response           dao.AboutUsResponse
		)
		logger := a.Logger.Named("DeleteAboutUsInfo")

		aboutUsDataList := make([]dao.AboutUsData, 0)
		sql := "select * from b_about_us;"
		records, _ := a.MysqlClient.Query(sql)
		for _, record := range records {
			aboutUsData = dao.AboutUsData{
				Id:          string(record["id"]),
				Content:     string(record["content"]),
			}
			aboutUsDataList = append(aboutUsDataList, aboutUsData)
		}

		aboutUsId, _ := strconv.ParseInt(c.Param("id"), 10, 64)

		_, err := a.MysqlClient.Where("id = ?", aboutUsId).Delete(dao.BAboutUs{})
		if err != nil {
			response.Data = aboutUsDataList
			logger.Errorf("Failed to delete the aboutUs record: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}


		for index, data := range aboutUsDataList {
			if data.Id == c.Param("id") {
				indexDeleted = index
				break
			}
		}

		aboutUsDataList = append(aboutUsDataList[:indexDeleted], aboutUsDataList[indexDeleted + 1:]...)
		response.Data = aboutUsDataList
		logger.Debugf("Succeeded to delete a record.")
		dao.Success(c, &response, http.StatusOK)
	}
}
