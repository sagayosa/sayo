package sayolog

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var (
	infoLog *log.Logger
	errLog  *log.Logger
)

func init() {
	infoLog = log.New(os.Stdout, "[Info] ", log.LstdFlags|log.Llongfile)
	errLog = log.New(os.Stderr, "[Error] ", log.LstdFlags|log.Llongfile)
}

type Log struct {
	Message     string `json:"msg"`
	ErrorString string `json:"err"`
}

func (l *Log) Err(err error) *Log {
	l.ErrorString = err.Error()
	return l
}

func (l *Log) Msg(format string, a ...interface{}) *Log {
	l.Message = fmt.Sprintf(format, a...)
	return l
}

func (l *Log) toString() string {
	bts, err := json.Marshal(l)
	if err != nil {
		errLog.Printf("Log toString failed because of: %v", err)
	}

	return string(bts)
}

func (l *Log) Info() {
	infoLog.Println(l.toString())
}

func (l *Log) Error() {
	errLog.Println(l.toString())
}

func Err(err error) *Log {
	return &Log{
		ErrorString: err.Error(),
	}
}

func Msg(format string, a ...interface{}) *Log {
	return &Log{
		Message: fmt.Sprintf(format, a...),
	}
}
