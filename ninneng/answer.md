# 1

```proto
syntax = "proto3";

option go_package = "/;ans1";

message LoginReq{
  string username = 1;
  string password = 2;
}

message AccountInfo{
  string account_id = 1;
  string username = 2;
  //  ... othters
}

message LoginResp{
  int32 code = 1;
  string errMsg = 2;
  AccountInfo account_info = 3;
}

service AuthService{
  rpc Login(LoginReq)  returns(LoginResp);
}
```

# 2


# 3

```sql
select sale_date_year, sale_date_month, category_id, sum(units_sold) total_num
from sales_data
group by category_id, sale_date_year, sale_date_month
order by sale_date_year, sale_date_month, category_id

```
优化：
1. 创建复合索引，可提高按月份统计得查询效率
2. 分区表，按月份进行分区
3. 每隔一段时间，将一部分旧数据进行搬迁，归档到历史表中，减少主表得大小
4. 新建表对每天每个商品得总和进行统计和存储，从这里算更快，同时也可为每天的情况提供一个支持

# 4

方法1：日记记录谁请求了什么接口，

方法2：用redis记录用户对某一个接口最近的的一次请求时间，有新请求到来时，先判断，再更新

# 5

相当于 确保 meta 实现了Meta接口

# 6

sove1、2、3对应三个小问

```go
package main

import "fmt"

var source = []string{
	"APTZvA", "BddOIt", "ctuuYn", "BCd5js", "cVCuqR", "AQynrL", "AoZ62r", "BV9DXI", "cqkYj7", "ALSKpF", "CEkB4M", "By6jE3", "Aclr2o", "cLiix5", "AClM5o", "BN36oa", "BYj4K0", "cKtPyI", "BGOn7c", "BQreVu", "B7kQ15", "BHhAY0", "cbQBTI", "A2KDsf", "AwmbeJ", "BsNdy0", "BoIVCB", "C3pHMS", "CP9Wc6", "C6vyPb", "A6BTpf", "AguFNY", "AoeaF8", "AyQ3dP", "CzlhVY", "BkFrls", "C4WncK", "ASTebw", "CTpdJi", "BtGzKA", "cWtmeT", "BgLz5G", "A9Ohfh", "ASv3qg", "A4du4s", "BstIGr", "BSIkmq", "CKxdNR", "BgCF6g", "CWkjqZ",
}

func solve1() {
	source2 := make([]string, len(source))
	for i := 0; i < len(source); i++ {
		source2[i] = source[i][1:len(source[i])]
	}
	for _, v := range source2 {
		fmt.Println(v)
	}
}
func solve2() {
	for i := 0; i < len(source); i++ {
		source[i] = source[i][:3] + "A" + source[i][4:]
	}
	for _, v := range source {
		fmt.Println(v)
	}
}
func solve3() {
	dic := make(map[rune]int)
	for i := 0; i < len(source); i++ {
		dic[rune(source[i][0])]++
	}
	for k, v := range dic {
		fmt.Println(string(k), v)
	}
}

func main() {
	solve3()
}

```

# 7

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	c := &sync.WaitGroup{}
	ss := []string{"A", "B", "C"}
	ch := make([]chan struct{}, 3)
	for i := 0; i < len(ch); i++ {
		ch[i] = make(chan struct{})
	}
	for it, vv := range ss {
		cc := it
		v := vv
		c.Add(1)
		go func() {
			// 收到消息后可发送
			for i := 0; i < 5; i++ {
				//fmt.Println(cc)
				<-ch[cc]
				fmt.Println(v)
				ch[(cc+1)%3] <- struct{}{}
			}
			if cc == 0 {
				// 把三号的最后一次信号接受了，防止阻塞
				<-ch[0]
			}
			c.Done()
		}()
	}
	ch[0] <- struct{}{}
	c.Wait()
}




```

# 8
```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		log.Println("get request ")
		fmt.Fprintf(w, "Hello, get")
	} else if r.Method == "POST" {
		log.Println("post request ")
		fmt.Fprintf(w, "Hello, post")

	} else {
		log.Println("other request")
		fmt.Fprintf(w, "Hello")
	}
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8080", nil)
}

```

# 9

solve1使用管道，当一个id已经在执行时，相同的id再来会向全局map网对于id下添加管道，第一个id执行完了之后像这些id发消息，能通过测试

solve2使用redis发布订阅机制，本来以为这个会更合适，但耗时较长，过不了测试
```go
package main

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"
)

/*
*
你需要实现的目标函数 target

@param id 是一个随机字符串，例如 6A10A467-2842-A460-5353-DBE7D41986B7；
@param job 函数是一个耗时操作，例如：去数据库 query 数据，可能耗时 500ms；
@return count 表示在执行本次 job 期间有多少相同的 id 调用过 target

关键特性：相同 id 并发调用 target，target 只执行一次 job 函数，举例来说：
第一个线程传入 id 为 "id-123" 调用 target，job 函数开始执行，在此期间，又有其他 4 个线程以 id 为 "id-123" 调用了 target；
在此期间，只有一个 job 函数执行，等它执行完成后，上述 5 个线程均收到返回值 count=5，表示这段时间有 5 个相同 id 进行了调用；
*/
var ss = make(map[string][]chan int)
var ssMutex = &sync.Mutex{}

var ss2 = make(map[string]int)
var ssMutex2 = &sync.Mutex{}

func solve2(id string, job func()) (count int) {
	ssMutex.Lock()
	_, ok := ss2[id]

	if ok {
		ss2[id] += 1
		ssMutex.Unlock()
		pubSub := rds.Subscribe(context.Background(), id)
		defer pubSub.Close()
		msg, err := pubSub.ReceiveMessage(context.Background())
		if err != nil {
			panic("redis::" + err.Error())
		}
		//log.Fatalln(msg.Payload)
		count, err = strconv.Atoi(msg.Payload)
		if err != nil {
			panic(err)
		}
		//fmt.Println("###############count ", count)

	} else {
		ss2[id] = 1
		ssMutex.Unlock()
		// ok为false，那么就是第一次访问
		job()
		ssMutex.Lock()
		count = ss2[id]
		rds.Publish(context.Background(), id, count)
		delete(ss2, id)
		ssMutex.Unlock()
	}
	return count
}

func solve1(id string, job func()) (count int) {

	ssMutex.Lock()
	_, ok := ss[id]

	if ok {
		c := make(chan int)
		defer close(c)
		ss[id] = append(ss[id], c)
		ssMutex.Unlock()
		count = <-c
		//fmt.Println("###############count ", count)

	} else {
		ss[id] = nil
		ssMutex.Unlock()
		// ok为false，那么就是第一次访问
		job()
		ssMutex.Lock()
		count = len(ss[id]) + 1
		for i := 0; i < len(ss[id]); i++ {
			ss[id][i] <- count
		}
		delete(ss, id)
		ssMutex.Unlock()
	}
	return count
}
func target(id string, job func()) (count int) {
	//TODO implement this

	return solve1(id, job)
}

//////////////////////////////////////////////
///////// 接下来的代码为测试代码，请勿修改 /////////
//////////////////////////////////////////////

// 用来模拟 job 函数的变量
// 不要修改
var (
	counter     int
	counterLock sync.Mutex
)

// 用来模拟耗时，时间不固定，实现 target 时不能依赖此时间
// 不要修改
const (
	mockJobTimeout = 300 * time.Millisecond
	tolerate       = 30 * time.Millisecond
)

// 测试用的 job 函数，是一个计数器，用来模拟耗时操作
// 不要修改
func mockJob() {
	time.Sleep(mockJobTimeout)
	counterLock.Lock()
	counter++
	counterLock.Unlock()
}

// 相同 id 并行调用
// 不要修改
func testCaseSampleIdParallel() {
	counter = 0 //重置计数器
	const (
		id     = "CBD225E1-B7D9-BE76-9735-1D0A9B62EE4D"
		repeat = 5 //用来模拟相同 id 的多次重复调用，调用次数不固定，实现 target 时不能依赖此调用次数
	)
	wg := sync.WaitGroup{}
	wg.Add(repeat)
	tStart := time.Now()
	for i := 0; i < repeat; i++ {
		go func() {
			count := target(id, mockJob)
			wg.Done()
			if count != repeat {
				panic(fmt.Sprintln("[testCaseSampleIdConcurrence] count:", count, "!= repeat:", repeat))
			}
		}()
	}
	wg.Wait()
	if counter != 1 { //应该只调用了一次 job 函数
		panic(fmt.Sprintln("[testCaseSampleIdConcurrence] counter:", counter, "!= 1"))
	}
	var (
		tDelta  = time.Now().Sub(tStart)
		tExpect = mockJobTimeout + tolerate
	)
	if tDelta > tExpect {
		panic(fmt.Sprintln("[testCaseRandomId] timeout", tDelta, ">", tExpect))
	}
}

// 相同 id 串行调用
// 不要修改
func testCaseSampleIdSerial() {
	counter = 0
	const (
		id     = "3E5A5C8D-B254-383B-4F33-F6927578FD11"
		repeat = 2
	)
	tStart := time.Now()
	for i := 0; i < repeat; i++ {
		count := target(id, mockJob)
		if count != 1 {
			panic(fmt.Sprintln("[testCaseSampleIdSerial] count:", count, "!= 1"))
		}
	}
	if counter != repeat { //虽然是相同 id，但因为是串行调用，应该执行 repeat 次 job 函数
		panic(fmt.Sprintln("[testCaseSampleIdSerial] counter:", counter, "!= repeat:", repeat))
	}
	var (
		tDelta  = time.Now().Sub(tStart)
		tExpect = repeat*mockJobTimeout + tolerate
	)
	if tDelta > tExpect {
		panic(fmt.Sprintln("[testCaseSampleIdSerial] timeout", tDelta, ">", tExpect))
	}
}

// 不同 id 并行调用
// 不要修改
func testCaseRandomId() {
	counter = 0 //重置计数器
	ids := []string{
		"id-3",
		"id-3",
		"id-3",

		"id-2",
		"id-2",

		"id-1",
	}
	wg := sync.WaitGroup{}
	wg.Add(len(ids))
	tStart := time.Now()
	for _, id := range ids {
		id := id
		go func() {
			count := target(id, mockJob)
			wg.Done()
			var expectedCount int
			switch id {
			case "id-1":
				expectedCount = 1
			case "id-2":
				expectedCount = 2
			case "id-3":
				expectedCount = 3
			}
			if count != expectedCount {
				panic(fmt.Sprintln("[testCaseRandomId] count:", count, "!= expectedCount:", expectedCount, "id:", id))
			}
		}()
	}
	wg.Wait()
	if counter != 3 { //3个不同的 id 同时并发调用，job 函数应该执行 3 次
		panic(fmt.Sprintln("[testCaseSampleIdConcurrence] counter:", counter, "!= 3"))
	}
	var (
		tDelta  = time.Now().Sub(tStart)
		tExpect = 3*mockJobTimeout + tolerate
	)
	if tDelta > tExpect {
		panic(fmt.Sprintln("[testCaseRandomId] timeout", tDelta, ">", tExpect))
	}
}

// 不要修改
func main() {
	const repeat = 50
	for i := 0; i < repeat; i++ {
		testCaseSampleIdParallel()
		testCaseSampleIdSerial()
		testCaseRandomId()
		fmt.Print("\r", i+1, "/", repeat, " ✔ ")
	}
	fmt.Println("🎉 All Tests Passed!")
}

```