package producerconsumer

import (
	"os"
	"os/signal"
	"syscall"

	"go0base/goprjs/05producer-consumer/out"
)

func Test_produ_comsum01() {
	o := out.NewOut()
	go o.OutPut()

	out.Println("ABCDEFG")
	out.Println("ABCDEFG")
	out.Println("ABCDEFG")
	out.Println("ABCDEFG")
	out.Println("ABCDEFG")
	out.Println("ABCDEFG")

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
