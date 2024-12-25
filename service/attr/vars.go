package attr

const (
	ATTR_ID_PENGUINS = iota + 1
)

const (
	ATTR_TYPE_TEXT     = iota // 纯文本值
	ATTR_TYPE_BASIC           // 基础白值
	ATTR_TYPE_ADDITION        // 加算数值
	ATTR_TYPE_FACTOR          // 乘算系数
)
