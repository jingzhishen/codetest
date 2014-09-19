package main

import  (
	"fmt"
	"runtime"
	"time"
	"reflect"
	"errors"
	"sync"
	"io"
	"log"
	"net/http"
	"encoding/json"
)

type Integer int
type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
}
type Lesser interface{
	Less(b Integer) bool
}

func (a Integer) Less(b Integer) bool {
	return a<b
}

func (a *Integer) Add(b Integer){
	*a+=b
}

func Add(a int, b int) (ret int, err error) {
	if a < 0 || b < 0 { // 假设这个函数只支持两个非负数字的加法
		err= errors.New("Should be non-negative numbers!")
		return 
	}
	return a + b, nil
// 支持多重返回值
}


func MyPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an  int value")
		case string:
			fmt.Println(arg, "is an string value")
		case int64:
			fmt.Println(arg, "is an int64 value")
		default:
			fmt.Println(arg, "is an unknown value")
		}
	}
}
func test1(){
	var v1 int = 1
	var v2 int64 = 234
	var v3 string = "hello"
	var v4 float32 = 1.234

	var a Integer=1
	var b LessAdder=&a
	var b1 Lesser=a
	if a.Less(2){
		fmt.Println(a," Less 2")
	}
	if b1.Less(2){
		fmt.Println(a," Less 2")
	}
	MyPrintf(v1, v2, v3,v4,a,b )
	var t1,t2=Add(-1,1)
	if t2 ==nil {
		fmt.Println(t1,t2)
	}
	if a1,ok := b1.(Integer);ok{
		fmt.Println(a1,v2)
	}
	//var b1 LessAdder=a

	fmt.Println("Hello",runtime.Version());
}

var counter int =0
func Count(lock *sync.Mutex){
	lock.Lock()
	counter++
	fmt.Println(counter)
	lock.Unlock()
}
func test_goroutine1(){
	lock:=&sync.Mutex{}
	for i:=0;i<10;i++{
		go Count(lock)
	}
	for{
		lock.Lock()
		c:=counter
		lock.Unlock()
		runtime.Gosched()
		if c>=10{
			break
		}
	}
}

func Count01(ch chan int){
	counter++
	ch <-counter
	//ch <- 1 
	fmt.Println("Counting=>",counter)
}
func test_goroutine4(){
	chs:=make([]chan int ,10)
	for i:=0;i<10;i++{
		chs[i]=make(chan int)
		go Count01(chs[i])
	}

	for _,ch:=range(chs){
		<-ch
	}
	time.Sleep(5)
}

func test_select01(){
	ch:=make(chan int ,1)
	for{
		select{
		case ch<-0:
		case ch<-1:
		}
		i:=<-ch
		fmt.Println("Value received:",i)
		if counter>10 {
			break
		}
		counter++
	}
}


func hellohandler(w http.ResponseWriter,r *http.Request){
	io.WriteString(w,"Hello,world!")
}

func test_http1(){
	http.HandleFunc("/hello",hellohandler)
	err:=http.ListenAndServe(":8080",nil)
	if err!=nil {
		log.Fatal("ListenAndServe: ",err.Error())
	}
}

type Book struct{
	Title string
	Authors []string
	Publisher string
	IsPublished bool
	Price float32
}
func test_json01(){
	gobook :=Book{
		"Go program",
		[]string{"12","23","34"},
		"nationanchip",
		true,
		9.99}
	b,err:=json.Marshal(gobook)
	if err!=nil {
		fmt.Println("json book:",b)
	}else {
		fmt.Println("json book:",string(b))
	}
	var book Book
	err=json.Unmarshal(b,&book)
	if err!=nil {
		fmt.Println("json book:",err.Error())
	}else {
		fmt.Println("value:%v",book)
	}
}
func test_reflect(){
	var x float64 =3.4
	fmt.Println("type:",reflect.TypeOf(x))
	v :=reflect.ValueOf(x)
	fmt.Println("v type:",v.Type())
	fmt.Println("kind of float64:",v.Kind()==reflect.Float64)
	fmt.Println("value float:",v.Float())
	//fmt.Println("value int:",v.Int())
}
func main(){
	//test_http1()
	//test_goroutine1()
	//test_goroutine4()
	//test_select01()
	//test_json01()
	test_reflect()
}
