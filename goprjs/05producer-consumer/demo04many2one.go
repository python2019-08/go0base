package producerconsumer

import (
	"os"
	"os/signal"
	"syscall"

	"go0base/goprjs/05producer-consumer/many2one"
	"go0base/goprjs/05producer-consumer/out"
)

func Test_many2one() {
	o := out.NewOut()
	go o.OutPut()

	many2one.Exec()

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
