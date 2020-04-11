package pool

import (
	"log"

	wechatclient "github.com/rdxsl/go-wechat-work/client"
)

type Worker struct {
	ID            int
	WorkerChannel chan chan wechatclient.WechatMsg // used to communicate between dispatcher and workers
	Channel       chan wechatclient.WechatMsg
	End           chan bool
}

// start worker
func (w *Worker) Start() {
	go func() {
		for {
			w.WorkerChannel <- w.Channel // when the worker is available place channel in queue
			select {
			case wechatText := <-w.Channel: // worker has received job
				wechatclient.SendText(wechatText)
			case <-w.End:
				return
			}
		}
	}()
}

// end worker
func (w *Worker) Stop() {
	log.Printf("worker [%d] is stopping", w.ID)
	w.End <- true
}
