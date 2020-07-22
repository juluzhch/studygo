package pressuretest

import (
	"fmt"
	"net/http"
	"runtime"
	"sort"
	"sync"
	"time"
)

func DoHttpTest(url string) (success bool, timespan int) {
	begin := GetCurrentTimestamp()
	if resp, err := http.Get(url); err != nil {
		return false, 0
	} else {
		defer resp.Body.Close()
		if resp.StatusCode == 200 {
			end := GetCurrentTimestamp()
			timespan = int(end - begin)
			return true, timespan
		}
		return false, 0
	}
}

//获取当前时间毫秒
func GetCurrentTimestamp() int64 {
	return time.Now().UnixNano() / 1e6
}

//输入url,请求总次数，并发数，反馈平均耗时，95%耗时，失败请求数（失败请求不计入耗时）
func DoHttpPressureTest(url string, testCount, concurrentSize int) (averageSpan, p95Span, failCount int) {
	//runtime.GOMAXPROCS(runtime.NumCPU())
	runtime.GOMAXPROCS(concurrentSize + 10)
	f := DoHttpTest
	ready := make(chan struct{})
	results := make(chan TestResult, concurrentSize)
	var wg sync.WaitGroup
	wg.Add(concurrentSize)
	pTestCount := testCount / concurrentSize //每个线程执行的测试次数
	for t := 0; t < concurrentSize; t++ {
		go func() {
			defer wg.Done()
			<-ready
			var failCount = 0
			var timeSpan []int
			for i := 0; i < pTestCount; i++ {
				if ok, span := f(url); ok {
					timeSpan = append(timeSpan, span)
				}
			}
			//收集测试结果
			result := TestResult{failCount: failCount, span: timeSpan}
			results <- result
		}()
	}
	time.Sleep(time.Second)
	close(ready)
	wg.Wait()
	close(results)
	//所有测试结果合并，统计
	var totalResul TestResult
	for result := range results {
		totalResul.add(&result)
	}
	return totalResul.GetSpanStatistic()
}

//测试结果记录
type TestResult struct {
	failCount int
	span      []int //所有成功请求耗时记录（ms ）
}

func (this *TestResult) add(newResult *TestResult) {
	this.failCount = newResult.failCount
	this.span = append(this.span, newResult.span...)
}

//耗时统计：平均耗时，95%耗时，失败记录数
func (this *TestResult) GetSpanStatistic() (averageSpan, p95Span, failCount int) {
	var successCount = len(this.span)
	sort.Ints(this.span)
	total := 0
	for _, s := range this.span {
		fmt.Println(s)
		total += s
	}
	averageSpan = total / successCount
	p95Count := successCount * 95 / 100
	if p95Count > 0 {
		p95Total := 0
		for i := 0; i < p95Count; i++ {
			p95Total += this.span[i]
		}
		p95Span = p95Total / p95Count
	}
	return averageSpan, p95Span, this.failCount
}
