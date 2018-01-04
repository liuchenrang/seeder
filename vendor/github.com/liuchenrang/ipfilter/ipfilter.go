package ipfilter

import (
	"encoding/binary"
	"net"
	"sort"
	"strings"
)

//IP2long ipstr convert to uint32
func IP2long(ipstr string) uint32 {
	ip := net.ParseIP(ipstr)
	if ip == nil {
		return 0
	}
	ip = ip.To4()
	return binary.BigEndian.Uint32(ip)
}

//Long2IP convert uint32 to string
func Long2IP(ipLong uint32) string {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, ipLong)
	ip := net.IP(ipByte)
	return ip.String()
}

type BySize []uint32

func (s BySize) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s BySize) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s BySize) Len() int {
	return len(s)
}

//GetIPV4 return slice uint32
func GetIPV4() []uint32 {
	var iplist []uint32
	addres, _ := net.InterfaceAddrs()
	for add := range addres {
		ip := addres[add].String()
		posSlash := strings.Index(ip, "/")
		ip = ip[:posSlash]
		trial := net.ParseIP(ip)
		if trial.To4() == nil {
			continue
		}
		iplist = append(iplist, IP2long(ip))
	}
	sort.Sort(BySize(iplist))
	return iplist
}

//GetPrivateIP return private slice ip string
func GetPrivateIP(skipLocal bool) []string {
	var ips []string
	iplist := GetIPV4()
	for _, ipDex := range iplist {
		if IsPrivate(ipDex) {
			if skipLocal {
				if IsLocal(ipDex) {
					continue
				}
			}
			ips = append(ips, Long2IP(ipDex))
		}
	}
	return ips
}
func IsPrivate(ip_decimal uint32) bool {
	if IsLocal(ip_decimal) || //-- 127.0.0.0 ~ 127.255.255.255
		ip_decimal >= 0x0a000000 && ip_decimal <= 0x0affffff || //-- 10.0.0.0 ~ 10.255.255.255
		ip_decimal >= 0xac100000 && ip_decimal <= 0xac1fffff || //-- 172.16.0.0 ~ 172.31.255.255
		ip_decimal >= 0xc0a80000 && ip_decimal <= 0xc0a8ffff { //  -- 192.168.0.0 ~ 192.168.255.255
		return true
	} else {
		return false
	}
}

//IsLocal check is 127.x.x.x
func IsLocal(ipDex uint32) bool {
	if ipDex >= 0x7f000000 && ipDex <= 0x7fffffff {
		return true
	}
	return false
}

//GetPublicIP is internet ip
func GetPublicIP() []string {
	var ips []string
	iplist := GetIPV4()
	for _, ipDex := range iplist {
		if !IsPrivate(ipDex) {
			ips = append(ips, Long2IP(ipDex))
		}
	}
	return ips
}
