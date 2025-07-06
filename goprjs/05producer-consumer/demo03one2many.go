package producerconsumer

import (
	"os"
	"os/signal"
	"syscall"

	"go0base/goprjs/05producer-consumer/one2many"
	"go0base/goprjs/05producer-consumer/out"
)

func Test_one2many() {
	o := out.NewOut()
	go o.OutPut()

	one2many.Exec()

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
