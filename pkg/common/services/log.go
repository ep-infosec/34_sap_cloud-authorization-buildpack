package services

type Logger interface {
	Info(format string, args ...interface{})
	Error(format string, args ...interface{})
}
