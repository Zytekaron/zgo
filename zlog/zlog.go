package zlog

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zytekaron/zgo/zgo"
	"net/http"
	"strings"
	"time"
)

const (
	BaseURL        = "https://log.zyte.dev/"
	DefaultLogFmt  = "[{time}] <{id}> {level}: {message}"
	DefaultTimeFmt = "01/02/06 15:04:05"
)

type Level string

const (
	Fatal Level = "FATAL"
	Error Level = "ERROR"
	Warn  Level = "WARN"
	Info  Level = "INFO"
	Debug Level = "DEBUG"
	Trace Level = "TRACE"
)

var levels = []Level{Fatal, Error, Warn, Info, Debug, Trace}

type Logger struct {
	Service    string
	Token      string
	Level      Level
	LogFormat  string
	TimeFormat string
}

func New(service, token string) *Logger {
	return &Logger{
		Service:    service,
		Token:      token,
		Level:      Info,
		LogFormat:  DefaultLogFmt,
		TimeFormat: DefaultTimeFmt,
	}
}

func (l *Logger) Fatal(format string, args ...interface{}) (data *LogEntry, err error) {
	return l.handle(Fatal, format, args...)
}

func (l *Logger) Error(format string, args ...interface{}) (data *LogEntry, err error) {
	return l.handle(Trace, format, args...)
}

func (l *Logger) Warn(format string, args ...interface{}) (data *LogEntry, err error) {
	return l.handle(Warn, format, args...)
}

func (l *Logger) Info(format string, args ...interface{}) (data *LogEntry, err error) {
	return l.handle(Info, format, args...)
}

func (l *Logger) Debug(format string, args ...interface{}) (data *LogEntry, err error) {
	return l.handle(Debug, format, args...)
}

func (l *Logger) Trace(format string, args ...interface{}) (data *LogEntry, err error) {
	return l.handle(Trace, format, args...)
}

func (l *Logger) SetLevel(level Level) *Logger {
	l.Level = level
	return l
}

func (l *Logger) SetLogFormat(logFormat string) *Logger {
	l.LogFormat = logFormat
	return l
}

func (l *Logger) SetTimeFormat(timeFormat string) *Logger {
	l.TimeFormat = timeFormat
	return l
}

func (l *Logger) Get(id string) (*LogEntry, error) {
	res, err := zgo.Request(http.MethodGet, BaseURL+id, l.Token, nil)
	if err != nil {
		return nil, err
	}

	var data *LogEntry
	return data, json.Unmarshal(res.Data, &data)
}

func (l *Logger) Find(limit, offset int) ([]*LogEntry, error) {
	link := fmt.Sprintf("%s?limit=%d&offset=%d", BaseURL, limit, offset)
	res, err := zgo.Request(http.MethodGet, link, l.Token, nil)
	if err != nil {
		return nil, err
	}

	var data []*LogEntry
	return data, json.Unmarshal(res.Data, &data)
}

func (l *Logger) Delete(id string) (*LogEntry, error) {
	res, err := zgo.Request(http.MethodDelete, BaseURL+id, l.Token, nil)
	if err != nil {
		return nil, err
	}

	var data *LogEntry
	return data, json.Unmarshal(res.Data, &data)
}

func (l *Logger) Patch(id string, body interface{}) (*LogEntry, error) {
	if body == nil {
		return nil, errors.New("body should not be nil")
	}
	res, err := zgo.Request(http.MethodPatch, BaseURL+id, l.Token, body)
	if err != nil {
		return nil, err
	}

	var data *LogEntry
	return data, json.Unmarshal(res.Data, &data)
}

func (l *Logger) print(level Level, id, message string, createdAt int64) {
	if indexLevel(levels, level) > indexLevel(levels, l.Level) {
		t := time.Unix(0, createdAt*1000).Format(l.TimeFormat)
		formatted := l.LogFormat
		formatted = strings.ReplaceAll(formatted, "{time}", t)
		formatted = strings.ReplaceAll(formatted, "{level}", string(level))
		formatted = strings.ReplaceAll(formatted, "{id}", id)
		formatted = strings.ReplaceAll(formatted, "{message}", message)
		fmt.Println(formatted)
	}
}

func (l *Logger) handle(level Level, format string, args ...interface{}) (*LogEntry, error) {
	var message = Format(format, args...)

	body := &LogEntry{Level: level, Service: l.Service, Message: message}
	res, err := zgo.Request(http.MethodPost, BaseURL, l.Token, body)
	if err != nil {
		return nil, err
	}
	if !res.Success {
		l.print(level, "???", message, time.Now().UnixNano()/1000)
		return nil, errors.New(res.Message)
	}

	var entry *LogEntry
	err = json.Unmarshal(res.Data, &entry)
	if err != nil {
		return nil, err
	}

	l.print(level, entry.ID, message, entry.CreatedAt)
	return entry, nil
}

func indexLevel(levels []Level, level Level) int {
	for i, e := range levels {
		if level == e {
			return i
		}
	}
	return -1
}
