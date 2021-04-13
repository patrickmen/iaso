package handler

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"
	"xorm.io/xorm"

	"iaso/dao"
	"iaso/mysql"
)

type Users interface {
	Get() gin.HandlerFunc
	Login() gin.HandlerFunc
}

type users struct {
	Logger            *zap.SugaredLogger
	MysqlClient       xorm.Interface
}

func NewUsers(logger *zap.SugaredLogger) Users {
	return &users{
		Logger:        logger.Named("users"),
		MysqlClient:   mysql.Client,
	}
}

func (u *users) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			username      string
			user          dao.BAuthUsers
			response      dao.UserResponse
		)

		logger := u.Logger.Named("GetUserInfo")
		username = c.Param("username")
		_, err := u.MysqlClient.Where("username=?", username).Get(&user)
		if err != nil {
			logger.Errorf("Failed to get the user %s: %s", username, err.Error())
			dao.FailWithMessage(c, &response, http.StatusConflict, fmt.Sprintf("Encounter a conflict when query user!"))
			return
		}
		response = dao.UserResponse{
			Id:   user.Id,
		}
		dao.Success(c, &response, http.StatusOK)
		return
	}
}

func (u *users) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			user          dao.BAuthUsers
			response      dao.UserResponse
		)

		logger := u.Logger.Named("CreateUserInfo")
		if err := c.ShouldBindBodyWith(&user, binding.JSON); err != nil {
			logger.Errorf("Failed to bind request: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}

		result := &dao.BAuthUsers{}
		has, err := u.MysqlClient.Where("username = ? AND password = ?", user.Username, user.Password).Get(result)
		response.Id = result.Id
		if !has {
			dao.FailWithMessage(c, &response, http.StatusNotFound, fmt.Sprintf("The user does not existed!"))
			return
		}
		if err != nil {
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}

		rand.Seed(time.Now().Unix())
		response.Data = dao.UserData{
			User:     user.Username,
			Token:    "Authorization:" + strconv.FormatFloat(rand.Float64(), 'E', -1, 64),
			ExpireAt: time.Now().Add(time.Hour),
		}
		logger.Infof("Get the record from database: %+v", response.Id)
		dao.Success(c, &response, http.StatusOK)
		return
	}
}
