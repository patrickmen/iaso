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

type Pipeline interface {
	List() gin.HandlerFunc
	Create() gin.HandlerFunc
	Update() gin.HandlerFunc
	Delete() gin.HandlerFunc
}

type pipeline struct {
	Logger            *zap.SugaredLogger
	MysqlClient       xorm.Interface
}

func NewPipeline(logger *zap.SugaredLogger) AboutUs {
	return &pipeline{
		Logger:       logger.Named("pipeline"),
		MysqlClient:  mysql.Client,
	}
}

func (p *pipeline) List() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			pipelineData   dao.PipelineData
			response       dao.PipelineResponse
		)

		logger := p.Logger.Named("GetPipelineInfo")

		pipelineDataList := make([]dao.PipelineData, 0)
		sql := "select * from b_pipeline;"
		records, _ := p.MysqlClient.Query(sql)
		for _, record := range records {
			pipelineData = dao.PipelineData{
				Id:          string(record["id"]),
				Content:     string(record["content"]),
			}
			pipelineDataList = append(pipelineDataList, pipelineData)
		}

		logger.Infof("Succeeded to get the record from database.")
		response.Data = pipelineDataList
		dao.Success(c, &response, http.StatusOK)
		return
	}
}

func (p *pipeline) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			pipelineData    dao.PipelineData
			response        dao.PipelineResponse
		)

		logger := p.Logger.Named("CreatePipelineInfo")

		pipelineDataList := make([]dao.PipelineData, 0)
		sql := "select * from b_pipeline;"
		records, _ := p.MysqlClient.Query(sql)
		for _, record := range records {
			pipelineData = dao.PipelineData{
				Id:          string(record["id"]),
				Content:     string(record["content"]),
			}
			pipelineDataList = append(pipelineDataList, pipelineData)
		}

		if err := c.ShouldBindBodyWith(&pipelineData, binding.JSON); err != nil {
			response.Data = pipelineDataList
			logger.Errorf("Faliled to bind request: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}

		record := &dao.BAboutUs{
			Content:     pipelineData.Content,
		}
		_, err := p.MysqlClient.Omit("created_time", "updated_time").InsertOne(record)
		if err != nil {
			response.Data = pipelineDataList
			logger.Errorf("Fail to add the pipeline record: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError,
				fmt.Sprintf("Failed to add the pipeline record!"))
			return
		}
		pipelineData.Id = strconv.FormatInt(record.Id, 10)
		pipelineDataList = append(pipelineDataList, pipelineData)
		response.Data = pipelineDataList
		logger.Debugf("Add a record id: %s into database.", record.Id)
		dao.Success(c, &response, http.StatusCreated)
		return
	}
}

func (p *pipeline) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			pipelineData     dao.PipelineData
			response        dao.PipelineResponse
		)

		logger := p.Logger.Named("UpdatePipelineInfo")

		pipelineDataList := make([]dao.PipelineData, 0)
		sql := "select * from b_pipeline;"
		records, _ := p.MysqlClient.Query(sql)
		for _, record := range records {
			pipelineData = dao.PipelineData{
				Id:          string(record["id"]),
				Content:     string(record["content"]),
			}
			pipelineDataList = append(pipelineDataList, pipelineData)
		}

		if err := c.ShouldBindBodyWith(&pipelineData, binding.JSON); err != nil {
			response.Data = pipelineDataList
			logger.Errorf("Failed to bind request: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}

		pipelineId, _ := strconv.ParseInt(c.Param("id"), 10, 64)

		record := &dao.BPipeline{
			Content:      pipelineData.Content,
		}

		_, err := p.MysqlClient.Omit("created_time", "updated_time").Where(
			"id = ?", pipelineId).Update(record)
		if err != nil {
			response.Data = pipelineDataList
			logger.Errorf("Failed to update the pipeline record: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}

		for index, data := range pipelineDataList {
			if data.Id == c.Param("id") {
				pipelineDataList[index] = dao.PipelineData{
					Id:          data.Id,
					Content:     record.Content,
				}
				break
			}
		}
		response.Data = pipelineDataList
		logger.Debugf("Update a record id: %s into database.", record.Id)
		dao.Success(c, &response, http.StatusOK)
		return
	}
}

func (p *pipeline) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			indexDeleted       int
			pipelineData       dao.PipelineData
			response           dao.PipelineResponse
		)
		logger := p.Logger.Named("DeletePipelineDataInfo")

		pipelineDataList := make([]dao.PipelineData, 0)
		sql := "select * from b_pipeline;"
		records, _ := p.MysqlClient.Query(sql)
		for _, record := range records {
			pipelineData = dao.PipelineData{
				Id:          string(record["id"]),
				Content:     string(record["content"]),
			}
			pipelineDataList = append(pipelineDataList, pipelineData)
		}

		pipelineId, _ := strconv.ParseInt(c.Param("id"), 10, 64)

		_, err := p.MysqlClient.Where("id = ?", pipelineId).Delete(dao.BPipeline{})
		if err != nil {
			response.Data = pipelineDataList
			logger.Errorf("Failed to delete the pipeline record: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}


		for index, data := range pipelineDataList {
			if data.Id == c.Param("id") {
				indexDeleted = index
				break
			}
		}

		pipelineDataList = append(pipelineDataList[:indexDeleted], pipelineDataList[indexDeleted + 1:]...)
		response.Data = pipelineDataList
		logger.Debugf("Succeeded to delete a record.")
		dao.Success(c, &response, http.StatusOK)
	}
}