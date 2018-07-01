package pipeline

import (
	"sort"
	"io"
	"encoding/binary"
	"math/rand"
)

// 将Array中的数据都存放到channel中,箭头告诉使用此函数只能读取到chan值，只出不进,使用他的人只能从这个管道里拿东西
func ArraySource(arr ...int)  <-  chan int {
	out:=make(chan int)
	// 因为我们的channel是在goroutine之间传输的通道，所以我们不能自己将自己的arr输送到自己的channel
	go func() {
		for _,v := range arr{
			out <- v
		}
		// 数据传输完毕，并行计算pipeline中最好是使用close结束表明数据传输已经完成了
		close(out)
	}()
	return out
}

func IntMemorySort(in <- chan int) <- chan int  {
	out := make(chan int)

	go func() {
		// 将数据加载到内存
		arr := []int{}
		for v:= range in  {
			arr = append(arr,v)
		}

		// 快速排序
		sort.Ints(arr)

		// 排序后加载到管道中,即排序好前下游接收会处于等待状态。
		for _,v := range arr {
			out <- v
		}
		close(out)
	}()


	return out
}

// 归并排序
func Merge(in1, in2 <- chan int) <- chan int {
	out := make(chan int)

	go func() {
		v1, ok1 := <- in1
		v2, ok2 := <- in2
		for ok1 || ok2{
			if !ok2 ||(ok1 && v1<=v2){
				out <- v1
				v1,ok1 = <- in1
			}else {
				out <- v2
				v2, ok2 = <- in2
			}
		}
		close(out)
	}()


	return out
}

// 读取文件到管道
func ReaderSource(reader io.Reader) <- chan int  {
	out := make(chan int)

	go func() {
		buffer := make([]byte, 8)
		for {
			n, err := reader.Read(buffer)
			if n>0 {
				v := int(binary.BigEndian.Uint64(buffer))
				out <- v
			}
			if err != nil {
				break
			}
		}
		close(out)
	}()
	return out
}

// 将数据写到Sink
func WriterSink(writer io.Writer, in <- chan int)  {
	for v := range in{
		buffer := make([]byte, 8)
		binary.BigEndian.PutUint64(buffer,uint64(v))
		writer.Write(buffer)
	}
}

func RandomSource(count int) <- chan int {
	out := make(chan int)
	go func() {
		for i:=0;i<count;i++ {
			out <- rand.Int()
		}
	}()
	return out
}

