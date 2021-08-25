package exception

import "github.com/sirupsen/logrus"

type errorHandler func(err error, logger *logrus.Entry, httpStatus int32) Throwable

var Eb5gcErrorMap = make(map[string]errorHandler)
