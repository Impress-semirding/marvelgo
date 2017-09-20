package marvel

import (
	"testing"
	"fmt"
	"os"
	"image"
	"time"
	"reflect"
)

func Test_marvel(t *testing.T)  {
	fmt.Println(1)
}



func Test_loadIma(t *testing.T)  {
	path := "../uploadfile/2017-7/22/b2608701a802a47d9d55228d4801bded.png"
	fmt.Println("path:", path)
	file, err := os.Open(path)
	fmt.Println(err)
	if err != nil {
		return
	}


	fmt.Println("file:", file)
	defer file.Close()
	img, name, err := image.Decode(file)

	fmt.Println("img:", img, "name:", name, "err:", err)
	return
}

func GetDateMH(timestamp int64) string {
	if timestamp <= 0 {
		return ""
	}

	tm := time.Unix(timestamp, 0)

	fmt.Println("tm:", tm)
	return tm.Format("2006-01-02 15:04")
}

func TestTime(t *testing.T)  {
	time := time.Now().Unix()

	times := GetDateMH(time)

	fmt.Println(times)
}

func TestTime1(t *testing.T)  {

	// 获取时间戳
	timess := time.Now().Unix()

	times := time.Now()
	fmt.Println(times)
	// 获取月份
	fmt.Println(times.Month())
	// 返回现在的时间 h m s
	fmt.Println(times.Clock())
	// 返回星期几 Saturday
	fmt.Println(times.Weekday())
	// 时间戳
	fmt.Println(timess)
}

func TestMonth(t *testing.T)  {

	month := time.Month.String(1)
	wen := time.Wednesday.String()
	wend := time.Weekday.String(0)
	fmt.Println(month)
	fmt.Println(wen)
	fmt.Println(wend)
}

func TestArray(t *testing.T)  {
	arr := []string{
		"1",
		"2",
		"3",
	}

	arr = append(arr, "4")

}

var d = [...]string {
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
}

type WebMarvel struct {

}

/*
  TODO:如果结构体里面定义了 String 方法，直接输出对象时会自动调用该方法,相当于 Java 里面的 toString 方法
*/
func (w WebMarvel) String() string {
	return d[1]
}

func TestSt(t *testing.T)  {
	m := new(WebMarvel)
	fmt.Println(m)
}

func TestSecond(t *testing.T) {
	fmt.Println(time.Second)
	fmt.Println(time.Wednesday)
	fmt.Println(time.August)
	var a int64
	var b string
	//a = time.Hour
	b = time.Second.String()
	fmt.Println(a)
	fmt.Println(b)
}

func TestAddTime(t *testing.T)  {

	s := time.Now()
	fmt.Println(s)

	ss := s.Add(time.Hour)
	fmt.Println(ss)

	sss := s.Sub(ss)
	fmt.Println(sss)
}

type int664 int64

func TestType(t *testing.T)  {
	var a int664
	var b int64

	a = 100
	b = 100

	/*
		新类型不是原类型的别名，除拥有数据存储结构外，他们之间没有任何关系，
		不会持有原类型任何信息。除非目标类型是未命名类型，否则必须显示转换
	 */
	// 虽然都是 int64 类型，但是给了别名之后就不是同一个类型了
	fmt.Println(int664(b) == a) // 必须强转
	fmt.Println(reflect.TypeOf(a) == reflect.TypeOf(b)) // false

}

func TestAfter(t *testing.T)  {
	c := make(<-chan time.Time)

	fmt.Println("the 1")
	tc:=time.After(time.Second) //返回一个time.C这个管道，1秒(time.Second)后会在此管道中放入一个时间点(time.Now())

	fmt.Println(reflect.TypeOf(c), "   ", reflect.TypeOf(tc))

	//时间点记录的是放入管道那一刻的时间值
	fmt.Println("the 2")
	fmt.Println("the 3")
	<- tc   //阻塞中，直到取出tc管道里的数据
	fmt.Println("the 4")
}


func TestA(t *testing.T)  {
	// 报错：invalid operation: a <- 1 (send to receive-only type <-chan int)
	// 意思是 <-chan int 类型只支持读
	/*a := make(<-chan int, 1)
	b := <- a
	fmt.Println(b)*/

	// 下面这样就不会报错了
	// 如果 make(chan int) 后面不加缓冲区，则默认是阻塞的，下面的代码如果没有其他 goroutine的话会报错
	// fatal error: all goroutines are asleep - deadlock!
	a := make(chan int, 1)
	a <- 10 // 给 a 赋值
	//aa := AA()
	go Parse(a)
}
func Parse(ch <-chan int) {
	for value := range ch {
		fmt.Println("Parsing value", value)
	}
}

func TestChannel(t *testing.T) {
	count := make(chan int)
	go func(count chan int) {
		current := 0
		for {
			current = <-count
			current++
			count <- current
			fmt.Println(count)
		}
	}(count)
	count <- 1
	fmt.Println(<-count)
	//time.Sleep(time.Second*2)
}
