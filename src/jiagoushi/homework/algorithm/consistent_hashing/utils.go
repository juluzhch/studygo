package consistent_hashing

import (
	"crypto/md5"
	"fmt"
	"hash/crc32"
	"math"
	"time"
)

//切片添加删除操作

func insertBefor(source []uint32, newValue uint32, insertIndex int) []uint32 {
	if len(source) == 0 {
		return append(source, newValue)
	}
	var newCodes []uint32
	if insertIndex == 0 { //头部
		//newCodes:=make([]uint32, len(source)+1)

		newCodes = append(append(newCodes, newValue), source...)
		return newCodes
	}
	if insertIndex > len(source) { //尾部
		return append(source, newValue)
	}
	newCodes = append(append(append(newCodes, source[:insertIndex]...), newValue), source[insertIndex:]...)
	return newCodes
}

func deleteAt(source []uint32, index int) []uint32 {
	return append(source[:index], source[index+1:]...)
}

//哈希算法-crc32(md5(key+md5(key)))
func getHashCode4KeyWithMd5(key string) uint32 {
	var newKey = key + Md5([]byte(key))
	return getHashCode(newKey)
}

//哈希算法-crc32(md5(key))
func getHashCode(key string) uint32 {
	return getCrcHash(Md5([]byte(key)))
}

//哈希算法-crc32(key)
func getCrcHash(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}

func StandDeviation(data []int) float64 {
	v := Variance(data)
	return math.Sqrt(v)

}
func Variance(data []int) float64 {
	var sum = Sum(data)
	var average = float64(sum) / float64(len(data))
	var total float64
	for _, v := range data {
		diff := float64(v) - average
		total += diff * diff
	}
	return total / float64(len(data))
}
func Sum(data []int) int64 {
	var sum int64
	for _, v := range data {
		sum += int64(v)
	}
	return sum
}
func Md5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}
func GetCurrentTimestamp() int64 {
	return time.Now().UnixNano() / 1e6
}

//func GetRandomString(lens int,sed) string {
//	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
//	bytes := []byte(str)
//	result := []byte{}
//	r := rand.New(rand.NewSource(time.Now().UnixNano()))
//	for i := 0; i < lens; i++ {
//		result = append(result, bytes[r.Intn(len(bytes))])
//	}
//	return string(result)
//}
