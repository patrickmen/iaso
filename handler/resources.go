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

type Resources interface {
	List() gin.HandlerFunc
	Create() gin.HandlerFunc
	Update() gin.HandlerFunc
	Delete() gin.HandlerFunc
}

type resources struct {
	Logger            *zap.SugaredLogger
	MysqlClient       xorm.Interface
}

func NewResources(logger *zap.SugaredLogger) Resources {
	return &resources{
		Logger:        logger.Named("resources"),
		MysqlClient:   mysql.Client,
	}
}

func (r *resources) List() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			resourceData    dao.ResourceData
			response        dao.ResourcesResponse
		)

		logger := r.Logger.Named("GetResourceInfo")
		title := c.Param("title")

		resourceDataList := make([]dao.ResourceData, 0)
		sql := "select * from b_resources;"
		records, _ := r.MysqlClient.Query(sql)
		for _, record := range records {
			resourceData = dao.ResourceData{
				Id:          string(record["id"]),
				Title:       string(record["title"]),
				Cover:       string(record["cover"]),
				Description: string(record["description"]),
				Content:     string(record["content"]),
			}
			if title != "" && resourceData.Title != title {
				continue
			}
			resourceDataList = append(resourceDataList, resourceData)
		}

		resourceDataList = reverseResource(resourceDataList)
		logger.Infof("Succeeded to get the record from database.")
		response.Data = resourceDataList
		dao.Success(c, &response, http.StatusOK)
		return
	}
}

func (r *resources) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			resourceData     dao.ResourceData
			response         dao.ResourcesResponse
		)

		logger := r.Logger.Named("CreateResourceInfo")

		resourceDataList := make([]dao.ResourceData, 0)
		sql := "select * from b_resources;"
		records, _ := r.MysqlClient.Query(sql)
		for _, record := range records {
			resourceData = dao.ResourceData{
				Id:          string(record["id"]),
				Title:       string(record["title"]),
				Cover:       string(record["cover"]),
				Description: string(record["description"]),
				Content:     string(record["content"]),
			}
			resourceDataList = append(resourceDataList, resourceData)
		}

		if err := c.ShouldBindBodyWith(&resourceData, binding.JSON); err != nil {
			resourceDataList = reverseResource(resourceDataList)
			response.Data = resourceDataList
			logger.Errorf("Failed to bind request: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}

		record := &dao.BResources{}
		has, err := r.MysqlClient.Where("title = ?", resourceData.Title).Get(record)
		if err != nil {
			resourceDataList = reverseResource(resourceDataList)
			response.Data = resourceDataList
			logger.Errorf("Failed to add the resource record: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}
		if has {
			resourceDataList = reverseResource(resourceDataList)
			response.Data = resourceDataList
			logger.Errorf("The record existed, submit duplicated!")
			dao.FailWithMessage(c, &response, http.StatusConflict, fmt.Sprintf("The record existed, submit duplicated!"))
			return
		}

		record = &dao.BResources{
			Title:       resourceData.Title,
			Cover:       resourceData.Cover,
			Description: resourceData.Description,
			Content:     resourceData.Content,
		}
		_, err = r.MysqlClient.Omit("created_time", "updated_time").InsertOne(record)
		if err != nil {
			response.Data = resourceDataList
			logger.Errorf("Failed to add the resource record: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, fmt.Sprintf("Failed to add the resource record!"))
			return
		}
		resourceData.Id = strconv.FormatInt(record.Id, 10)
		resourceDataList = append(resourceDataList, resourceData)
		resourceDataList = reverseResource(resourceDataList)
		response.Data = resourceDataList
		logger.Debugf("Add a record named %s into database.", record.Title)
		dao.Success(c, &response, http.StatusCreated)
		return
	}
}

func (r *resources) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			resourceData     dao.ResourceData
			response         dao.ResourcesResponse
		)

		logger := r.Logger.Named("UpdateResourceInfo")

		resourceDataList := make([]dao.ResourceData, 0)
		sql := "select * from b_resources;"
		records, _ := r.MysqlClient.Query(sql)
		for _, record := range records {
			resourceData = dao.ResourceData{
				Id:          string(record["id"]),
				Title:       string(record["title"]),
				Cover:       string(record["cover"]),
				Description: string(record["description"]),
				Content:     string(record["content"]),
			}
			resourceDataList = append(resourceDataList, resourceData)
		}
		resourceDataList = reverseResource(resourceDataList)

		if err := c.ShouldBindBodyWith(&resourceData, binding.JSON); err != nil {
			response.Data = resourceDataList
			logger.Errorf("Failed to bind request: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}

		resourceId, _ := strconv.ParseInt(c.Param("id"), 10, 64)

		record := &dao.BResources{
			Title:        resourceData.Title,
			Cover:        resourceData.Cover,
			Description:  resourceData.Description,
			Content:      resourceData.Content,
		}

		_, err := r.MysqlClient.Omit("created_time", "updated_time").Where(
			"id = ?", resourceId).Update(record)

		if err != nil {
			response.Data = resourceDataList
			logger.Errorf("Failed to update the resource record: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}

		for index, data := range resourceDataList {
			if data.Id == c.Param("id") {
				resourceDataList[index] = dao.ResourceData{
					Id:          data.Id,
					Title:       record.Title,
					Cover:       record.Cover,
					Description: record.Description,
					Content:     record.Content,
				}
				break
			}
		}
		response.Data = resourceDataList
		logger.Debugf("Updated a record named %s into database.", record.Title)
		dao.Success(c, &response, http.StatusOK)
		return
	}
}

func (r *resources) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			indexDeleted       int
			resourceData       dao.ResourceData
			response           dao.ResourcesResponse
		)
		logger := r.Logger.Named("DeleteResourceInfo")

		resourceDataList := make([]dao.ResourceData, 0)
		sql := "select * from b_resources;"
		records, _ := r.MysqlClient.Query(sql)
		for _, record := range records {
			resourceData = dao.ResourceData{
				Id:          string(record["id"]),
				Title:       string(record["title"]),
				Cover:       string(record["cover"]),
				Description: string(record["description"]),
				Content:     string(record["content"]),
			}
			resourceDataList = append(resourceDataList, resourceData)
		}
		resourceDataList = reverseResource(resourceDataList)


		resourceId, _ := strconv.ParseInt(c.Param("id"), 10, 64)

		_, err := r.MysqlClient.Where("id = ?", resourceId).Delete(dao.BResources{})
		if err != nil {
			response.Data = resourceDataList
			logger.Errorf("Failed to delete the resource record: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}


		for index, data := range resourceDataList {
			if data.Id == c.Param("id") {
				indexDeleted = index
				break
			}
		}

		resourceDataList = append(resourceDataList[:indexDeleted], resourceDataList[indexDeleted + 1:]...)
		response.Data = resourceDataList
		logger.Debugf("Succeeded to delete a record.")
		dao.Success(c, &response, http.StatusOK)
	}
}

func reverseResource(arr []dao.ResourceData) []dao.ResourceData{
	length := len(arr)
	for index := range arr[:length / 2] {
		temp := arr[length - 1 - index]
		arr[length - 1 - index] = arr[index]
		arr[index] = temp
	}
	return arr
}
