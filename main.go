package main

import (
	"github.com/sirupsen/logrus"
	"github.com/zjyl1994/catchsdbot/startup"
)

func main() {
	if err := startup.Startup(); err != nil {
		logrus.Fatalln(err.Error())
	}
}
