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

type Partnering interface {
	AcademicInstitutionList() gin.HandlerFunc
	AcademicInstitutionCreate() gin.HandlerFunc
	AcademicInstitutionUpdate() gin.HandlerFunc
	AcademicInstitutionDelete() gin.HandlerFunc
	IndustrialInstitutionList()   gin.HandlerFunc
	IndustrialInstitutionCreate() gin.HandlerFunc
	IndustrialInstitutionUpdate() gin.HandlerFunc
	IndustrialInstitutionDelete() gin.HandlerFunc
}

type partnering struct {
	Logger            *zap.SugaredLogger
	MysqlClient       xorm.Interface
}

func NewPartnering(logger *zap.SugaredLogger) Partnering {
	return &partnering{
		Logger:        logger.Named("partnering"),
		MysqlClient:   mysql.Client,
	}
}

func (pt *partnering) IndustrialInstitutionList() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			partneringData    dao.PartneringData
			response          dao.PartneringResponse
		)

		logger := pt.Logger.Named("GetIndustrialInstitutionInfo")
		lang := c.Query("lang")

		partneringDataList := make([]dao.PartneringData, 0)
		sql := fmt.Sprintf("select * from b_industrial_institution where lang='%s';", lang)
		records, _ := pt.MysqlClient.Query(sql)
		for _, record := range records {
			partneringData = dao.PartneringData{
				Id:          string(record["id"]),
				Content:     string(record["content"]),
				Image:       string(record["image"]),
				Align:       string(record["align"]),
			}
			partneringDataList = append(partneringDataList, partneringData)
		}

		logger.Infof("Succeeded to get the record from database.")
		response.Data = partneringDataList
		dao.Success(c, &response, http.StatusOK)
		return
	}
}

func (pt *partnering) AcademicInstitutionList() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			partneringData    dao.PartneringData
			response          dao.PartneringResponse
		)

		logger := pt.Logger.Named("GetAcademicInstitutionInfo")
		lang := c.Query("lang")

		partneringDataList := make([]dao.PartneringData, 0)
		sql := fmt.Sprintf("select * from b_academic_institution where lang='%s';", lang)
		records, _ := pt.MysqlClient.Query(sql)
		for _, record := range records {
			partneringData = dao.PartneringData{
				Id:          string(record["id"]),
				Content:     string(record["content"]),
				Image:       string(record["image"]),
				Align:       string(record["align"]),
			}
			partneringDataList = append(partneringDataList, partneringData)
		}

		logger.Infof("Succeeded to get the record from database.")
		response.Data = partneringDataList
		dao.Success(c, &response, http.StatusOK)
		return
	}
}

func (pt *partnering) IndustrialInstitutionCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			partneringData     dao.PartneringData
			response           dao.PartneringResponse
		)

		logger := pt.Logger.Named("CreateIndustrialInstitutionInfo")
		lang := c.Query("lang")

		partneringDataList := make([]dao.PartneringData, 0)
		sql := fmt.Sprintf("select * from b_industrial_institution where lang='%s';", lang)
		records, _ := pt.MysqlClient.Query(sql)
		for _, record := range records {
			partneringData = dao.PartneringData{
				Id:          string(record["id"]),
				Content:     string(record["content"]),
				Image:       string(record["image"]),
				Align:       string(record["align"]),
			}
			partneringDataList = append(partneringDataList, partneringData)
		}

		if err := c.ShouldBindBodyWith(&partneringData, binding.JSON); err != nil {
			response.Data = partneringDataList
			logger.Errorf("Failed to bind request: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}

		record := &dao.BIndustrialInstitution{
			Content:     partneringData.Content,
			Image:       partneringData.Image,
			Align:       partneringData.Align,
			Lang:        lang,
		}
		_, err := pt.MysqlClient.Omit("created_time", "updated_time").InsertOne(record)
		if err != nil {
			response.Data = partneringDataList
			logger.Errorf("Failed to add the industrialInstitution record: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, fmt.Sprintf("Failed to add the industrialInstitution record!"))
			return
		}
		partneringData.Id = strconv.FormatInt(record.Id, 10)
		partneringDataList = append(partneringDataList, partneringData)
		response.Data = partneringDataList
		logger.Debugf("Add a record id: %s into database.", record.Id)
		dao.Success(c, &response, http.StatusCreated)
		return
	}
}

func (pt *partnering) AcademicInstitutionCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			partneringData     dao.PartneringData
			response           dao.PartneringResponse
		)

		logger := pt.Logger.Named("CreateAcademicInstitutionInfo")
		lang := c.Query("lang")

		partneringDataList := make([]dao.PartneringData, 0)
		sql := fmt.Sprintf("select * from b_academic_institution where lang='%s';", lang)
		records, _ := pt.MysqlClient.Query(sql)
		for _, record := range records {
			partneringData = dao.PartneringData{
				Id:          string(record["id"]),
				Content:     string(record["content"]),
				Image:       string(record["image"]),
				Align:       string(record["align"]),
			}
			partneringDataList = append(partneringDataList, partneringData)
		}

		if err := c.ShouldBindBodyWith(&partneringData, binding.JSON); err != nil {
			response.Data = partneringDataList
			logger.Errorf("Failed to bind request: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}

		record := &dao.BAcademicInstitution{
			Content:     partneringData.Content,
			Image:       partneringData.Image,
			Align:       partneringData.Align,
			Lang:        lang,
		}
		_, err := pt.MysqlClient.Omit("created_time", "updated_time").InsertOne(record)
		if err != nil {
			response.Data = partneringDataList
			logger.Errorf("Failed to add the academicInstitution record: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, fmt.Sprintf("Failed to add the academicInstitution record!"))
			return
		}
		partneringData.Id = strconv.FormatInt(record.Id, 10)
		partneringDataList = append(partneringDataList, partneringData)
		response.Data = partneringDataList
		logger.Debugf("Add a record id: %s into database.", record.Id)
		dao.Success(c, &response, http.StatusCreated)
		return
	}
}

func (pt *partnering) IndustrialInstitutionUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			partneringData   dao.PartneringData
			response         dao.PartneringResponse
		)

		logger := pt.Logger.Named("UpdateIndustrialInstitutionInfo")
		lang := c.Query("lang")

		partneringDataList := make([]dao.PartneringData, 0)
		sql := fmt.Sprintf("select * from b_industrial_institution where lang='%s';", lang)
		records, _ := pt.MysqlClient.Query(sql)
		for _, record := range records {
			partneringData = dao.PartneringData{
				Id:          string(record["id"]),
				Content:     string(record["content"]),
				Image:       string(record["image"]),
				Align:       string(record["align"]),
			}
			partneringDataList = append(partneringDataList, partneringData)
		}

		if err := c.ShouldBindBodyWith(&partneringData, binding.JSON); err != nil {
			response.Data = partneringDataList
			logger.Errorf("Failed to bind request: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}

		partneringId, _ := strconv.ParseInt(c.Param("id"), 10, 64)

		record := &dao.BIndustrialInstitution{
			Content:      partneringData.Content,
			Image:        partneringData.Image,
			Align:        partneringData.Align,
			Lang:         lang,
		}

		_, err := pt.MysqlClient.Omit("created_time", "updated_time").Where(
			"id = ?", partneringId).Update(record)
		if err != nil {
			response.Data = partneringDataList
			logger.Errorf("Failed to update the industrialInstitution record: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}

		for index, data := range partneringDataList {
			if data.Id == c.Param("id") {
				partneringDataList[index] = dao.PartneringData{
					Id:          data.Id,
					Content:     record.Content,
					Image:       record.Image,
					Align:       record.Align,
				}
				break
			}
		}
		response.Data = partneringDataList
		logger.Debugf("Updated a record named %s into database.", record.Id)
		dao.Success(c, &response, http.StatusOK)
		return
	}
}

func (pt *partnering) AcademicInstitutionUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			partneringData   dao.PartneringData
			response         dao.PartneringResponse
		)

		logger := pt.Logger.Named("UpdateAcademicInstitutionInfo")
		lang := c.Query("lang")

		partneringDataList := make([]dao.PartneringData, 0)
		sql := fmt.Sprintf("select * from b_academic_institution where lang='%s';", lang)
		records, _ := pt.MysqlClient.Query(sql)
		for _, record := range records {
			partneringData = dao.PartneringData{
				Id:          string(record["id"]),
				Content:     string(record["content"]),
				Image:       string(record["image"]),
				Align:       string(record["align"]),
			}
			partneringDataList = append(partneringDataList, partneringData)
		}

		if err := c.ShouldBindBodyWith(&partneringData, binding.JSON); err != nil {
			response.Data = partneringDataList
			logger.Errorf("Failed to bind request: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}

		partneringId, _ := strconv.ParseInt(c.Param("id"), 10, 64)

		record := &dao.BAcademicInstitution{
			Content:      partneringData.Content,
			Image:        partneringData.Image,
			Align:        partneringData.Align,
			Lang:         lang,
		}

		_, err := pt.MysqlClient.Omit("created_time", "updated_time").Where(
			"id = ?", partneringId).Update(record)
		if err != nil {
			response.Data = partneringDataList
			logger.Errorf("Failed to update the academicInstitution record: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}

		for index, data := range partneringDataList {
			if data.Id == c.Param("id") {
				partneringDataList[index] = dao.PartneringData{
					Id:          data.Id,
					Content:     record.Content,
					Image:       record.Image,
					Align:       record.Align,
				}
				break
			}
		}
		response.Data = partneringDataList
		logger.Debugf("Updated a record named %s into database.", record.Id)
		dao.Success(c, &response, http.StatusOK)
		return
	}
}

func (pt *partnering) IndustrialInstitutionDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			indexDeleted       int
			partneringData     dao.PartneringData
			response           dao.PartneringResponse
		)
		logger := pt.Logger.Named("DeleteIndustrialInstitutionInfo")
		lang := c.Query("lang")

		partneringDataList := make([]dao.PartneringData, 0)
		sql := fmt.Sprintf("select * from b_industrial_institution where lang='%s';", lang)
		records, _ := pt.MysqlClient.Query(sql)
		for _, record := range records {
			partneringData = dao.PartneringData{
				Id:          string(record["id"]),
				Content:     string(record["content"]),
				Image:       string(record["image"]),
				Align:       string(record["align"]),
			}
			partneringDataList = append(partneringDataList, partneringData)
		}

		partneringId, _ := strconv.ParseInt(c.Param("id"), 10, 64)

		_, err := pt.MysqlClient.Where("id = ?", partneringId).Delete(dao.BIndustrialInstitution{})
		if err != nil {
			response.Data = partneringDataList
			logger.Errorf("Failed to delete the industrialInstitution record: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}


		for index, data := range partneringDataList {
			if data.Id == c.Param("id") {
				indexDeleted = index
				break
			}
		}

		partneringDataList = append(partneringDataList[:indexDeleted], partneringDataList[indexDeleted + 1:]...)
		response.Data = partneringDataList
		logger.Debugf("Succeeded to delete a record.")
		dao.Success(c, &response, http.StatusOK)
	}
}

func (pt *partnering) AcademicInstitutionDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			indexDeleted       int
			partneringData     dao.PartneringData
			response           dao.PartneringResponse
		)
		logger := pt.Logger.Named("DeleteAcademicInstitutionInfo")
		lang := c.Query("lang")

		partneringDataList := make([]dao.PartneringData, 0)
		sql := fmt.Sprintf("select * from b_academic_institution where lang='%s';", lang)
		records, _ := pt.MysqlClient.Query(sql)
		for _, record := range records {
			partneringData = dao.PartneringData{
				Id:          string(record["id"]),
				Content:     string(record["content"]),
				Image:       string(record["image"]),
				Align:       string(record["align"]),
			}
			partneringDataList = append(partneringDataList, partneringData)
		}

		partneringId, _ := strconv.ParseInt(c.Param("id"), 10, 64)

		_, err := pt.MysqlClient.Where("id = ?", partneringId).Delete(dao.BAcademicInstitution{})
		if err != nil {
			response.Data = partneringDataList
			logger.Errorf("Failed to delete the academicInstitution record: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}


		for index, data := range partneringDataList {
			if data.Id == c.Param("id") {
				indexDeleted = index
				break
			}
		}

		partneringDataList = append(partneringDataList[:indexDeleted], partneringDataList[indexDeleted + 1:]...)
		response.Data = partneringDataList
		logger.Debugf("Succeeded to delete a record.")
		dao.Success(c, &response, http.StatusOK)
	}
}
