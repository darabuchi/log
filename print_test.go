package log_test

import (
	"github.com/darabuchi/log"
	"sync"
	"testing"
)

func TestPrint(t *testing.T) {
	// fs , _, err := zap.Open("./out.log")
	// if err != nil {
	//    log.Errorf("err:%v", err)
	//    return
	// }

	log.SetTrace("")
	log.Info("msg")
	log.Infof("msgf")
	log.Infof("%d", 1)

	log.Clone().Caller(true).Info("not caller")
	log.Info("has caller")

	var w sync.WaitGroup
	w.Add(1)
	go func() {
		defer w.Done()
		log.Info("go")
	}()

	w.Wait()
}
