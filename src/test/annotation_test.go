package test

import (
	"errors"
	"fmt"
	"github.com/niceSong/goCicada/src/exception"
	"github.com/sirupsen/logrus"
	"net/http"
	"testing"
)

func TestAnnotation(t *testing.T) {
	l := logrus.New()
	entry := l.WithFields(logrus.Fields{"component": "UDM", "category": "HTTP"})
	exception.CicadaScan("src/obj")
	throwable := exception.CicadaErrorMap["AuthenticationError"](entry, http.StatusBadRequest, errors.New("haha"))
	fmt.Println(throwable)
}
