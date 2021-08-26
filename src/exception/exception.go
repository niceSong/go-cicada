package exception

import "github.com/sirupsen/logrus"

type errorHandler func(logger *logrus.Entry, httpStatus int32, err ...error) Throwable

var CicadaErrorMap = make(map[string]errorHandler)
