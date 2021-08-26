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
	exception.AnnotationHandler("src/obj")
	throwable := exception.Eb5gcErrorMap["AuthenticationError"](errors.New("haha"), entry, http.StatusBadRequest)
	fmt.Println(throwable)
}