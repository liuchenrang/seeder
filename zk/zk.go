package zk

import (
	"github.com/samuel/go-zookeeper/zk"
	"seeder/bootstrap"
	"time"
	"strings"
	"strconv"
	"encoding/binary"
	error2 "seeder/error"
	"fmt"
)

type ServerSoa struct {
	host        string
	zkClient    *zk.Conn
	ZkPrefix string
	application *bootstrap.Application
}

func (soa *ServerSoa) Register() {

	_, error := soa.zkClient.Create( soa.ZkPrefix + "/seeder/servers/"+soa.host, nil, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	if error != nil && error.Error() != error2.ZK_NODE_EXITSTS {
		soa.application.GetLogger().Warn(error)
	}
}

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}
func (soa *ServerSoa) UpdateSnowTime(time int64) {

	snowTimePath := soa.ZkPrefix  +  "/seeder/data/" + soa.host + "/time"
	_, stat, err := soa.zkClient.Get(snowTimePath)
	if err != nil {
		soa.application.GetLogger().Error(err)
	}
	timeData := fmt.Sprintf("%d", time)
	stat, err = soa.zkClient.Set(snowTimePath, []byte(timeData), stat.Version)
	if err != nil {
		soa.application.GetLogger().Error(err)
	}
}
func (soa *ServerSoa) GetSnowTime() (int64) {

	snowTimePath := soa.ZkPrefix  + "/seeder/data/" + soa.host + "/time"
	data, _, err := soa.zkClient.Get(snowTimePath)
	if err != nil {
		soa.application.GetLogger().Error(err)
		panic(err)
	}
	time, _ := strconv.ParseInt(string(data), 10, 64)
	return time

}
func (soa *ServerSoa) GeneratorID() (int string) {
	path, error := soa.zkClient.Create(soa.ZkPrefix  + "/seeder/workers/", nil, zk.FlagEphemeral+zk.FlagSequence, zk.WorldACL(zk.PermAll))
	if error != nil {
		soa.application.GetLogger().Error(path)
		panic(error)
	}
	lastSlash := strings.LastIndex(path, "/")
	path = strings.TrimLeft(path[lastSlash+1:], "0")
	nodeId, _ := strconv.ParseInt(path, 10, 32)
	int = strconv.FormatInt(nodeId, 10)
	return
}
func (soa *ServerSoa) Initialize(serverAddr string) *ServerSoa {
	var (
		nodeId int64
	)
	if len(serverAddr) <= 0 {
		panic("serverAddr must set!")
	}
	soa.host = serverAddr

	c, _, err := zk.Connect(soa.application.GetConfig().Zookeeper, time.Second) //*10)
	if err != nil {
		soa.application.GetLogger().Error(err)
		panic(err)
	}

	soa.zkClient = c
	soa.AddNode(soa.ZkPrefix, nil)
	soa.AddNode(soa.ZkPrefix  + "/seeder", nil)
	soa.AddNode(soa.ZkPrefix  + "/seeder/servers", nil)
	soa.AddNode(soa.ZkPrefix  + "/seeder/workers", nil)
	soa.AddNode(soa.ZkPrefix  + "/seeder/data", nil)
	soa.AddNode(soa.ZkPrefix  + "/seeder/data/"+soa.host, nil)
	soa.AddNode(soa.ZkPrefix  + "/seeder/data/"+soa.host+"/time", nil)
	if exists, _, _ := soa.zkClient.Exists(soa.ZkPrefix  + "/seeder/data/" + soa.host + "/id"); !exists {
		int := soa.GeneratorID()
		soa.zkClient.Create(soa.ZkPrefix  + "/seeder/data/"+soa.host+"/id", []byte(int), 0, zk.WorldACL(zk.PermAll))
		nodeId, _ = strconv.ParseInt(int, 10, 32)
	} else {
		nodeData, _, _ := soa.zkClient.Get(soa.ZkPrefix  + "/seeder/data/" + soa.host + "/id")
		tid, _ := strconv.Atoi(string(nodeData[:]))
		nodeId = int64(tid)
	}
	configSeeder := soa.application.GetConfig()
	soa.application.GetLogger().Info("snow work id %d", nodeId)
	configSeeder.Snow.Node = nodeId

	return soa
}
func (soa *ServerSoa) AddNode(path string, data []byte) (result string, err error) {
	if exists, _, _ := soa.zkClient.Exists(path); !exists {
		result, err = soa.zkClient.Create(path, data, 0, zk.WorldACL(zk.PermAll))
	}
	return
}

func NewServerSoa(application *bootstrap.Application, host string) *ServerSoa {
	soa := &ServerSoa{
		application: application,
		ZkPrefix: application.GetConfig().Common.ZkPrefix,
	}
	return soa.Initialize(host)
}
