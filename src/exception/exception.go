package exception

import "github.com/sirupsen/logrus"

type errorHandler func(logger *logrus.Entry, httpStatus int32, err ...error) Throwable

var Eb5gcErrorMap = make(map[string]errorHandler)
