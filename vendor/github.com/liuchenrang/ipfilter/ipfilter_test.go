package ipfilter_test

import (
	"fmt"
	"ipfilter"
	"testing"
)

func TestGetPrivateIP(t *testing.T) {
	for _, v := range ipfilter.GetPrivateIP(false) {
		fmt.Println(v)
	}
}

func TestGetPublicIP(t *testing.T) {
	for _, v := range ipfilter.GetPublicIP() {
		fmt.Println(v)
	}
}
