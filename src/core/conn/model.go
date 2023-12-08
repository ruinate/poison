// Package conn -----------------------------
// @file      : model.go
// @author    : fzf
// @contact   : fzf54122@163.com
// @time      : 2023/12/8 下午1:19
// -------------------------------------------
package conn

import "poison/src/model"

type LayerModel interface {
	init() model.Messages
	Send() model.Messages
}
