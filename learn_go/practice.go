package main

import
(
	"golang.org/x/tour/wc"
	"strings"
	"golang.org/x/tour/pic"
	"fmt"
	//"strconv"
	"time"
	"math"
	"github.com/Go-zh/tour/reader"
	"os"
	"io"
	"bytes"
	"image"
	"image/color"
	"golang.org/x/tour/tree"
	//"math/rand"
	"sync"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}
type ErrNegativeSqrt float64
func (e ErrNegativeSqrt) Error() string{
	if e <= 0{
		return fmt.Sprintf("	cannot Sqrt negative number: %f", e)	// %v 会陷入死循环
	}else{
		return ""
	}
}
func Sqrt(x float64) (float64, error) {
	var z = 1.0
	if x <= 0{
		return x, ErrNegativeSqrt(x)
	}
	for ;math.Abs(z-x) <= 0.000001;{
		z -= (z*z - x) / (2*z)
	}
	return z, nil
}
type MyReader struct{}
// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (mr MyReader) Read(buf []byte) (int, error){
	cnt := 0
	for i:= range buf{
		buf[i] = 'A'
		cnt++
	}
	return cnt, nil
}
func excercise_reader() {
	//r := strings.NewReader("Hello, Reader!")
	//
	//b := make([]byte, 8)
	//for {
	//	n, err := r.Read(b)
	//	fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
	//	fmt.Printf("b[:n] = %q\n", b[:n])
	//	if err == io.EOF {
	//		break
	//	}
	//}
	println("validate")
	reader.Validate(MyReader{})
}

func WordCount(s string) map[string]int {
	fields := strings.Fields(s)
	var m = make(map[string]int)
	for _, field := range fields{
		if cnt, ok := m[field]; ok == false{
			m[field] = 1
		} else{
			m[field] = cnt + 1
		}
	}
	return m
}
func Pic(dx, dy int) [][]uint8 {
	//var arr [][] uint8;
	arr := make([][]uint8,  dy)
	for i:=0; i < dy; i++ {
		arr[i] = make([]uint8, dx)
		for j:=0; j < dx; j++{
			arr[i][j] = uint8(i * j)
		}
	}
	return arr
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	////count 用于计数调用次数,只要考虑0和1,不知有没有其他更好的方法
	////sum1 和 sum2 用于储存 fib函数的两次和的值
	//var count,sum1,sum2 int= 0,1,1
	//return func() int {
	//	switch count {
	//	//0 和 1 fib=1
	//	case 0,1:
	//		count++
	//		//其他 fib(n) = fib(n-1) + fib(n-2)
	//	default:
	//		sum1,sum2 = sum2, sum1+sum2
	//	}
	//	return sum2
	//}

	a, b := -1, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}
type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ipAddr IPAddr) String() string  {
	//s := ""
	//for _, v := range ipAddr{
	//	strings.Join(string(v-''))
	//}
	//return strings.Join([]string(ipAddr), ".")
	return fmt.Sprintf("%d.%d.%d.%d", ipAddr[0], ipAddr[1], ipAddr[2], ipAddr[3])
	//return fmt.Sprintf("%s", strconv.Itoa(int(ipAddr[0])))
}
func stringer_practice() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}


type rot13Reader struct {
	r io.Reader
}

var ascii_uppercase = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var ascii_lowercase = []byte("abcdefghijklmnopqrstuvwxyz")
var ascii_uppercase_len = len(ascii_uppercase)
var ascii_lowercase_len = len(ascii_lowercase)

func rot13(b byte) byte {
	pos := bytes.IndexByte(ascii_uppercase, b)
	if pos != -1 {
		return ascii_uppercase[(pos+13) % ascii_uppercase_len]
	}
	pos = bytes.IndexByte(ascii_lowercase, b)
	if pos != -1 {
		return ascii_lowercase[(pos+13) % ascii_lowercase_len]
	}
	return b
}
func (r rot13Reader) Read(p []byte) (int, error){
	n, err := r.r.Read(p)
	for i := 0; i < n; i++ {
		p[i] = rot13(p[i])
	}
	return n, err
}
func rot13_exercise() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

func image_exercise() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())//(0,0)-(100,100)
	fmt.Println(m.At(0, 0).RGBA())//0 0 0 0
}

type Image struct{}

func (i Image) ColorModel() color.Model{
	return color.RGBAModel
}
func (i Image) Bounds() image.Rectangle{
	rt := image.Rectangle{image.Point{0, 0 }, image.Point{20, 20}}
	return rt
}
func (i Image) At(x, y int) color.Color{
	uint8x := uint8(x)
	return color.NYCbCrA{color.YCbCr{uint8x, uint8x, uint8x},uint8(y)}
}
func image_exercise2() {
	m := Image{}
	pic.ShowImage(m)
}


func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 将和送入 c
}

func goroutine1() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从 c 中接收
	fmt.Println(x, y, x+y)
}

func fibonacci_go(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func goroutine2() {
	c := make(chan int, 10)
	go fibonacci_go(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
func fibonacci3(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func goroutine3() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci3(c, quit)
}
func goroutine4() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}


//等价二叉树
//不同二叉树的叶节点上可以保存相同的值序列。
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}
// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int){
	//for{
	//	select {
	//		case ch <- t.Value:
	//			t = t.Left
	//			fmt.Println("walk %v", t.Value)
	//		default:
	//			t = t.Right
	//			fmt.Println("default")
	//	}
	//}
	//close(ch)

	//	t := tree.New(1); ch := make(chan int)
	// go Walk(t, ch) got this stdout:
		// walk %v 5
		// default
		// default
		// default
		// 10

	//fmt.Println("hhh: %d", t.Value)
	sendValue(t, ch)

	close(ch)
}
func sendValue(t *tree.Tree, ch chan int){
	if t != nil{
		if t.Left != nil{
			sendValue(t.Left, ch)
		}
		ch <- t.Value
		if t.Right != nil{
			sendValue(t.Right, ch)
		}
	}
}
// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool{
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i:= range ch1{
		if i != <- ch2{
			return false
		}
	}
	return true
}

func binarytree_go_test1() {
	//a := rand.Perm(10)
	//fmt.Println(a)
	t := tree.New(1)
	ch := make(chan int)
	go Walk(t, ch)

	for i:=0; i<10; i++{
		fmt.Println(<-ch)
	}

	b1 := Same(tree.New(1), tree.New(1))
	b2 := Same(tree.New(1), tree.New(2))
	println(b1, b2)
}

// SafeCounter 的并发使用是安全的。
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc 增加给定 key 的计数器的值。
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
	c.v[key]++
	c.mux.Unlock()
}

// Value 返回给定 key 的计数器的当前值。
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
	defer c.mux.Unlock()
	return c.v[key]
}

func mutex_test1() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}
	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}


type Fetcher interface {
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: 并行的抓取 URL。
	// TODO: 不重复抓取页面。
	// 下面并没有实现上面两种情况：å
	//if depth <= 0 {
	//	return
	//}
	//body, urls, err := fetcher.Fetch(url)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Printf("found: %s %q\n", url, body)
	//for _, u := range urls {
	//	Crawl(u, depth-1, fetcher)
	//}
	//return


	if depth <= 0 {
		return
	}
	type fetched_res struct {
		url string
		body string
		urls []string
		depth int
		err error
	}

	fetched_set := make(map[string]bool)
	fetched_ch := make(chan *fetched_res)
	fetch_routine := func(url string, depth int){
		body, urls, err := fetcher.Fetch(url)
		fetched_ch <- &fetched_res{url, body,urls, depth,err}
	}
	go fetch_routine(url, depth)
	for progress := 1; progress > 0; progress-- {
		res_ptr := <- fetched_ch
		if res_ptr.err != nil{
			fmt.Println(res_ptr.err)
			continue
		}
		fmt.Printf("found: %s %q\n", res_ptr.url, res_ptr.body)

		fetched_set[res_ptr.url] = true

		cur_depth := res_ptr.depth - 1
		if cur_depth > 0{
			for _, candidate := range res_ptr.urls{
				if ! fetched_set[candidate]{
					progress++
					go fetch_routine(candidate, cur_depth)
				}else{
					fmt.Printf("fetched already: %s\n", candidate)
					continue
				}
			}
		}
	}
}

func co_test() {
	Crawl("http://golang.org/", 4, fetcher)
	//Crawl("https://golang.org/", 4, fakeFetcher{})
}

// fakeFetcher 是返回若干结果的 Fetcher。
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher 是填充后的 fakeFetcher。
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}

func main() {
	// map
	wc.Test(WordCount)

	// make
	pic.Show(Pic)

	// closure
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

	//stringer built-in interface
	stringer_practice()


	// error built-in interface
	if err := run(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))


	// reader
	excercise_reader()

	//rot13 reader
	rot13_exercise()


	//image
	image_exercise()
	image_exercise2()


	//goroutine
	goroutine1()
	goroutine2()
	goroutine3()
	goroutine4()
	binarytree_go_test1()


	// mutex
	mutex_test1()

	//cocrutine crawler
	co_test()
}


