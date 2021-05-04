package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"
	"iaso/dao"
	"iaso/mysql"
	"net/http"
	"strconv"
	"xorm.io/xorm"
)

type Technology interface {
	TargetProteinList() gin.HandlerFunc
	TargetProteinCreate() gin.HandlerFunc
	TargetProteinUpdate() gin.HandlerFunc
	TargetProteinDelete() gin.HandlerFunc
	SBDDList() gin.HandlerFunc
	SBDDCreate() gin.HandlerFunc
	SBDDUpdate() gin.HandlerFunc
	SBDDDelete() gin.HandlerFunc
}

type technology struct {
	Logger            *zap.SugaredLogger
	MysqlClient       xorm.Interface
}

func NewTechnology(logger *zap.SugaredLogger) Technology {
	return &technology{
		Logger:        logger.Named("technology"),
		MysqlClient:   mysql.Client,
	}
}

func (t *technology) TargetProteinList() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			technologyData  dao.TechnologyData
			response        dao.TechnologyResponse
		)

		logger := t.Logger.Named("GetTargetProteinPlatformInfo")
		lang := c.Query("lang")

		technologyDataList := make([]dao.TechnologyData, 0)
		sql := fmt.Sprintf("select * from b_target_protein where lang='%s';", lang)
		records, _ := t.MysqlClient.Query(sql)
		for _, record := range records {
			technologyData = dao.TechnologyData{
				Id:          string(record["id"]),
				Content:     string(record["content"]),
			}
			technologyDataList = append(technologyDataList, technologyData)
		}

		logger.Infof("Succeeded to get the record from database.")
		response.Data = technologyDataList
		dao.Success(c, &response, http.StatusOK)
		return
	}
}

func (t *technology) SBDDList() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			technologyData  dao.TechnologyData
			response        dao.TechnologyResponse
		)

		logger := t.Logger.Named("GetSBDDPlatformInfo")
		lang := c.Query("lang")

		technologyDataList := make([]dao.TechnologyData, 0)
		sql := fmt.Sprintf("select * from b_sbdd where lang='%s';", lang)
		records, _ := t.MysqlClient.Query(sql)
		for _, record := range records {
			technologyData = dao.TechnologyData{
				Id:          string(record["id"]),
				Content:     string(record["content"]),
			}
			technologyDataList = append(technologyDataList, technologyData)
		}

		logger.Infof("Succeeded to get the record from database.")
		response.Data = technologyDataList
		dao.Success(c, &response, http.StatusOK)
		return
	}
}

func (t *technology) TargetProteinCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			technologyData  dao.TechnologyData
			response        dao.TechnologyResponse
		)

		logger := t.Logger.Named("CreateTargetProteinPlatformInfo")
		lang := c.Query("lang")

		technologyDataList := make([]dao.TechnologyData, 0)
		sql := fmt.Sprintf("select * from b_target_protein where lang='%s';", lang)
		records, _ := t.MysqlClient.Query(sql)
		for _, record := range records {
			technologyData = dao.TechnologyData{
				Id:          string(record["id"]),
				Content:     string(record["content"]),
			}
			technologyDataList = append(technologyDataList, technologyData)
		}

		if err := c.ShouldBindBodyWith(&technologyData, binding.JSON); err != nil {
			response.Data = technologyDataList
			logger.Errorf("Faliled to bind request: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}

		record := &dao.BTargetProtein{
			Content:     technologyData.Content,
			Lang:        lang,
		}
		_, err := t.MysqlClient.Omit("created_time", "updated_time").InsertOne(record)
		if err != nil {
			response.Data = technologyDataList
			logger.Errorf("Fail to add the targetProtein record: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError,
				fmt.Sprintf("Failed to add the targetProtein record!"))
			return
		}
		technologyData.Id = strconv.FormatInt(record.Id, 10)
		technologyDataList = append(technologyDataList, technologyData)
		response.Data = technologyDataList
		logger.Debugf("Add a record id: %s into database.", record.Id)
		dao.Success(c, &response, http.StatusCreated)
		return
	}
}

func (t *technology) SBDDCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			technologyData  dao.TechnologyData
			response        dao.TechnologyResponse
		)

		logger := t.Logger.Named("CreateSBDDPlatformInfo")
		lang := c.Query("lang")

		technologyDataList := make([]dao.TechnologyData, 0)
		sql := fmt.Sprintf("select * from b_sbdd where lang='%s';", lang)
		records, _ := t.MysqlClient.Query(sql)
		for _, record := range records {
			technologyData = dao.TechnologyData{
				Id:          string(record["id"]),
				Content:     string(record["content"]),
			}
			technologyDataList = append(technologyDataList, technologyData)
		}

		if err := c.ShouldBindBodyWith(&technologyData, binding.JSON); err != nil {
			response.Data = technologyDataList
			logger.Errorf("Faliled to bind request: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}

		record := &dao.BSBDD{
			Content:     technologyData.Content,
			Lang:        lang,
		}
		_, err := t.MysqlClient.Omit("created_time", "updated_time").InsertOne(record)
		if err != nil {
			response.Data = technologyDataList
			logger.Errorf("Fail to add the SBDD record: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError,
				fmt.Sprintf("Failed to add the SBDD record!"))
			return
		}
		technologyData.Id = strconv.FormatInt(record.Id, 10)
		technologyDataList = append(technologyDataList, technologyData)
		response.Data = technologyDataList
		logger.Debugf("Add a record id: %s into database.", record.Id)
		dao.Success(c, &response, http.StatusCreated)
		return
	}
}

func (t *technology) TargetProteinUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			technologyData  dao.TechnologyData
			response        dao.TechnologyResponse
		)

		logger := t.Logger.Named("UpdateTargetProteinPlatformInfo")
		lang := c.Query("lang")

		technologyDataList := make([]dao.TechnologyData, 0)
		sql := fmt.Sprintf("select * from b_target_protein where lang='%s';", lang)
		records, _ := t.MysqlClient.Query(sql)
		for _, record := range records {
			technologyData = dao.TechnologyData{
				Id:          string(record["id"]),
				Content:     string(record["content"]),
			}
			technologyDataList = append(technologyDataList, technologyData)
		}

		if err := c.ShouldBindBodyWith(&technologyData, binding.JSON); err != nil {
			response.Data = technologyDataList
			logger.Errorf("Failed to bind request: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}

		technologyId, _ := strconv.ParseInt(c.Param("id"), 10, 64)

		record := &dao.BTargetProtein{
			Content:      technologyData.Content,
			Lang:         lang,
		}

		_, err := t.MysqlClient.Omit("created_time", "updated_time").Where(
			"id = ?", technologyId).Update(record)
		if err != nil {
			response.Data = technologyDataList
			logger.Errorf("Failed to update the targetProtein record: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}

		for index, data := range technologyDataList {
			if data.Id == c.Param("id") {
				technologyDataList[index] = dao.TechnologyData{
					Id:          data.Id,
					Content:     record.Content,
				}
				break
			}
		}
		response.Data = technologyDataList
		logger.Debugf("Update a record id: %s into database.", record.Id)
		dao.Success(c, &response, http.StatusOK)
		return
	}
}

func (t *technology) SBDDUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			technologyData  dao.TechnologyData
			response        dao.TechnologyResponse
		)

		logger := t.Logger.Named("UpdateSBDDPlatformInfo")
		lang := c.Query("lang")

		technologyDataList := make([]dao.TechnologyData, 0)
		sql := fmt.Sprintf("select * from b_sbdd where lang='%s';", lang)
		records, _ := t.MysqlClient.Query(sql)
		for _, record := range records {
			technologyData = dao.TechnologyData{
				Id:          string(record["id"]),
				Content:     string(record["content"]),
			}
			technologyDataList = append(technologyDataList, technologyData)
		}

		if err := c.ShouldBindBodyWith(&technologyData, binding.JSON); err != nil {
			response.Data = technologyDataList
			logger.Errorf("Failed to bind request: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}

		technologyId, _ := strconv.ParseInt(c.Param("id"), 10, 64)

		record := &dao.BSBDD{
			Content:      technologyData.Content,
			Lang:         lang,
		}

		_, err := t.MysqlClient.Omit("created_time", "updated_time").Where(
			"id = ?", technologyId).Update(record)
		if err != nil {
			response.Data = technologyDataList
			logger.Errorf("Failed to update the SBDD record: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}

		for index, data := range technologyDataList {
			if data.Id == c.Param("id") {
				technologyDataList[index] = dao.TechnologyData{
					Id:          data.Id,
					Content:     record.Content,
				}
				break
			}
		}
		response.Data = technologyDataList
		logger.Debugf("Update a record id: %s into database.", record.Id)
		dao.Success(c, &response, http.StatusOK)
		return
	}
}


func (t *technology) TargetProteinDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			indexDeleted       int
			technologyData     dao.TechnologyData
			response           dao.TechnologyResponse
		)
		logger := t.Logger.Named("DeleteTargetProteinPlatformInfo")
		lang := c.Query("lang")

		technologyDataList := make([]dao.TechnologyData, 0)
		sql := fmt.Sprintf("select * from b_target_protein where lang='%s';", lang)
		records, _ := t.MysqlClient.Query(sql)
		for _, record := range records {
			technologyData = dao.TechnologyData{
				Id:          string(record["id"]),
				Content:     string(record["content"]),
			}
			technologyDataList = append(technologyDataList, technologyData)
		}

		technologyId, _ := strconv.ParseInt(c.Param("id"), 10, 64)

		_, err := t.MysqlClient.Where("id = ?", technologyId).Delete(dao.BTargetProtein{})
		if err != nil {
			response.Data = technologyDataList
			logger.Errorf("Failed to delete the targetProtein record: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}

		for index, data := range technologyDataList {
			if data.Id == c.Param("id") {
				indexDeleted = index
				break
			}
		}

		technologyDataList = append(technologyDataList[:indexDeleted], technologyDataList[indexDeleted + 1:]...)
		response.Data = technologyDataList
		logger.Debugf("Succeeded to delete a record.")
		dao.Success(c, &response, http.StatusOK)
	}
}

func (t *technology) SBDDDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			indexDeleted       int
			technologyData     dao.TechnologyData
			response           dao.TechnologyResponse
		)
		logger := t.Logger.Named("DeleteSBDDPlatformInfo")
		lang := c.Query("lang")

		technologyDataList := make([]dao.TechnologyData, 0)
		sql := fmt.Sprintf("select * from b_sbdd where lang='%s';", lang)
		records, _ := t.MysqlClient.Query(sql)
		for _, record := range records {
			technologyData = dao.TechnologyData{
				Id:          string(record["id"]),
				Content:     string(record["content"]),
			}
			technologyDataList = append(technologyDataList, technologyData)
		}

		technologyId, _ := strconv.ParseInt(c.Param("id"), 10, 64)

		_, err := t.MysqlClient.Where("id = ?", technologyId).Delete(dao.BSBDD{})
		if err != nil {
			response.Data = technologyDataList
			logger.Errorf("Failed to delete the SBDD record: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}

		for index, data := range technologyDataList {
			if data.Id == c.Param("id") {
				indexDeleted = index
				break
			}
		}

		technologyDataList = append(technologyDataList[:indexDeleted], technologyDataList[indexDeleted + 1:]...)
		response.Data = technologyDataList
		logger.Debugf("Succeeded to delete a record.")
		dao.Success(c, &response, http.StatusOK)
	}
}
