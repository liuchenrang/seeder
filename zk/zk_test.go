package zk_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

func TestRegister(t *testing.T) {

	c, _, err := zk.Connect([]string{"127.0.0.1"}, time.Second) //*10)
	if err != nil {
		panic(err)
	}
	if succ, _, _ := c.Exists("/seeder"); !succ {
		c.Create("/seeder", nil, 0, zk.WorldACL(zk.PermAll))
	}
	if succ, _, _ := c.Exists("/seeder/workers"); !succ {
		c.Create("/seeder/workers", nil, 0, zk.WorldACL(zk.PermAll))
	}
	go func() {
		c.Create("/seeder/workers/0", nil, zk.FlagEphemeral+zk.FlagSequence, zk.WorldACL(zk.PermAll))
	}()
	go func() {
		c.Create("/seeder/workers/0", nil, zk.FlagEphemeral+zk.FlagSequence, zk.WorldACL(zk.PermAll))
	}()
	go func() {
		res, _ := c.Create("/seeder/workers/0", nil, zk.FlagEphemeral+zk.FlagSequence, zk.WorldACL(zk.PermAll))
		fmt.Println("%+v", res)
	}()
	go func() {
	ccc:
		exists, stats, change, error := c.ExistsW("/seeder")
		if error != nil {
			fmt.Println(error)
		}
		fmt.Printf("node exist %b", exists)
		fmt.Println("status %+v", stats)
		event := <-change

		fmt.Println("event %+v", event)
		goto ccc
	}()
	// children, stat, ch, err := c.ChildrenW("/seeder/workers")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%+v %+v\n", children, stat)
	// e := <-ch
	// fmt.Printf("%+v\n", e)
	time.Sleep(time.Second * 35)
}
