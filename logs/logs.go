package hook

import (
	"bytes"
	"errors"
	"exec/config"
	conf "exec/config"
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type TraceIdHook struct {
	TraceId  string
}

func (hook *TraceIdHook) Fire(entry *logrus.Entry) error {
	entry.Data["traceId"] = hook.TraceId
	return nil
}

func (hook *TraceIdHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

type LogFormatter struct{}

func (s *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("2006/01/02 15:04:05")
	var file string
	var len int
	if entry.Caller != nil {
		file = filepath.Base(entry.Caller.File)
		len = entry.Caller.Line
	}
	//fmt.Println(entry.Data)
	var conf *config.Config
	config, err := conf.LogConfig()
	if err != nil {
		logrus.WithFields(logrus.Fields{
		}).Info(err.Error())
	}
	msg := fmt.Sprintf("%s [GOID:%d] [%s] [%s] [%s:%d] %s\n",timestamp, getGID(), config.System.ServiceName, strings.ToUpper(entry.Level.String()),  file, len, entry.Message)
	return []byte(msg), nil
}

func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

type logFileWriter struct {
	file     *os.File
	logPath  string
	fileDate string //判断日期切换目录
	appName  string
	encoding string
}

func (p *logFileWriter) Write(data []byte) (n int, err error) {
	if p == nil {
		return 0, errors.New("logFileWriter is nil")
	}
	if p.file == nil {
		return 0, errors.New("file not opened")
	}
	//判断是否需要切换日期
	fileDate := time.Now().Format("20060102")
		filename := fmt.Sprintf("%s/%s-%s.log", p.logPath,  p.appName, fileDate)

		p.file, err = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_SYNC, 0666)
		if err != nil {
			return 0, err
		}
	n, err = p.file.Write(data)
	return n, err
}

func SetLogLevel() error {
	var conf *conf.Config
	config2, err := conf.LogConfig()
	if err !=nil {
		logrus.WithFields(logrus.Fields{
		}).Fatal(err.Error())
	}
	level := config2.SiteLogs.LogLevel
	if level == "info" {
		logrus.SetLevel(logrus.InfoLevel)
	} else if level == "debug" || level == "DEBUG" {
		logrus.SetLevel(logrus.DebugLevel)
	} else if level == "trace" || level == "TRACE"{
		logrus.SetLevel(logrus.TraceLevel)
	} else if level == "fatal" || level == "FATAL" {
		logrus.SetLevel(logrus.FatalLevel)
	} else if level == "error" || level == "ERROR" {
		logrus.SetLevel(logrus.ErrorLevel)
	} else if level == "warn" || level == "WARN" {
		logrus.SetLevel(logrus.WarnLevel)
	} else if level == "panic" || level == "PANIC" {
		logrus.SetLevel(logrus.PanicLevel)
	} else {
		logrus.WithFields(logrus.Fields{
		}).Fatal("Wrong Log level Configed...")
	}
	return nil
}

func InitLog(logPath string) {
	writer, _ := rotatelogs.New(
		logPath + ".%Y%m%d%H%M",
		rotatelogs.WithLinkName(logPath),
		rotatelogs.WithMaxAge(time.Duration(360)*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
	)
	SetLogLevel()
	logrus.SetOutput(writer)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(new(LogFormatter))
}