package go_file_util

import ( 
	"time"
)

// 定义函数类型
type Fn func() error

// 定时器中的成员
type MyTicker struct {
	MyTick *time.Ticker
	Runner Fn
}
type MyTimer struct {
	MyTime *time.Timer
	Runner Fn
}

func NewMyTick(interval int, f Fn) *MyTicker {
	return &MyTicker{
		MyTick: time.NewTicker(time.Duration(interval) * time.Second),
		Runner: f,
	}
}

// 一次性
func NewMyTimer(interval int, f Fn) *MyTimer {
	return &MyTimer{
		MyTime: time.NewTimer(time.Duration(interval) * time.Second),
		Runner: f,
	}
}

// 启动定时器需要执行的任务
func (t *MyTicker) Start() {
	for {
		select {
		case <-t.MyTick.C:
			t.Runner()
		}
	}
}

// 启动定时器需要执行的任务
func (t *MyTimer) Start() {
	// for {
	select {
	case <-t.MyTime.C:
		t.Runner()
	}
	// }
}

// func over() error {
// 	fmt.Println("token过期")
// 	return nil
// }

// func main() {
// 	t := NewMyTimer(2, over)
// 	t.Start()
// }
