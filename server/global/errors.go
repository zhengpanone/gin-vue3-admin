package global

import "errors"

var (
	ErrMenuRepetition = errors.New("存在重复名称,请修改名称")
)
