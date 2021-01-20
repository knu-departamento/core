package logging

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"os"
)

func SetupLog() error {
	errorFile, err := os.OpenFile("logs/errors.log", os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		return err
	}

	infoFile, err := os.OpenFile("logs/info.log", os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		return err
	}

	logrus.SetOutput(ioutil.Discard)
	logrus.SetFormatter(&nested.Formatter{})
	logrus.AddHook(&WriterHook{ // Send logrus with level higher than warning to stderr
		Writer: io.MultiWriter(os.Stderr, errorFile),
		LogLevels: []logrus.Level{
			logrus.PanicLevel,
			logrus.FatalLevel,
			logrus.ErrorLevel,
			logrus.WarnLevel,
		},
	})
	logrus.AddHook(&WriterHook{ // Send info and debug logrus to stdout
		Writer: io.MultiWriter(os.Stdout, infoFile),
		LogLevels: []logrus.Level{
			logrus.InfoLevel,
			logrus.DebugLevel,
		},
	})
	return nil
}
