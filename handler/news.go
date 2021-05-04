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

type News interface {
	List() gin.HandlerFunc
	Create() gin.HandlerFunc
	Update() gin.HandlerFunc
	Delete() gin.HandlerFunc
}

type news struct {
	Logger            *zap.SugaredLogger
	MysqlClient       xorm.Interface
}

func NewNews(logger *zap.SugaredLogger) News {
	return &news{
		Logger:        logger.Named("news"),
		MysqlClient:   mysql.Client,
	}
}

func (n *news) List() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			newsData    dao.NewsData
			response    dao.NewsResponse
		)

		logger := n.Logger.Named("GetNewsInfo")
		lang := c.Query("lang")
		title := c.Query("title")

		newsDataList := make([]dao.NewsData, 0)
		sql := fmt.Sprintf("select * from b_news where lang='%s';", lang)
		records, _ := n.MysqlClient.Query(sql)
		for _, record := range records {
			newsData = dao.NewsData{
				Id:          string(record["id"]),
				Title:       string(record["title"]),
				Description: string(record["description"]),
				Content:     string(record["content"]),
				CreatedAt:   string(record["created_time"]),
				UpdatedAt:   string(record["updated_time"]),
			}
			if title != "" && newsData.Title != title {
				continue
			}
			newsDataList = append(newsDataList, newsData)
		}

		newsDataList = reverseNews(newsDataList)
		logger.Infof("Succeeded to get the record from database.")
		response.Data = newsDataList
		dao.Success(c, &response, http.StatusOK)
		return
	}
}

func (n *news) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			newsData     dao.NewsData
			response     dao.NewsResponse
		)

		logger := n.Logger.Named("CreateNewsInfo")
		lang := c.Query("lang")

		newsDataList := make([]dao.NewsData, 0)
		sql := fmt.Sprintf("select * from b_news where lang='%s';", lang)
		records, _ := n.MysqlClient.Query(sql)
		for _, record := range records {
			newsData = dao.NewsData{
				Id:          string(record["id"]),
				Title:       string(record["title"]),
				Description: string(record["description"]),
				Content:     string(record["content"]),
				CreatedAt:   string(record["created_time"]),
				UpdatedAt:   string(record["updated_time"]),
			}
			newsDataList = append(newsDataList, newsData)
		}

		if err := c.ShouldBindBodyWith(&newsData, binding.JSON); err != nil {
			newsDataList = reverseNews(newsDataList)
			response.Data = newsDataList
			logger.Errorf("Failed to bind request: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}

		record := &dao.BNews{}
		has, err := n.MysqlClient.Where("title = ? and lang = ?", newsData.Title, lang).Get(record)
		if err != nil {
			newsDataList = reverseNews(newsDataList)
			response.Data = newsDataList
			logger.Errorf("Failed to add the news record: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}
		if has {
			newsDataList = reverseNews(newsDataList)
			response.Data = newsDataList
			logger.Errorf("The record existed, submit duplicated!")
			dao.FailWithMessage(c, &response, http.StatusConflict, fmt.Sprintf("The record existed, submit duplicated!"))
			return
		}

		record = &dao.BNews{
			Title:       newsData.Title,
			Description: newsData.Description,
			Content:     newsData.Content,
			Lang:        lang,
		}
		_, err = n.MysqlClient.Omit("created_time", "updated_time").InsertOne(record)
		if err != nil {
			response.Data = newsDataList
			logger.Errorf("Failed to add the news record: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, fmt.Sprintf("Failed to add the news record!"))
			return
		}
		insertedData := &dao.BNews{}
		newsData.Id = strconv.FormatInt(record.Id, 10)
		has, err = n.MysqlClient.ID(newsData.Id).Get(insertedData)
		if err != nil {
			response.Data = newsDataList
			logger.Errorf("Failed to add the news record: %s", newsData.Title)
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, fmt.Sprintf("Failed to add the news record!"))
			return
		}
		if has {
			newsData.CreatedAt = insertedData.CreatedAt.Format("2006-01-02 15:04:05")
			newsData.UpdatedAt = insertedData.UpdatedAt.Format("2006-01-02 15:04:05")
		}
		newsDataList = append(newsDataList, newsData)
		newsDataList = reverseNews(newsDataList)
		response.Data = newsDataList
		logger.Debugf("Add a record named %s into database.", record.Title)
		dao.Success(c, &response, http.StatusCreated)
		return
	}
}

func (n *news) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			newsData     dao.NewsData
			response     dao.NewsResponse
		)

		logger := n.Logger.Named("UpdateNewsInfo")
		lang := c.Query("lang")

		newsDataList := make([]dao.NewsData, 0)
		sql := fmt.Sprintf("select * from b_news where lang='%s';", lang)
		records, _ := n.MysqlClient.Query(sql)
		for _, record := range records {
			newsData = dao.NewsData{
				Id:          string(record["id"]),
				Title:       string(record["title"]),
				Description: string(record["description"]),
				Content:     string(record["content"]),
				CreatedAt:   string(record["created_time"]),
				UpdatedAt:   string(record["updated_time"]),
			}
			newsDataList = append(newsDataList, newsData)
		}
		newsDataList = reverseNews(newsDataList)

		if err := c.ShouldBindBodyWith(&newsData, binding.JSON); err != nil {
			response.Data = newsDataList
			logger.Errorf("Failed to bind request: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}

		newsId, _ := strconv.ParseInt(c.Param("id"), 10, 64)

		record := &dao.BNews{
			Id:           newsId,
			Title:        newsData.Title,
			Description:  newsData.Description,
			Content:      newsData.Content,
			Lang:         lang,
		}

		_, err := n.MysqlClient.Omit("created_time", "updated_time").Where("id = ?", newsId).Update(record)
		if err != nil {
			response.Data = newsDataList
			logger.Errorf("Failed to update the news record: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}

		for index, data := range newsDataList {
			if data.Id == c.Param("id") {
				newsDataList[index] = dao.NewsData{
					Id:          data.Id,
					Title:       record.Title,
					Description: record.Description,
					Content:     record.Content,
					CreatedAt:   record.CreatedAt.Format("2006-01-02 15:04:05"),
					UpdatedAt:   record.UpdatedAt.Format("2006-01-02 15:04:05"),
				}
				break
			}
		}
		response.Data = newsDataList
		logger.Infof("Updated a record named %s into database.", record.Title)
		dao.Success(c, &response, http.StatusOK)
		return
	}
}

func (n *news) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			indexDeleted   int
			newsData       dao.NewsData
			response       dao.NewsResponse
		)
		logger := n.Logger.Named("DeleteNewsInfo")
		lang := c.Query("lang")

		newsDataList := make([]dao.NewsData, 0)
		sql := fmt.Sprintf("select * from b_news where lang='%s';", lang)
		records, _ := n.MysqlClient.Query(sql)
		for _, record := range records {
			newsData = dao.NewsData{
				Id:          string(record["id"]),
				Title:       string(record["title"]),
				Description: string(record["description"]),
				Content:     string(record["content"]),
				CreatedAt:   string(record["created_time"]),
				UpdatedAt:   string(record["updated_time"]),
			}
			newsDataList = append(newsDataList, newsData)
		}
		newsDataList = reverseNews(newsDataList)

		newsId, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		_, err := n.MysqlClient.Where("id = ?", newsId).Delete(dao.BNews{})
		if err != nil {
			response.Data = newsDataList
			logger.Errorf("Failed to delete the news record: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}


		for index, data := range newsDataList {
			if data.Id == c.Param("id") {
				indexDeleted = index
				break
			}
		}

		newsDataList = append(newsDataList[:indexDeleted], newsDataList[indexDeleted + 1:]...)
		response.Data = newsDataList
		logger.Debugf("Succeeded to delete a record.")
		dao.Success(c, &response, http.StatusOK)
	}
}

func reverseNews(arr []dao.NewsData) []dao.NewsData{
	length := len(arr)
	for index := range arr[:length / 2] {
		temp := arr[length - 1 - index]
		arr[length - 1 - index] = arr[index]
		arr[index] = temp
	}
	return arr
}

