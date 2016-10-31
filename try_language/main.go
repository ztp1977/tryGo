package main

import (
	"errors"
	"github.com/Sirupsen/logrus"
)

func getType (input interface{}) error {
	switch v := input.(type) {
	case string:
		logrus.Debug(v)
		return nil
	case []byte:
		logrus.Debug(v)
		return nil
	default:
		return errors.New("not support")
	}
}

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	getType("this is string")
	getType([]byte("this is string"))

}
