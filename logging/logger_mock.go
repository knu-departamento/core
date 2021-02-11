package logging

import "github.com/sirupsen/logrus"

type LoggerMock struct{}

func (l LoggerMock) WithField(key string, value interface{}) *logrus.Entry {
	return nil
}

func (l LoggerMock) WithFields(fields logrus.Fields) *logrus.Entry {
	return nil
}

func (l LoggerMock) WithError(err error) *logrus.Entry {
	return nil
}

func (l LoggerMock) Debugf(format string, args ...interface{}) {
	return
}

func (l LoggerMock) Infof(format string, args ...interface{}) {
	return
}

func (l LoggerMock) Printf(format string, args ...interface{}) {
	return
}

func (l LoggerMock) Warnf(format string, args ...interface{}) {
	return
}

func (l LoggerMock) Warningf(format string, args ...interface{}) {
	return
}

func (l LoggerMock) Errorf(format string, args ...interface{}) {
	return
}

func (l LoggerMock) Fatalf(format string, args ...interface{}) {
	return
}

func (l LoggerMock) Panicf(format string, args ...interface{}) {
	return
}

func (l LoggerMock) Debug(args ...interface{}) {
	return
}

func (l LoggerMock) Info(args ...interface{}) {
	return
}

func (l LoggerMock) Print(args ...interface{}) {
	return
}

func (l LoggerMock) Warn(args ...interface{}) {
	return
}

func (l LoggerMock) Warning(args ...interface{}) {
	return
}

func (l LoggerMock) Error(args ...interface{}) {
	return
}

func (l LoggerMock) Fatal(args ...interface{}) {
	return
}

func (l LoggerMock) Panic(args ...interface{}) {
	return
}

func (l LoggerMock) Debugln(args ...interface{}) {
	return
}

func (l LoggerMock) Infoln(args ...interface{}) {
	return
}

func (l LoggerMock) Println(args ...interface{}) {
	return
}

func (l LoggerMock) Warnln(args ...interface{}) {
	return
}

func (l LoggerMock) Warningln(args ...interface{}) {
	return
}

func (l LoggerMock) Errorln(args ...interface{}) {
	return
}

func (l LoggerMock) Fatalln(args ...interface{}) {
	return
}

func (l LoggerMock) Panicln(args ...interface{}) {
	return
}
