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

type Technology interface {
	TargetValidationList() gin.HandlerFunc
	TargetValidationCreate() gin.HandlerFunc
	TargetValidationUpdate() gin.HandlerFunc
	TargetValidationDelete() gin.HandlerFunc
	SBDDList()   gin.HandlerFunc
	SBDDCreate() gin.HandlerFunc
	SBDDUpdate() gin.HandlerFunc
	SBDDDelete() gin.HandlerFunc
	BiomarkerList()   gin.HandlerFunc
	BiomarkerCreate() gin.HandlerFunc
	BiomarkerUpdate() gin.HandlerFunc
	BiomarkerDelete() gin.HandlerFunc
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

func (t *technology) TargetValidationList() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			technologyData  dao.TechnologyData
			response        dao.TechnologyResponse
		)

		logger := t.Logger.Named("GetTargetValidationPlatformInfo")
		lang := c.Query("lang")

		technologyDataList := make([]dao.TechnologyData, 0)
		sql := fmt.Sprintf("select * from b_target_validation where lang='%s';", lang)
		records, _ := t.MysqlClient.Query(sql)
		for _, record := range records {
			technologyData = dao.TechnologyData{
				Id:          string(record["id"]),
				Content:     string(record["content"]),
				Image:       string(record["image"]),
				Align:       string(record["align"]),
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
				Image:       string(record["image"]),
				Align:       string(record["align"]),
			}
			technologyDataList = append(technologyDataList, technologyData)
		}

		logger.Infof("Succeeded to get the record from database.")
		response.Data = technologyDataList
		dao.Success(c, &response, http.StatusOK)
		return
	}
}

func (t *technology) BiomarkerList() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			technologyData  dao.TechnologyData
			response        dao.TechnologyResponse
		)

		logger := t.Logger.Named("GetBiomarkerPlatformInfo")
		lang := c.Query("lang")

		technologyDataList := make([]dao.TechnologyData, 0)
		sql := fmt.Sprintf("select * from b_biomarker where lang='%s';", lang)
		records, _ := t.MysqlClient.Query(sql)
		for _, record := range records {
			technologyData = dao.TechnologyData{
				Id:          string(record["id"]),
				Content:     string(record["content"]),
				Image:       string(record["image"]),
				Align:       string(record["align"]),
			}
			technologyDataList = append(technologyDataList, technologyData)
		}

		logger.Infof("Succeeded to get the record from database.")
		response.Data = technologyDataList
		dao.Success(c, &response, http.StatusOK)
		return
	}
}

func (t *technology) TargetValidationCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			technologyData  dao.TechnologyData
			response        dao.TechnologyResponse
		)

		logger := t.Logger.Named("CreateTargetValidationPlatformInfo")
		lang := c.Query("lang")

		technologyDataList := make([]dao.TechnologyData, 0)
		sql := fmt.Sprintf("select * from b_target_validation where lang='%s';", lang)
		records, _ := t.MysqlClient.Query(sql)
		for _, record := range records {
			technologyData = dao.TechnologyData{
				Id:          string(record["id"]),
				Content:     string(record["content"]),
				Image:       string(record["image"]),
				Align:       string(record["align"]),
			}
			technologyDataList = append(technologyDataList, technologyData)
		}

		if err := c.ShouldBindBodyWith(&technologyData, binding.JSON); err != nil {
			response.Data = technologyDataList
			logger.Errorf("Failed to bind request: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}

		record := &dao.BTargetValidation{
			Content:     technologyData.Content,
			Image:       technologyData.Image,
			Align:       technologyData.Align,
			Lang:        lang,
		}
		_, err := t.MysqlClient.Omit("created_time", "updated_time").InsertOne(record)
		if err != nil {
			response.Data = technologyDataList
			logger.Errorf("Fail to add the targetValidation record: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError,
				fmt.Sprintf("Failed to add the targetValidation record!"))
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
				Image:       string(record["image"]),
				Align:       string(record["align"]),
			}
			technologyDataList = append(technologyDataList, technologyData)
		}

		if err := c.ShouldBindBodyWith(&technologyData, binding.JSON); err != nil {
			response.Data = technologyDataList
			logger.Errorf("Failed to bind request: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}

		record := &dao.BSBDD{
			Content:     technologyData.Content,
			Image:       technologyData.Image,
			Align:       technologyData.Align,
			Lang:        lang,
		}
		_, err := t.MysqlClient.Omit("created_time", "updated_time").InsertOne(record)
		if err != nil {
			response.Data = technologyDataList
			logger.Errorf("Failed to add the SBDD record: %s", err.Error())
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

func (t *technology) BiomarkerCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			technologyData  dao.TechnologyData
			response        dao.TechnologyResponse
		)

		logger := t.Logger.Named("CreateBiomarkerPlatformInfo")
		lang := c.Query("lang")

		technologyDataList := make([]dao.TechnologyData, 0)
		sql := fmt.Sprintf("select * from b_biomarker where lang='%s';", lang)
		records, _ := t.MysqlClient.Query(sql)
		for _, record := range records {
			technologyData = dao.TechnologyData{
				Id:          string(record["id"]),
				Content:     string(record["content"]),
				Image:       string(record["image"]),
				Align:       string(record["align"]),
			}
			technologyDataList = append(technologyDataList, technologyData)
		}

		if err := c.ShouldBindBodyWith(&technologyData, binding.JSON); err != nil {
			response.Data = technologyDataList
			logger.Errorf("Failed to bind request: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}

		record := &dao.BBiomarker{
			Content:     technologyData.Content,
			Image:       technologyData.Image,
			Align:       technologyData.Align,
			Lang:        lang,
		}
		_, err := t.MysqlClient.Omit("created_time", "updated_time").InsertOne(record)
		if err != nil {
			response.Data = technologyDataList
			logger.Errorf("Failed to add the biomarker record: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError,
				fmt.Sprintf("Failed to add the biomarker record!"))
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

func (t *technology) TargetValidationUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			technologyData  dao.TechnologyData
			response        dao.TechnologyResponse
		)

		logger := t.Logger.Named("UpdateTargetValidationPlatformInfo")
		lang := c.Query("lang")

		technologyDataList := make([]dao.TechnologyData, 0)
		sql := fmt.Sprintf("select * from b_target_validation where lang='%s';", lang)
		records, _ := t.MysqlClient.Query(sql)
		for _, record := range records {
			technologyData = dao.TechnologyData{
				Id:          string(record["id"]),
				Content:     string(record["content"]),
				Image:       string(record["image"]),
				Align:       string(record["align"]),
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

		record := &dao.BTargetValidation{
			Content:     technologyData.Content,
			Image:       technologyData.Image,
			Align:       technologyData.Align,
			Lang:        lang,
		}

		_, err := t.MysqlClient.Omit("created_time", "updated_time").Where(
			"id = ?", technologyId).Update(record)
		if err != nil {
			response.Data = technologyDataList
			logger.Errorf("Failed to update the targetValidation record: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}

		for index, data := range technologyDataList {
			if data.Id == c.Param("id") {
				technologyDataList[index] = dao.TechnologyData{
					Id:          data.Id,
					Content:     record.Content,
					Image:       record.Image,
					Align:       record.Align,
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
				Image:       string(record["image"]),
				Align:       string(record["align"]),
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
			Content:     technologyData.Content,
			Image:       technologyData.Image,
			Align:       technologyData.Align,
			Lang:        lang,
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
					Image:       record.Image,
					Align:       record.Align,
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

func (t *technology) BiomarkerUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			technologyData  dao.TechnologyData
			response        dao.TechnologyResponse
		)

		logger := t.Logger.Named("UpdateBiomarkerPlatformInfo")
		lang := c.Query("lang")

		technologyDataList := make([]dao.TechnologyData, 0)
		sql := fmt.Sprintf("select * from b_biomarker where lang='%s';", lang)
		records, _ := t.MysqlClient.Query(sql)
		for _, record := range records {
			technologyData = dao.TechnologyData{
				Id:          string(record["id"]),
				Content:     string(record["content"]),
				Image:       string(record["image"]),
				Align:       string(record["align"]),
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

		record := &dao.BBiomarker{
			Content:     technologyData.Content,
			Image:       technologyData.Image,
			Align:       technologyData.Align,
			Lang:        lang,
		}

		_, err := t.MysqlClient.Omit("created_time", "updated_time").Where(
			"id = ?", technologyId).Update(record)
		if err != nil {
			response.Data = technologyDataList
			logger.Errorf("Failed to update the biomarker record: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}

		for index, data := range technologyDataList {
			if data.Id == c.Param("id") {
				technologyDataList[index] = dao.TechnologyData{
					Id:          data.Id,
					Content:     record.Content,
					Image:       record.Image,
					Align:       record.Align,
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

func (t *technology) TargetValidationDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			indexDeleted       int
			technologyData     dao.TechnologyData
			response           dao.TechnologyResponse
		)
		logger := t.Logger.Named("DeleteTargetValidationPlatformInfo")
		lang := c.Query("lang")

		technologyDataList := make([]dao.TechnologyData, 0)
		sql := fmt.Sprintf("select * from b_target_validation where lang='%s';", lang)
		records, _ := t.MysqlClient.Query(sql)
		for _, record := range records {
			technologyData = dao.TechnologyData{
				Id:          string(record["id"]),
				Content:     string(record["content"]),
				Image:       string(record["image"]),
				Align:       string(record["align"]),
			}
			technologyDataList = append(technologyDataList, technologyData)
		}

		technologyId, _ := strconv.ParseInt(c.Param("id"), 10, 64)

		_, err := t.MysqlClient.Where("id = ?", technologyId).Delete(dao.BTargetValidation{})
		if err != nil {
			response.Data = technologyDataList
			logger.Errorf("Failed to delete the targetValidation record: %s", err.Error())
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
				Image:       string(record["image"]),
				Align:       string(record["align"]),
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

func (t *technology) BiomarkerDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			indexDeleted       int
			technologyData     dao.TechnologyData
			response           dao.TechnologyResponse
		)
		logger := t.Logger.Named("DeleteBiomarkerPlatformInfo")
		lang := c.Query("lang")

		technologyDataList := make([]dao.TechnologyData, 0)
		sql := fmt.Sprintf("select * from b_biomarker where lang='%s';", lang)
		records, _ := t.MysqlClient.Query(sql)
		for _, record := range records {
			technologyData = dao.TechnologyData{
				Id:          string(record["id"]),
				Content:     string(record["content"]),
				Image:       string(record["image"]),
				Align:       string(record["align"]),
			}
			technologyDataList = append(technologyDataList, technologyData)
		}

		technologyId, _ := strconv.ParseInt(c.Param("id"), 10, 64)

		_, err := t.MysqlClient.Where("id = ?", technologyId).Delete(dao.BBiomarker{})
		if err != nil {
			response.Data = technologyDataList
			logger.Errorf("Failed to delete the biomarker record: %s", err.Error())
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
