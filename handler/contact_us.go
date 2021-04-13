package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"
	"xorm.io/xorm"

	"iaso/dao"
	"iaso/mysql"
)

type ContactUs interface {
	Create() gin.HandlerFunc
}

type contactUs struct {
	Logger            *zap.SugaredLogger
	MysqlClient       xorm.Interface
}

func NewContactUs(logger *zap.SugaredLogger) ContactUs {
	return &contactUs{
		Logger:        logger.Named("contact-us"),
		MysqlClient:   mysql.Client,
	}
}

func (ct *contactUs) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			contactUsData   dao.BContactUs
			response        dao.ContactUsResponse
		)

		logger := ct.Logger.Named("CreateContactUsInfo")
		if err := c.ShouldBindBodyWith(&contactUsData, binding.JSON); err != nil {
			logger.Errorf("Failed to bind request: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}


		result := &dao.BContactUs{}
		has, err := ct.MysqlClient.Where("name = ? AND phone = ? AND email = ?", contactUsData.Name, contactUsData.Phone, contactUsData.Email).Get(result)
		if err != nil {
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}
		if has {
			dao.FailWithMessage(c, &response, http.StatusConflict, fmt.Sprintf("The record existed, submit duplicated!"))
			return
		}

		_, err = ct.MysqlClient.Omit("created_time", "updated_time").InsertOne(contactUsData)
		if err != nil {
			logger.Errorf("Failed to insert record into contact-us table: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, fmt.Sprintf("Failed to insert the record!"))
			return
		}
		logger.Infof("Insert a record into database: %+v", contactUsData)
		dao.Success(c, &response, http.StatusCreated)
		return
	}
}