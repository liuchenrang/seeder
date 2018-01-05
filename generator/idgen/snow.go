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
// timestamtp(41) + idc(4) + node(6) + step(12) = 63
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

	WRITE_SNOW_TIME_INTERVAL = 10
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


func checkConfig(idc int64, node int64) error{
	if node < 0 || node > nodeMax {
		return  errors.New("Node number must be between 0 and " + fmt.Sprintf("%d",nodeMax))
	}
	if idc < 0 || idc > idcMax {
		return  errors.New("idc number must be between 0 and " +  fmt.Sprintf("%d",idcMax))
	}
	return nil
}
func NewNodeWithTime(application  *bootstrap.Application, idc int64, node int64, lastTime int64, step int64) (*Node, error) {
	if error := checkConfig(idc, node); error !=nil {
		return nil, error
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
	ticker :=  time.NewTicker(time.Second * WRITE_SNOW_TIME_INTERVAL)
	soa := n.application.GetServerSoa().(*zk.ServerSoa)
	fmt.Println("start at", n.time)
	soa.UpdateSnowTime(n.getNowTime() + WRITE_SNOW_TIME_INTERVAL * 1000) //初期启动时, 时间跳读到下次刷入时间

	go func() {
		for _ = range ticker.C {
			now := n.getNowTime()
			if now >= n.time {
				soa.UpdateSnowTime(n.getNowTime() + WRITE_SNOW_TIME_INTERVAL * 1000)
			}else{
				n.application.GetLogger().Error("tick time back to old value, now=%d, time=%d",now,n.time)
			}
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
	n.application.GetLogger().Debug("snow time  value, now=%d, time=%d",now,n.time)

	// 服务器时间回拨 后者 是 异常终端
	if now < n.time {
		for now < n.time {
			//fmt.Printf("wait snow time  value, now=%d, time=%d",now,n.time)
			//differ := time.Duration(n.time-now)
			//time.Sleep(time.Millisecond  * differ)
			now = n.getNowTime()
		}
	}

	if n.time == now {
		n.step = (n.step + 1) & stepMask
		if n.step == 0 {
			for n.time > now {
				now =  n.getNowTime()
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
