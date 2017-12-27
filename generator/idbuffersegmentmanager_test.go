package generator

import (
	"fmt"
	"seeder/bootstrap"
	"seeder/config"
	"seeder/logger"
	"testing"
	"github.com/liuchenrang/log4go"
	"os"
	"bufio"
	"bytes"
	"log"
	"time"
	"math/rand"
)

func TestSegManager(t *testing.T) {
	// Different allocations should not be equal.

	Application := bootstrap.NewApplication()
	seederConfig := config.NewSeederConfig("../seeder.yaml")
	Application.Set("globalSeederConfig", seederConfig)
	Application.Set("globalLogger", SeederLogger.NewLogger4g(log4go.DEBUG, seederConfig))
	m := NewIDBufferSegmentManager(Application)

	i := 0
	runTime := 30
	for i <= runTime {
		i++
		id, _ := m.GetId("uts", 1)
		fmt.Println("id====== ", id)
	}

}

func BenchmarkLoopsM(b *testing.B) {
	Application := bootstrap.NewApplication()
	seederConfig := config.NewSeederConfig("../seeder.yaml")
	Application.Set("globalSeederConfig", seederConfig)
	Application.Set("globalLogger", SeederLogger.NewLogger4g(log4go.CRITICAL, seederConfig))
	m := NewIDBufferSegmentManager(Application)
	i := func(pb *testing.PB) {

		for pb.Next() {

			id, _ := m.GetId("uts", 1)
			fmt.Println("id====== ", id)

		}
	}
	b.RunParallel(i)
}

func TestFmt(t *testing.T) {
	// Different allocations should not be equal.
	fmt.Println("xx")
}
func TestIDBufferSegmentManager_StartHotPreLoad(b *testing.T) {
	Application := bootstrap.NewApplication()
	seederConfig := config.NewSeederConfig("../seeder.yaml")
	Application.Set("globalSeederConfig", seederConfig)
	Application.Set("globalLogger", SeederLogger.NewLogger4g(log4go.CRITICAL, seederConfig))
	m := NewIDBufferSegmentManager(Application)
	m.StartHotPreLoad()
}
func BenchmarkLoopsMultiTag(b *testing.B) {
	Application := bootstrap.NewApplication()
	seederConfig := config.NewSeederConfig("../seeder.yaml")
	Application.Set("globalSeederConfig", seederConfig)
	path := "../logo4go.xml"
	Application.Set("globalLogger", SeederLogger.NewLogger4gWithConfig(log4go.DEBUG, seederConfig, &path))
	f,_ := os.Open("../client/tags.csv")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var tags []string
	for scanner.Scan() {
		tags = append(tags, string(bytes.TrimSpace([]byte(scanner.Text()))))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	len := len(tags)
	rand.Seed(time.Now().Unix())
	m := NewIDBufferSegmentManager(Application)

	i := func(pb *testing.PB) {

		for pb.Next() {
			tag := tags[rand.Intn(len)]
			id, _ := m.GetId(tag, 1)
			fmt.Printf("GetTag %s, GetId=%d ",tag, id)

		}
	}
	b.SetParallelism(200)
	b.RunParallel(i)
}