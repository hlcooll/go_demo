package main

/*
通道（channel）与并发goroutine搭配，实现用通信代替内存共享的CSP模型
*/

//消费者
func consumer(data chan int, done chan bool) {
	//接受数据，直到通道被关闭
	for x := range data {
		println("recv:", x)
	}
	//通知main,消费结束
	done <- true
}

//生产者
func producer(data chan int) {
	for i := 0; i < 4; i++ {
		//发送数据
		data <- i
	}
	//生产结束，关闭通道
	close(data)
}

func main() {
	//用于接收消费结束信号
	done := make(chan bool)
	//数据管道
	data := make(chan int)

	//启动消费者
	go consumer(data, done)
	//启动生产者
	go producer(data)

	//阻塞，直到消费者发回结束信号
	<-done
}
