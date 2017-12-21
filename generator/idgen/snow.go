package idgen

import (
	"errors"
	"strconv"
	"sync"
	"time"
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
	s    store
}

type store interface {
	get() int64
	set(int64)
}

func NewNode(idc int64, node int64, s store) (*Node, error) {
	if node < 0 || node > nodeMax {
		return nil, errors.New("Node number must be between 0 and 1023")
	}

	return &Node{
		s:    s,
		idc:  idc,
		node: node,
		time: 0,
		step: 0,
	}, nil
}

func (n *Node) generate() ID {
	n.Lock()
	defer n.Unlock()

	now := time.Now().UnixNano() / 1000000

	// 服务器时间回拨
	if now < n.time {
		for now >= n.time {
			now = time.Now().UnixNano() / 1000000
		}
	}

	if n.time == now {
		n.step = (n.step + 1) & stepMask
		if n.step == 0 {
			for n.time > now {
				n.time = time.Now().UnixNano() / 1000000
			}
		}
	} else {
		n.step = 0
	}

	n.time = now

	// 保存上次的使用时间,防止服务器重启后时间回拨
	go func() {
		n.s.set(n.time)
	}()

	return ID((now-epoch)<<timeShift |
		(n.idc << idcShift) |
		(n.node << nodeShift) |
		(n.step))
}

func (f ID) Int64() int64 {
	return int64(f)
}

func (f ID) String() string {
	return strconv.FormatInt(int64(f), 10)
}
