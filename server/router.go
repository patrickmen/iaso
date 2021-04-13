package server

import (
	"net/http"
	"os"
	"os/signal"
	"path"
	"syscall"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

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
func AllowCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Origin", "http://localhost:8080")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, Content-Type, X-CSRF-Token, Token, session")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func Init(logger *zap.SugaredLogger, verbose bool, distFilePath string) *gin.Engine {
	routerLogger := logger.Named("Router")
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(log.AccessLoggerForGin())
	//router.Use(AllowCors())

	if !verbose {
		gin.SetMode(gin.ReleaseMode)
	}

	staticFilePath := path.Join(distFilePath)
	faviconPath := path.Join(distFilePath, "/favicon.png")
	entryHtmlPath := path.Join(distFilePath, "/index.html")
	routerLogger.Debugf("The UI static file path: %s", staticFilePath)
	routerLogger.Debugf("The UI favicon.png path: %s", faviconPath)
	routerLogger.Debugf("The UI html path: %s", entryHtmlPath)

	router.Static("/dist", staticFilePath)
	//router.StaticFile("/favicon.png", faviconPath)
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
	userv1Group := v1Group.Group("/users")
	users := handler.NewUsers(logger)
	{
		userv1Group.GET("/:username", users.Get())
		userv1Group.POST("", users.Login())
	}

	//about us handler
	aboutUsv1Group := v1Group.Group("/about-us")
	aboutUs := handler.NewAboutUs(logger)
	{
		aboutUsv1Group.GET("", aboutUs.List())
		aboutUsv1Group.POST("", aboutUs.Create())
		aboutUsv1Group.PUT("/:id", aboutUs.Update())
		aboutUsv1Group.DELETE("/:id", aboutUs.Delete())
	}

	//products handler
	productsv1Group := v1Group.Group("/products")
	products := handler.NewProducts(logger)
	{
		productsv1Group.GET("", products.List())
		productsv1Group.POST("", products.Create())
		productsv1Group.PUT("/:id", products.Update())
		productsv1Group.DELETE("/:id", products.Delete())

	}

	//partnering handler
	partneringv1Group := v1Group.Group("/partnering")
	partnering := handler.NewPartnering(logger)
	{
		partneringv1Group.GET("", partnering.List())
		partneringv1Group.PUT("/:id", partnering.Update())
		partneringv1Group.DELETE("/:id", partnering.Delete())
		partneringv1Group.POST("", partnering.Create())
	}

	//resources handler
	resourcesv1Group := v1Group.Group("/resources")
	resources := handler.NewResources(logger)
	{
		resourcesv1Group.GET("", resources.List())
		resourcesv1Group.POST("", resources.Create())
		resourcesv1Group.PUT("/:id", resources.Update())
		resourcesv1Group.DELETE("/:id", resources.Delete())
	}

	//news handler
	newsv1Group := v1Group.Group("/news")
	news := handler.NewNews(logger)
	{
		newsv1Group.GET("", news.List())
		newsv1Group.PUT("/:id",news.Update())
		newsv1Group.DELETE("/:id", news.Delete())
		newsv1Group.POST("", news.Create())
	}

	//careers handler
	careersv1Group := v1Group.Group("/careers")
	careers := handler.NewCareers(logger)
	{
		careersv1Group.GET("", careers.List())
		careersv1Group.PUT("/:id",careers.Update())
		careersv1Group.DELETE("/:id", careers.Delete())
		careersv1Group.POST("", careers.Create())
	}

	//contact us handler
	contactUsv1Group := v1Group.Group("/contact-us")
	contactUs := handler.NewContactUs(logger)
	{
		contactUsv1Group.POST("", contactUs.Create())
	}

	////load html
	router.Use().StaticFile("/", entryHtmlPath)
	return router
}

