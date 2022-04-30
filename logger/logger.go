package logger

import "github.com/sirupsen/logrus"

// Public log var
var Log *logrus.Logger

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	Log = logrus.New()
}
