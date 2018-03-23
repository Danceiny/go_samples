package main

import (
	"time"
	"log"
	"net/http"
)
/*
State 类型
State 类型表示一个URL的状态。

Poller 会将 State 值发送到 StateMonitor，它维护了每一个URL当前状态的映射。


Resource 类型
Resource 表示URL被轮询的状态，即URL本身及其最后一次成功轮询之后遇到的错误编号。

当此程序启动时，它会为每个URL都分配一个 Resource。主Go程与 Poller Go程会在信道上互相发送 Resource。

Poller 函数
每个 Poller 都会从输入信道中接收到 Resource 的指针。在此程序中，我们约定发送者通过信道， 将底层数据的所有权传递给接收者。由此可知，不会出现两个Go程同时访问该 Resource 的情况。这就意味着我们无需担心锁会阻止对这些数据结构的并发访问。

Poller 通过调用其 Poll 方法来处理 Resource。

它会向 status 信道发送 State 值，以此将 Poll 的结果通知给 StateMonitor。

最后，它会将 Resource 的指针发送给 out 信道。这可以理解成 Poller 说：“我搞定这个 Resource 了”，然后将它的所有权返回给主Go程。

多个Go程运行多个 Poller，可以并行地处理 Resource。


Poll 方法
（Resource 类型的）Poll 方法为 Resource 的URL执行HTTP HEAD请求，并返回HTTP响应的状态码。 若有错误产生，Poll 就会将该信息记录到标准错误中，并转而返回该错误的字符串。
 */
const (
	numPollers     = 2                // number of Poller goroutines to launch // Poller Go程的启动数
	pollInterval   = 60 * time.Second // how often to poll each URL            // 轮询每一个URL的频率
	statusInterval = 10 * time.Second // how often to log status to stdout     // 将状态记录到标准输出的频率
	errTimeout     = 10 * time.Second // back-off timeout on error             // 回退超时的错误
)

var urls = []string{
	"http://www.google.com/",
	"http://golang.org/",
	"http://blog.golang.org/",
}

// State represents the last-known state of a URL.

// State 表示一个URL最后的已知状态。
type State struct {
	url    string
	status string
}

// StateMonitor maintains a map that stores the state of the URLs being
// polled, and prints the current state every updateInterval nanoseconds.
// It returns a chan State to which resource state should be sent.

// StateMonitor 维护了一个映射，它存储了URL被轮询的状态，并每隔 updateInterval
// 纳秒打印出其当前的状态。它向资源状态的接收者返回一个 chan State。
func StateMonitor(updateInterval time.Duration) chan<- State {
	updates := make(chan State)
	urlStatus := make(map[string]string)
	ticker := time.NewTicker(updateInterval)
	go func() {
		for {
			select {
			case <-ticker.C:
				logState(urlStatus)
			case s := <-updates:
				urlStatus[s.url] = s.status
			}
		}
	}()
	return updates
}

// logState prints a state map.

// logState 打印出一个状态映射。
func logState(s map[string]string) {
	log.Println("Current state:")
	for k, v := range s {
		log.Printf(" %s %s", k, v)
	}
}

// Resource represents an HTTP URL to be polled by this program.

// Resource 表示一个被此程序轮询的HTTP URL。
type Resource struct {
	url      string
	errCount int
}

// Poll executes an HTTP HEAD request for url
// and returns the HTTP status string or an error string.

// Poll 为 url 执行一个HTTP HEAD请求，并返回HTTP的状态字符串或一个错误字符串。
func (r *Resource) Poll() string {
	resp, err := http.Head(r.url)
	if err != nil {
		log.Println("Error", r.url, err)
		r.errCount++
		return err.Error()
	}
	r.errCount = 0
	return resp.Status
}

// Sleep sleeps for an appropriate interval (dependent on error state)
// before sending the Resource to done.

// Sleep 在将 Resource 发送到 done 之前休眠一段适当的时间（取决于错误状态）。
func (r *Resource) Sleep(done chan<- *Resource) {
	time.Sleep(pollInterval + errTimeout*time.Duration(r.errCount))
	done <- r
}

func Poller(in <-chan *Resource, out chan<- *Resource, status chan<- State) {
	for r := range in {
		s := r.Poll()
		status <- State{r.url, s}
		out <- r
	}
}

func main() {
	// Create our input and output channels.
	// 创建我们的输入和输出信道。
	pending, complete := make(chan *Resource), make(chan *Resource)

	// Launch the StateMonitor.
	// 启动 StateMonitor。
	status := StateMonitor(statusInterval)

	// Launch some Poller goroutines.
	// 启动一些 Poller Go程。
	for i := 0; i < numPollers; i++ {
		go Poller(pending, complete, status)
	}

	// Send some Resources to the pending queue.
	// 将一些 Resource 发送至 pending 序列。
	go func() {
		for _, url := range urls {
			pending <- &Resource{url: url}
		}
	}()

	for r := range complete {
		go r.Sleep(pending)
	}
}
