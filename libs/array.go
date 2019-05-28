package libs

// 判断value是否在array中
func In_array(value interface{}, array []interface{})bool {
	for _,val := range array {
		if val == value{
			return true
		}
	}
	return false
}
