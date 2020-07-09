package consistent_hashing

import "strconv"

///先验证下是否正确
//code值输出，
var nodePrefix = "192.168.1."

func GetMockNodeKey(size int) []string {
	var s []string
	for i := 0; i < size; i++ {
		s = append(s, nodePrefix+strconv.Itoa(i+1))
	}
	return s
}
func GetMockDataKey(size int) []string {
	var s []string
	for i := 0; i < size; i++ {
		s = append(s, "user-"+strconv.Itoa(i+1))
	}
	return s
}
