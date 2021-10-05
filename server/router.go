package server

import (
	"net/http"
	"os"
	"os/signal"
	"path"
	"syscall"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"iaso/config"
	"iaso/handler"
	log "iaso/logger"
)

var (
	onlyOneSignalHandler = make(chan struct{})
	shutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM}
)

// SetupSignalHandler registered for SIGTERM and SIGINT. A stop channel is returned
// which is closed on one of these signals. If a second signal is caught, the program
// is terminated with exit code 1.
func SetupSignalHandler() (stopCh <-chan struct{}) {
	close(onlyOneSignalHandler) // panics when called twice

	stop := make(chan struct{})
	c := make(chan os.Signal, 2)
	signal.Notify(c, shutdownSignals...)
	go func() {
		<-c
		close(stop)
		<-c
		os.Exit(1) // second signal. Exit directly.
	}()

	return stop
}

// fix cross domain problem
func AllowCors(crossConfig config.CrossConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Origin", crossConfig.AccessControlAllowOrigin)
			c.Header("Access-Control-Allow-Methods", crossConfig.AccessControlAllowMethods)
			c.Header("Access-Control-Allow-Headers", crossConfig.AccessControlAllowHeaders)
			c.Header("Access-Control-Expose-Headers", crossConfig.AccessControlExposeHeaders)
			c.Header("Access-Control-Allow-Credentials", crossConfig.AccessControlAllowCredentials)
		}

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func Init(logger *zap.SugaredLogger, verbose bool, crossConfig config.CrossConfig, distFilePath string) *gin.Engine {
	routerLogger := logger.Named("Router")
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(log.AccessLoggerForGin())
	//router.Use(AllowCors(crossConfig))

	if !verbose {
		gin.SetMode(gin.ReleaseMode)
	}

	staticFilePath := path.Join(distFilePath)
	faviconPath := path.Join(distFilePath, "/favicon.png")
	entryHtmlPath := path.Join(distFilePath, "/index.html")
	svcWorkerPath := path.Join(distFilePath, "/service-worker.js")
	routerLogger.Debugf("The UI static file path: %s", staticFilePath)
	routerLogger.Debugf("The UI favicon.png path: %s", faviconPath)
	routerLogger.Debugf("The UI html path: %s", entryHtmlPath)

	router.Static("/dist", staticFilePath)
	router.StaticFile("/favicon.ico", faviconPath)
	router.StaticFile("/service-worker.js", svcWorkerPath)
	router.LoadHTMLGlob(entryHtmlPath)

	//router.Use(static.Serve("/dist", static.LocalFile(distFilePath, false)))
	//router.StaticFS("/", http.Dir("/Users/jenson/go/src/iaso/ui/dist"))
	// Service API: v1
	routerLogger.Info("Init service API group")
	apiGroup := router.Group("/api")
	if verbose {
		requestLogger := log.NewHttpLogger(logger)
		apiGroup.Use(requestLogger.HttpLogger()) // Print the request body
	}

	v1Group := apiGroup.Group("/v1")

	// admin info handler
	//TODO: all users info handler
	userV1Group := v1Group.Group("/users")
	users := handler.NewUsers(logger)
	{
		userV1Group.GET("/:username", users.Get())
		userV1Group.POST("", users.Login())
	}

	//about us handler
	aboutUsV1Group := v1Group.Group("/about-us")
	aboutUs := handler.NewAboutUs(logger)
	{
		aboutUsV1Group.GET("", aboutUs.List())
		aboutUsV1Group.POST("", aboutUs.Create())
		aboutUsV1Group.PUT("/:id", aboutUs.Update())
		aboutUsV1Group.DELETE("/:id", aboutUs.Delete())
	}

	//technology handler
	technologyV1Group := v1Group.Group("/technology")
	technology := handler.NewTechnology(logger)
	{
		technologyV1Group.GET("/target-validation", technology.TargetValidationList())
		technologyV1Group.GET("/sbdd", technology.SBDDList())
		technologyV1Group.GET("/biomarker-development", technology.BiomarkerList())
		technologyV1Group.POST("/target-validation", technology.TargetValidationCreate())
		technologyV1Group.POST("/sbdd", technology.SBDDCreate())
		technologyV1Group.POST("/biomarker-development", technology.BiomarkerCreate())
		technologyV1Group.PUT("/target-validation/:id", technology.TargetValidationUpdate())
		technologyV1Group.PUT("/sbdd/:id", technology.SBDDUpdate())
		technologyV1Group.PUT("/biomarker-development/:id", technology.BiomarkerUpdate())
		technologyV1Group.DELETE("/target-validation/:id", technology.TargetValidationDelete())
		technologyV1Group.DELETE("/sbdd/:id", technology.SBDDDelete())
		technologyV1Group.DELETE("/biomarker-development/:id", technology.BiomarkerDelete())
	}

	////products handler
	//productsV1Group := v1Group.Group("/products")
	//products := handler.NewProducts(logger)
	//{
	//	productsV1Group.GET("", products.List())
	//	productsV1Group.POST("", products.Create())
	//	productsV1Group.PUT("/:id", products.Update())
	//	productsV1Group.DELETE("/:id", products.Delete())
	//
	//}

	//partnering handler
	partneringV1Group := v1Group.Group("/partnering")
	partnering := handler.NewPartnering(logger)
	{
		partneringV1Group.GET("/biotech-company", partnering.BiotechCompanyList())
		partneringV1Group.GET("/academic-institution", partnering.AcademicInstitutionList())
		partneringV1Group.PUT("/biotech-company/:id", partnering.BiotechCompanyUpdate())
		partneringV1Group.PUT("/academic-institution/:id", partnering.AcademicInstitutionUpdate())
		partneringV1Group.DELETE("/biotech-company/:id", partnering.BiotechCompanyDelete())
		partneringV1Group.DELETE("/academic-institution/:id", partnering.AcademicInstitutionDelete())
		partneringV1Group.POST("/biotech-company", partnering.BiotechCompanyCreate())
		partneringV1Group.POST("/academic-institution", partnering.AcademicInstitutionCreate())
	}

	//pipeline handler
	pipelineV1Group := v1Group.Group("/pipeline")
	pipeline := handler.NewPipeline(logger)
	{
		pipelineV1Group.GET("", pipeline.List())
		pipelineV1Group.POST("", pipeline.Create())
		pipelineV1Group.PUT("/:id", pipeline.Update())
		pipelineV1Group.DELETE("/:id", pipeline.Delete())
	}

	//news handler
	newsV1Group := v1Group.Group("/news")
	news := handler.NewNews(logger)
	{
		newsV1Group.GET("", news.List())
		newsV1Group.PUT("/:id",news.Update())
		newsV1Group.DELETE("/:id", news.Delete())
		newsV1Group.POST("", news.Create())
	}

	//careers handler
	careersV1Group := v1Group.Group("/careers")
	careers := handler.NewCareers(logger)
	{
		careersV1Group.GET("", careers.List())
		careersV1Group.PUT("/:id",careers.Update())
		careersV1Group.DELETE("/:id", careers.Delete())
		careersV1Group.POST("", careers.Create())
	}

	//contact us handler
	contactUsV1Group := v1Group.Group("/contact-us")
	contactUs := handler.NewContactUs(logger)
	{
		contactUsV1Group.POST("", contactUs.Create())
	}

	////load html
	router.Use().StaticFile("/", entryHtmlPath)
	return router
}

