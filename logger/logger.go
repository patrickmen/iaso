package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gokits/rfw"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type accessEntity struct {
	Time      string
	Method    string
	Path      string
	Status    int
	Client    string
	UserAgent string
	ReqSize   int64
	RspSize   int
	Latency   int
}

type HttpLogger interface {
	HttpLogger() gin.HandlerFunc
}

type httpLogger struct {
	Logger *zap.SugaredLogger
}

var (
	runtimeLw           *rfw.Rfw
	runtimeSyncer       *MutableWriteSyncer

	// Runtime logger for runtime logging
	Runtime             *zap.Logger
	// RuntimeSugar sugar logger for runtime logging
	RuntimeSugar        *zap.SugaredLogger
	DevModeRuntimeSugar *zap.SugaredLogger

	EventSugar          *zap.SugaredLogger
	eventLw             *rfw.Rfw
	eventSyncer         *MutableWriteSyncer

	AuditSugar          *zap.SugaredLogger
	auditLw             *rfw.Rfw
	auditSyncer         *MutableWriteSyncer

	accessWriter        *rfw.Rfw
	accessPool        = sync.Pool{
		New: func() interface{} {
			return &accessEntity{
				Status:  0,
				ReqSize: -1,
				Latency: -1,
			}
		},
	}
)

func (ae *accessEntity) Reset() {
	ae.Status = 0
	ae.ReqSize = -1
	ae.RspSize = -1
}

// MutableWriteSyncer a WriteSyncer implementation support change inner WriteSyncer on the fly
type MutableWriteSyncer struct {
	syncer atomic.Value
}

func NewMutableWriteSyncer(defaultSyncer zapcore.WriteSyncer) *MutableWriteSyncer {
	mws := &MutableWriteSyncer{}
	mws.syncer.Store(&defaultSyncer)
	return mws
}

func (mws *MutableWriteSyncer) get() zapcore.WriteSyncer {
	return *(mws.syncer.Load().(*zapcore.WriteSyncer))
}

func (mws *MutableWriteSyncer) SetWriteSyncer(newSyncer zapcore.WriteSyncer) {
	mws.syncer.Store(&newSyncer)
}

func (mws *MutableWriteSyncer) Write(p []byte) (n int, err error) {
	return mws.get().Write(p)
}

func (mws *MutableWriteSyncer) Sync() error {
	return mws.get().Sync()
}

func init() {
	runtimeSyncer = NewMutableWriteSyncer(zapcore.Lock(zapcore.AddSync(os.Stdout)))
	jsonEncoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})

	// instead of message of level, ts and msag, audit log only print the message about the audit info
	auditJsonEncoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{ //only output the kv in the audit log
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})
	runtimeCore := zapcore.NewCore(
		jsonEncoder,
		runtimeSyncer,
		zapcore.InfoLevel,
	)
	Runtime = zap.New(runtimeCore, zap.AddCaller())
	RuntimeSugar = Runtime.Sugar()

	DevRuntimeCore := zapcore.NewCore(
		jsonEncoder,
		runtimeSyncer,
		zapcore.DebugLevel,
	)
	DevModeRuntimeSugar = zap.New(DevRuntimeCore, zap.AddCaller()).Sugar()

	eventSyncer = NewMutableWriteSyncer(zapcore.Lock(zapcore.AddSync(os.Stdout)))
	EventCore := zapcore.NewCore(
		jsonEncoder,
		eventSyncer,
		zapcore.DebugLevel,
	)
	EventSugar = zap.New(EventCore, zap.AddCaller()).Sugar()

	auditSyncer = NewMutableWriteSyncer(zapcore.Lock(zapcore.AddSync(os.Stdout)))
	AuditCore := zapcore.NewCore(
		auditJsonEncoder,
		auditSyncer,
		zapcore.InfoLevel,
	)
	AuditSugar = zap.New(AuditCore, zap.AddCaller()).Sugar()

}

func AccessLoggerForGin() gin.HandlerFunc {
	encoder := json.NewEncoder(accessWriter)
	return func(c *gin.Context) {
		ae := accessPool.Get().(*accessEntity)
		defer func() {
			accessPool.Put(ae)
		}()
		ae.Reset()
		ae.Method = c.Request.Method
		ae.Path = c.Request.URL.Path
		ae.Client = c.ClientIP()
		if c.Request.URL.RawQuery != "" {
			ae.Path += "?" + c.Request.URL.RawQuery
		}
		ae.ReqSize = c.Request.ContentLength
		start := time.Now()
		defer func() {
			end := time.Now()
			ae.Latency = int(end.Sub(start) / time.Millisecond)
			ae.Status = c.Writer.Status()
			ae.RspSize = c.Writer.Size()
			ae.Time = end.Format("2006/01/02 - 15:04:05.000")
			encoder.Encode(ae)
		}()
		// Process request
		c.Next()
	}
}

//Init initializer of this module
func Init(runtimeLog string, runtimeRemaindays int, accessLog string, accessRemaindays int) (err error) {
	if runtimeLog == "" && accessLog == "" {
		return fmt.Errorf("At least one of runtimeLog and accessLog must be specified")
	}

	if runtimeLog != "" {
		runtimeLw, err = rfw.NewWithOptions(runtimeLog, rfw.WithCleanUp(runtimeRemaindays))
		if err != nil {
			return fmt.Errorf("open rfw for path %s failed: %v", runtimeLog, err)
		}
		runtimeSyncer.SetWriteSyncer(zapcore.AddSync(runtimeLw))
	}
	if accessLog != "" {
		accessWriter, err = rfw.NewWithOptions(accessLog, rfw.WithCleanUp(accessRemaindays))
		if err != nil {
			runtimeSyncer.SetWriteSyncer(zapcore.AddSync(os.Stdout))
			return fmt.Errorf("open rfw for path %s failed: %v", accessLog, err)
		}
	}

	return nil
}

func NewHttpLogger(logger *zap.SugaredLogger) HttpLogger {
	return &httpLogger{
		Logger: logger.Named("HttpLogger"),
	}
}

type bodyWriter struct {
	gin.ResponseWriter
	bodyBuf *bytes.Buffer
}

//func (w *bodyWriter) Write(b []byte) (int, error) {
//	//memory copy here!
//	w.bodyBuf.Write(b)
//	return w.ResponseWriter.Write(b)
//}

func (r *httpLogger) HttpLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, err := ctx.GetRawData()
		if err != nil {
			r.Logger.Errorf("Fail to get raw data in the request")
		}
		r.Logger.Debugf("The request body: %s", string(data))
		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
		responseBodyWriter := &bodyWriter{
			bodyBuf:        bytes.NewBufferString(""),
			ResponseWriter: ctx.Writer}
		ctx.Writer = responseBodyWriter
		ctx.Next()
		r.Logger.Debug("The response body: %s", responseBodyWriter.bodyBuf.String())
	}
}

func EventLog(eventLog string, remaindays int) (err error) {
	eventLw, err = rfw.NewWithOptions(eventLog, rfw.WithCleanUp(remaindays))
	if err != nil {
		return fmt.Errorf("open rfw for path %s failed: %v", eventLog, err)
	}
	eventSyncer.SetWriteSyncer(zapcore.AddSync(eventLw))
	return nil
}

//Final finalizer of this module
func Final() {
	Runtime.Sync()
	runtimeLw.Close()

	accessWriter.Close()
}

