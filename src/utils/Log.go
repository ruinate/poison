// Package utils -----------------------------
// @file      : Log.go
// @author    : fzf
// @time      : 2023/4/20 上午10:57
// -------------------------------------------
package utils

import (
	"github.com/sirupsen/logrus"
)

func LogDebug(p *ProtoAPP, err error) {
	if err != nil {
		logrus.Errorf(err.Error())
	} else {
		logrus.Infof(p.Result)
	}
}
