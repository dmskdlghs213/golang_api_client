package practice

import (
	"fmt"
	"sync"
	"time"
)


// // goroutine基本
// func main() {

// 	fmt.Println("main start")

// 	go Greet(i)

// 	time.Sleep(time.Second)
// 	fmt.Println("main end")

// 	end := time.Now();
// 	fmt.Printf("%f秒\n",(end.Sub(start)).Seconds())
// }

// func Greet(i int) {
// 	fmt.Println("hello", i)
// }


// goroutine時間比較

func Task1() {
	time.Sleep(2 * time.Second)
}

func Task2() {
	time.Sleep(3 * time.Second)
}

func Task3() {
	time.Sleep(2 * time.Second)
}

func concurrencyTask1(wg *sync.WaitGroup) {
	fmt.Println("Task1")
	time.Sleep(2 * time.Second)
	wg.Done()
}

func concurrencyTask2(wg *sync.WaitGroup) {
	fmt.Println("Task2")
	time.Sleep(3 * time.Second)
	wg.Done()
}

func concurrencyTask3(wg *sync.WaitGroup) {
	fmt.Println("Task3")
	time.Sleep(2 * time.Second)
	wg.Done()
}

func main() {
	// 	// 通常呼び出し
	// 	// start := time.Now()
	// 	// Task1()
	// 	// Task2()
	// 	// Task3()
	// 	// fmt.Println("done")

	// 	// end := time.Now()
	// 	// fmt.Printf("%f秒\n",(end.Sub(start)).Seconds())

	// goroutine使用
	start := time.Now()

	var wg sync.WaitGroup
	wg.Add(3)
	go concurrencyTask1(&wg)
	go concurrencyTask2(&wg)
	go concurrencyTask3(&wg)
	wg.Wait()
	fmt.Println("done")

	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}

//　受信するデータの数が分からない時
// func goroutine(ch chan int) {
//     for i := 0; i < 10; i++ {
// 		ch <- i
//     }
//     close(ch)
// }

// func main() {
//     ch := make(chan int)
//     go goroutine(ch)

//     for v := range ch {
// 		fmt.Println(v)
//     }
// }

// chanelの型
// func main() {
// 	// receiveStream := make(<-chan string)
// 	// sendStream := make(chan<- string)
// 	stream := make(chan int,4)

// 	// receiveStream = stream
// 	// sendStream = stream

// 	stream <- 1
// 	stream <- 2
// 	stream <- 3
// 	stream <- 4

// 	// sendStream <- "hello 5"
// 	// sendStream <- "hello 6"
// 	// sendStream <- "hello 7"
// 	// sendStream <- "hello 8"

// 	// キューから受信し、値を出力
// 	// キューに値がなければデッドロッグ

// 	// fmt.Println(<-receiveStream)
// 	// fmt.Println(<-receiveStream)
// 	// fmt.Println(<-receiveStream)
// 	// fmt.Println(<-receiveStream)

// 	fmt.Println(<-stream)
// 	fmt.Println(<-stream)
// 	fmt.Println(<-stream)
// 	fmt.Println("main end")
// }

// chanelのselect
// func main() {

// 	ch1 := make(chan int,1)
// 	ch2 := make(chan int,1)
// 	ch3 := make(chan int,1)

// 	ch1 <- 1
// 	ch2 <- 2

// 	select {
// 	case <- ch1:
// 		fmt.Println("ch1受信。",ch1)
// 	case <- ch2:
// 		fmt.Println("ch2受信。",ch2)
// 	case ch3 <- 3:
// 		fmt.Println("ch3送信",ch3)
// 	default:
// 		fmt.Println("到達しません")

// 	}
// }
