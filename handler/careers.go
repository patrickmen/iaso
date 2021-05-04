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

type Careers interface {
	List() gin.HandlerFunc
	Create() gin.HandlerFunc
	Update() gin.HandlerFunc
	Delete() gin.HandlerFunc
}

type careers struct {
	Logger            *zap.SugaredLogger
	MysqlClient       xorm.Interface
}

func NewCareers(logger *zap.SugaredLogger) Careers {
	return &careers{
		Logger:        logger.Named("careers"),
		MysqlClient:   mysql.Client,
	}
}

func (cr *careers) List() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			careerData        dao.CareerData
			response          dao.CareersResponse
		)

		logger := cr.Logger.Named("GetCareerInfo")
		lang := c.Query("lang")

		careerDataList := make([]dao.CareerData, 0)
		sql := fmt.Sprintf("select * from b_careers where lang='%s';", lang)
		records, _ := cr.MysqlClient.Query(sql)
		for _, record := range records {
			careerData = dao.CareerData{
				Id:          string(record["id"]),
				Content:     string(record["content"]),
			}
			careerDataList = append(careerDataList, careerData)
		}

		logger.Infof("Succeeded to get the record from database.")
		response.Data = careerDataList
		dao.Success(c, &response, http.StatusOK)
		return
	}
}

func (cr *careers) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			careerData     dao.CareerData
			response       dao.CareersResponse
		)

		logger := cr.Logger.Named("CreateCareerInfo")
		lang := c.Query("lang")

		careerDataList := make([]dao.CareerData, 0)
		sql := fmt.Sprintf("select * from b_careers where lang='%s';", lang)
		records, _ := cr.MysqlClient.Query(sql)
		for _, record := range records {
			careerData = dao.CareerData{
				Id:          string(record["id"]),
				Content:     string(record["content"]),
			}
			careerDataList = append(careerDataList, careerData)
		}

		if err := c.ShouldBindBodyWith(&careerData, binding.JSON); err != nil {
			response.Data = careerDataList
			logger.Errorf("Failed to bind request: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}

		record := &dao.BCareers{
			Content:     careerData.Content,
			Lang:        lang,
		}
		_, err := cr.MysqlClient.Omit("created_time", "updated_time").InsertOne(record)
		if err != nil {
			response.Data = careerDataList
			logger.Errorf("Fail to add the career record: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, fmt.Sprintf("Failed to add the career record!"))
			return
		}
		careerData.Id = strconv.FormatInt(record.Id, 10)
		careerDataList = append(careerDataList, careerData)
		response.Data = careerDataList
		logger.Debugf("Add a record named %s into database.", record.Id)
		dao.Success(c, &response, http.StatusCreated)
		return
	}
}

func (cr *careers) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			careerData     dao.CareerData
			response       dao.CareersResponse
		)

		logger := cr.Logger.Named("UpdateCareerInfo")
		lang := c.Query("lang")

		careerDataList := make([]dao.CareerData, 0)
		sql := fmt.Sprintf("select * from b_careers where lang='%s';", lang)
		records, _ := cr.MysqlClient.Query(sql)
		for _, record := range records {
			careerData = dao.CareerData{
				Id:          string(record["id"]),
				Content:     string(record["content"]),
			}
			careerDataList = append(careerDataList, careerData)
		}

		if err := c.ShouldBindBodyWith(&careerData, binding.JSON); err != nil {
			response.Data = careerDataList
			logger.Errorf("Failed to bind request: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}

		careerId, _ := strconv.ParseInt(c.Param("id"), 10, 64)

		record := &dao.BCareers{
			Content:      careerData.Content,
			Lang:        lang,
		}

		_, err := cr.MysqlClient.Omit("created_time", "updated_time").Where(
			"id = ?", careerId).Update(record)
		if err != nil {
			response.Data = careerDataList
			logger.Errorf("Failed to update the career record: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}

		for index, data := range careerDataList {
			if data.Id == c.Param("id") {
				careerDataList[index] = dao.CareerData{
					Id:          data.Id,
					Content:     record.Content,
				}
				break
			}
		}
		response.Data = careerDataList
		logger.Debugf("Update a record named %s into database.", record.Id)
		dao.Success(c, &response, http.StatusOK)
		return
	}
}

func (cr *careers) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			indexDeleted   int
			careerData     dao.CareerData
			response       dao.CareersResponse
		)
		logger := cr.Logger.Named("DeleteCareerInfo")
		lang := c.Query("lang")

		careerDataList := make([]dao.CareerData, 0)
		sql := fmt.Sprintf("select * from b_careers where lang='%s';", lang)
		records, _ := cr.MysqlClient.Query(sql)
		for _, record := range records {
			careerData = dao.CareerData{
				Id:          string(record["id"]),
				Content:     string(record["content"]),
			}
			careerDataList = append(careerDataList, careerData)
		}

		careerId, _ := strconv.ParseInt(c.Param("id"), 10, 64)

		_, err := cr.MysqlClient.Where("id = ?", careerId).Delete(dao.BCareers{})
		if err != nil {
			response.Data = careerDataList
			logger.Errorf("Failed to delete the career record: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}


		for index, data := range careerDataList {
			if data.Id == c.Param("id") {
				indexDeleted = index
				break
			}
		}

		careerDataList = append(careerDataList[:indexDeleted], careerDataList[indexDeleted + 1:]...)
		response.Data = careerDataList
		logger.Debugf("Succeeded to delete a record.")
		dao.Success(c, &response, http.StatusOK)
	}
}
