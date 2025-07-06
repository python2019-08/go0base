package out

import "fmt"

// io操作是串行的，不是并发的。打印输出是

type Out struct {
	data chan interface{}
}

var out *Out

func NewOut() *Out {
	if out == nil {
		out = &Out{
			// data: make(chan interface{}),// 缓冲区大小为0,
			data: make(chan interface{}, 65565), // 缓冲区大小为65535,
		}
	}

	return out
}

func Println(i interface{}) {
	out.data <- i
}

func (o *Out) OutPut() {
	for i := range o.data {
		fmt.Println(i)
		fmt.Println("out put")
	}
	fmt.Println("out put...END")
	/*
		for {
			select {
			case i := <-o.data:
				fmt.Println(i)
			}

			fmt.Println("out put")
		}
	*/
}
