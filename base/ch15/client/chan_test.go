package client

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)

	return "done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

func AsyncService() chan string {
	retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("return result")
		retCh <- ret
		fmt.Println("AsyncService done")
	}()

	return retCh
}

func TestAsync(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	t.Log(<-retCh)
	select {
		case ret := <-retCh:
			t.Log(ret)
		case <-time.After(time.Millisecond * 100):
			t.Error("time out")
		default:
			t.Log("default")
	}
}
