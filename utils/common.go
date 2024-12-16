package utils

type BizErr string

func (e BizErr) Error() string {
	return string(e)
}
