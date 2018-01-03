package idgen

import (
	"errors"
	"strconv"
	"sync"
	"time"
	"fmt"
	"seeder/bootstrap"
	"seeder/zk"
)

// snowFlake算法:
// timestamtp(41) + idc(3) + node(7) + step(12) = 63
const (
	idcBits         = 3
	nodeBits        = 7
	stepBits        = 12
	idcMax          = -1 ^ (-1 << idcBits)
	nodeMax         = -1 ^ (-1 << nodeBits)
	stepMask  int64 = -1 ^ (-1 << stepBits)
	nodeShift       = stepBits
	idcShift        = nodeBits + stepBits
	timeShift       = idcBits + idcShift
)

var epoch int64 = 1513856639

type ID int64

type Node struct {
	sync.Mutex
	idc  int64 "idc"
	time int64 "last used time"
	node int64 "server node"
	step int64 "last plus step"
	application *bootstrap.Application
}

func NewNode(idc int64, node int64) (*Node, error) {
	if node < 0 || node > nodeMax {
		return nil, errors.New("Node number must be between 0 and 1023")
	}

	return &Node{
		idc:  idc,
		node: node,
		time: 0,
		step: 0,
	}, nil
}
func NewNodeWithTime(application  *bootstrap.Application, idc int64, node int64, lastTime int64, step int64) (*Node, error) {
	if node < 0 || node > nodeMax {
		return nil, errors.New("Node number must be between 0 and 1023")
	}

	return &Node{
		idc:  idc,
		node: node,
		application: application,
		time: lastTime,
		step: step,
	}, nil
}
func (n *Node) StartReport()  {
	ticker :=  time.NewTicker(time.Second * 2)
	fmt.Println("start at", n.time)

	go func() {
		for _ = range ticker.C {
			soa := n.application.GetServerSoa().(*zk.ServerSoa)
			soa.UpdateSnowTime(n.getNowTime())
			fmt.Println("Tick at", n.time)
		}
	}()
}
func (n *Node) getNowTime() int64 {
	return  time.Now().UnixNano() / 1000000
}
func (n *Node) Generate() ID {
	n.Lock()
	defer n.Unlock()

	now := n.getNowTime()

	// 服务器时间回拨
	if now < n.time {
		for now >= n.time {
			now = n.getNowTime()
			n.application.GetLogger().Warn("snow time back to old value, now=%d, time=%d",now,n.time)
		}
	}

	if n.time == now {
		n.step = (n.step + 1) & stepMask
		if n.step == 0 {
			for n.time > now {
				n.time =  n.getNowTime()
			}
		}
	} else {
		n.step = 0
	}

	n.time = now

	return ID((now-epoch)<<timeShift |
		(n.idc << idcShift) |
		(n.node << nodeShift) |
		(n.step))
}

func (f ID) Int64() int64 {
	return int64(f)
}
func (f ID) UInt64() uint64 {
	return uint64(f)
}

func (f ID) String() string {
	return strconv.FormatInt(int64(f), 10)
}
