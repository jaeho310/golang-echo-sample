package common

//import (
//	"os"
//	"strings"
//
//	"github.com/jblim0125/golang-web-platform/model"
//	formatter "github.com/mobigen/gologger"
//	"github.com/sirupsen/logrus"
//)
//
//// Logger empty struct
//type Logger struct {
//	*logrus.Logger
//}
//
//// Logger log variable
//var l *Logger
//
//func init() {
//	l = &Logger{logrus.New()}
//	l.SetOutput(os.Stdout)
//	f := &formatter.Formatter{
//		TimestampFormat: "2006-01-02 15:04:05.000",
//		ShowFields:      true,
//	}
//	l.SetFormatter(f)
//	l.SetLevel(logrus.DebugLevel)
//	l.SetReportCaller(true)
//}
//
//// GetInstance return logger instance
//func (Logger) GetInstance() *Logger {
//	return l
//}
//
//// SetLogLevel set log level
//func (l *Logger) SetLogLevel(lv logrus.Level) {
//	switch lv {
//	case logrus.ErrorLevel:
//		l.SetLevel(lv)
//	case logrus.WarnLevel:
//		l.SetLevel(lv)
//	case logrus.InfoLevel:
//		l.SetLevel(lv)
//	case logrus.DebugLevel:
//		l.SetLevel(lv)
//	default:
//		l.Errorf("ERROR. Not Supported Log Level[ %d ]", lv)
//	}
//}
//
//// GetLogLevel get log level
//func (l *Logger) GetLogLevel() string {
//	text, _ := l.GetLevel().MarshalText()
//	return string(text)
//}
//
//// Start Print Start Banner
//func (l *Logger) Start() {
//	l.Errorf("%s", model.LINE90)
//	l.Errorf(" ")
//	l.Errorf("                         START. %s:%s-%s",
//		strings.ToUpper(model.Name), model.Version, model.BuildHash)
//	l.Errorf(" ")
//	l.Errorf("%90s", "Copyright(C) 2021 Mobigen Corporation.  ")
//	l.Errorf(" ")
//	l.Errorf("%s", model.LINE90)
//}
//
//// Shutdown Print Shutdown
//func (l *Logger) Shutdown() {
//	l.Errorf("%s", model.LINE90)
//	l.Errorf(" ")
//	l.Errorf("                        %s Bye Bye.", strings.ToUpper(model.Name))
//	l.Errorf(" ")
//	l.Errorf("%90s", "Copyright(C) 2021 Mobigen Corporation.  ")
//	l.Errorf(" ")
//	l.Errorf("%s", model.LINE90)
//}
//
//// GormPrint gorm log print function
//func (l *Logger) GormPrint(v ...interface{}) {
//	if v[0] == "sql" {
//		l.Debugf("[ ORM_SQL ] %s", v[3])
//	}
//	if v[0] == "log" {
//		l.Debugf("[ ORM_LOG ] %s", v[2])
//	}
//}
