package kodo

const (
	Standard = 0
	LowUsage = 1
	Archive  = 2
)

func GetType(sType int) string {
	if sType == Standard {
		return "标准存储"
	} else if sType == LowUsage {
		return "低频存储"
	} else if sType == LowUsage {
		return "归档存储"
	}
	return ""
}
