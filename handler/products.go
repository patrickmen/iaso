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

type Products interface {
	List() gin.HandlerFunc
	Create() gin.HandlerFunc
	Update() gin.HandlerFunc
	Delete() gin.HandlerFunc
}

type products struct {
	Logger            *zap.SugaredLogger
	MysqlClient       xorm.Interface
}

func NewProducts(logger *zap.SugaredLogger) Products {
	return &products{
		Logger:        logger.Named("products"),
		MysqlClient:   mysql.Client,
	}
}

func (pd *products) List() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			productData    dao.ProductData
			response       dao.ProductsResponse
		)

		logger := pd.Logger.Named("GetProductInfo")
		lang := c.Query("lang")
		title := c.Param("title")

		productDataList := make([]dao.ProductData, 0)
		sql := fmt.Sprintf("select * from b_products where lang='%s';", lang)
		records, _ := pd.MysqlClient.Query(sql)
		for _, record := range records {
			productData = dao.ProductData{
				Id:          string(record["id"]),
				Title:       string(record["title"]),
				Cover:       string(record["cover"]),
				Description: string(record["description"]),
				Content:     string(record["content"]),
			}
			if title != "" && productData.Title != title {
				continue
			}
			productDataList = append(productDataList, productData)
		}

		productDataList = reverseProduct(productDataList)
		logger.Infof("Succeeded to get the record from database.")
		response.Data = productDataList
		dao.Success(c, &response, http.StatusOK)
		return
	}
}

func (pd *products) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			productData     dao.ProductData
			response         dao.ProductsResponse
		)

		logger := pd.Logger.Named("CreateProductInfo")
		lang := c.Query("lang")

		productDataList := make([]dao.ProductData, 0)
		sql := fmt.Sprintf("select * from b_products where lang='%s';", lang)
		records, _ := pd.MysqlClient.Query(sql)
		for _, record := range records {
			productData = dao.ProductData{
				Id:          string(record["id"]),
				Title:       string(record["title"]),
				Cover:       string(record["cover"]),
				Description: string(record["description"]),
				Content:     string(record["content"]),
			}
			productDataList = append(productDataList, productData)
		}

		if err := c.ShouldBindBodyWith(&productData, binding.JSON); err != nil {
			productDataList = reverseProduct(productDataList)
			response.Data = productDataList
			logger.Errorf("Failed to bind request: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}

		record := &dao.BProducts{}
		has, err := pd.MysqlClient.Where("title = ? and lang = ?", productData.Title, lang).Get(record)
		if err != nil {
			productDataList = reverseProduct(productDataList)
			response.Data = productDataList
			logger.Errorf("Failed to add the product record: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}
		if has {
			productDataList = reverseProduct(productDataList)
			response.Data = productDataList
			logger.Errorf("The record existed, submit duplicated!")
			dao.FailWithMessage(c, &response, http.StatusConflict, fmt.Sprintf("The record existed, submit duplicated!"))
			return
		}

		record = &dao.BProducts{
			Title:       productData.Title,
			Cover:       productData.Cover,
			Description: productData.Description,
			Content:     productData.Content,
			Lang:        lang,
		}
		_, err = pd.MysqlClient.InsertOne(record)
		if err != nil {
			response.Data = productDataList
			logger.Errorf("Fail to add the product record: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, fmt.Sprintf("Failed to add the product record!"))
			return
		}
		productData.Id = strconv.FormatInt(record.Id, 10)
		productDataList = append(productDataList, productData)
		productDataList = reverseProduct(productDataList)
		response.Data = productDataList
		logger.Debugf("Add a record named %s into database.", record.Title)
		dao.Success(c, &response, http.StatusCreated)
		return
	}
}

func (pd *products) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			productData     dao.ProductData
			response        dao.ProductsResponse
		)

		logger := pd.Logger.Named("UpdateProductInfo")
		lang := c.Query("lang")

		productDataList := make([]dao.ProductData, 0)
		sql := fmt.Sprintf("select * from b_products where lang='%s';", lang)
		records, _ := pd.MysqlClient.Query(sql)
		for _, record := range records {
			productData = dao.ProductData{
				Id:          string(record["id"]),
				Title:       string(record["title"]),
				Cover:       string(record["cover"]),
				Description: string(record["description"]),
				Content:     string(record["content"]),
			}
			productDataList = append(productDataList, productData)
		}
		productDataList = reverseProduct(productDataList)

		if err := c.ShouldBindBodyWith(&productData, binding.JSON); err != nil {
			response.Data = productDataList
			logger.Errorf("Failed to bind request: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}

		productId, _ := strconv.ParseInt(c.Param("id"), 10, 64)

		record := &dao.BProducts{
			Title:        productData.Title,
			Cover:        productData.Cover,
			Description:  productData.Description,
			Content:      productData.Content,
			Lang:         lang,
		}

		_, err := pd.MysqlClient.Omit("created_time", "updated_time").Where(
			"id = ?", productId).Update(record)
		if err != nil {
			response.Data = productDataList
			logger.Errorf("Failed to update the product record: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}

		for index, data := range productDataList {
			if data.Id == c.Param("id") {
				productDataList[index] = dao.ProductData{
					Id:          data.Id,
					Title:       record.Title,
					Cover:       record.Cover,
					Description: record.Description,
					Content:     record.Content,
				}
				break
			}
		}
		response.Data = productDataList
		logger.Debugf("Updated a record named %s into database.", record.Title)
		dao.Success(c, &response, http.StatusOK)
		return
	}
}


func (pd *products) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			indexDeleted       int
			productData        dao.ProductData
			response           dao.ProductsResponse
		)
		logger := pd.Logger.Named("DeleteProductInfo")
		lang := c.Query("lang")

		productDataList := make([]dao.ProductData, 0)
		sql := fmt.Sprintf("select * from b_products where lang='%s';", lang)
		records, _ := pd.MysqlClient.Query(sql)
		for _, record := range records {
			productData = dao.ProductData{
				Id:          string(record["id"]),
				Title:       string(record["title"]),
				Cover:       string(record["cover"]),
				Description: string(record["description"]),
				Content:     string(record["content"]),
			}
			productDataList = append(productDataList, productData)
		}
		productDataList = reverseProduct(productDataList)


		productId, _ := strconv.ParseInt(c.Param("id"), 10, 64)

		_, err := pd.MysqlClient.Where("id = ?", productId).Delete(dao.BProducts{})
		if err != nil {
			response.Data = productDataList
			logger.Errorf("Failed to delete the product record: %s", err.Error())
			dao.FailWithMessage(c, &response, http.StatusInternalServerError, err.Error())
			return
		}


		for index, data := range productDataList {
			if data.Id == c.Param("id") {
				indexDeleted = index
				break
			}
		}

		productDataList = append(productDataList[:indexDeleted], productDataList[indexDeleted + 1:]...)
		response.Data = productDataList
		logger.Debugf("Succeeded to delete a record.")
		dao.Success(c, &response, http.StatusOK)
	}
}

func reverseProduct(arr []dao.ProductData) []dao.ProductData{
	length := len(arr)
	for index := range arr[:length / 2] {
		temp := arr[length - 1 - index]
		arr[length - 1 - index] = arr[index]
		arr[index] = temp
	}
	return arr
}
