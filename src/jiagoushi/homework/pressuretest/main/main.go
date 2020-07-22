package main

import (
	"fmt"
	"jiagoushi/homework/pressuretest"
)

func main() {
	doHttpTest4Baidu(100, 10)
}
func doHttpTest4Baidu(testCount, concurrentSize int) {
	averageSpan, p95Sapan, failCount := pressuretest.DoHttpPressureTest("http://www.baidu.com", testCount, concurrentSize)
	s := fmt.Sprintf("并发数%d,平均响应时间%d ms,p95响应时间 %d ms,总请求数%d 总失败数%d", concurrentSize, averageSpan, p95Sapan, testCount, failCount)
	fmt.Println(s)
}

//func TestReady(){
//	var wg sync.WaitGroup
//
//	ready:=make(chan struct{})
//	for i:=0;i<100;i++{
//		wg.Add(1)
//		go func(j int) {
//			defer wg.Done()
//			fmt.Println(j," :ready")
//			<-ready
//			fmt.Println(j," running ...")
//		}(i)
//	}
//	time.Sleep(time.Second)
//	close(ready)
//	wg.Wait()
//}
