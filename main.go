package main

import (
	"time"

	pool "github.com/rdxsl/go-wechat-work/workerpool"
)

const WORKER_COUNT = 5

func main() {
	// wechatclient.Init()
	// err := wechatclient.GetAccessTocken()

	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	// 	os.Exit(1)
	// }

	// wechatclient.Test1()
	//workerpool.Test()
	collector := pool.StartDispatcher(WORKER_COUNT)
	collector.Work <- pool.Work{Job: "test", ID: 1}
	collector.Work <- pool.Work{Job: "test", ID: 1}
	collector.Work <- pool.Work{Job: "test", ID: 1}
	collector.Work <- pool.Work{Job: "test", ID: 1}
	collector.Work <- pool.Work{Job: "test", ID: 1}
	collector.Work <- pool.Work{Job: "test", ID: 1}
	time.Sleep(5 * time.Second)
}
