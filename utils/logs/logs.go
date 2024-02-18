package logs

import "github.com/sirupsen/logrus"

func Info(fields map[string]interface{}, message string) {
	logrus.WithFields(fields).Info(message)
}

func Warning(fields map[string]interface{}, message string) {
	logrus.WithFields(fields).Warn(message)
}

func Deubg(fields map[string]interface{}, message string) {
	logrus.WithFields(fields).Debug(message)
}

func Error(fields map[string]interface{}, message string) {
	logrus.WithFields(fields).Error(message)
}

func Fatal(fields map[string]interface{}, message string) {
	logrus.WithFields(fields).Fatal(message)
}
