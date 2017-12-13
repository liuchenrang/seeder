package generator_test

import (
	//"errors"
	//"fmt"
	"fmt"
	"github.com/liuchenrang/log4go"
	"seeder/bootstrap"
	"seeder/config"
	"seeder/generator"
	"seeder/generator/idgen"
	"seeder/logger"
	"testing"
	"sync"
	"sort"
)

func TestGenIDGo(t *testing.T) {

	Application := bootstrap.NewApplication()
	seederConfig := config.NewSeederConfig("../seeder.yaml")
	Application.Set("globalSeederConfig", seederConfig)
	Application.Set("globalLogger", SeederLogger.NewLogger4g(log4go.DEBUG, seederConfig))
	segment := generator.NewIDBufferSegment("test", Application)
	wg := sync.WaitGroup{}
	var ids []int
	for j:=0; j < 5; j++ {
		wg.Add(1)
		go func() {
			var i uint64
			for i < 5 {
				id := segment.GetId()
				ids = append(ids, int(id))
				i++
			}
			wg.Done()
		}()
	}
	wg.Wait()
	sort.Ints(ids)
	len := len(ids)
	for k:=0; k< len; k++ {
		if k + 1 < len {
			if ids[k]+1 != ids[k+1] {

				t.Errorf("id=%d, next=%d",ids[k], ids[k+1])
				panic("exit")
			}
		}
		 fmt.Printf("%d\n",ids[k])
	}
}
func TestGenID(t *testing.T) {


	// Different allocations should not be equal.
	Application := bootstrap.NewApplication()
	seederConfig := config.NewSeederConfig("../seeder.yaml")
	Application.Set("globalSeederConfig", seederConfig)
	Application.Set("globalLogger", SeederLogger.NewLogger4g(log4go.DEBUG, seederConfig))
	segment := generator.NewIDBufferSegment("test", Application)

	var id uint64
	logger := SeederLogger.NewLogger(seederConfig)
	var i uint64
	for i < 1000 {
		id = segment.GetId()
		nextId := segment.GetId()
		logger.Debug("id ", id, "nextId", nextId)
		if id+1 != nextId {
			t.Error("id error")
			break
		}
		i++
	}

}

func BenchmarkIDBufferSegment_GetId(b *testing.B) {
	Application := bootstrap.NewApplication()
	seederConfig := config.NewSeederConfig("../seeder.yaml")
	Application.Set("globalSeederConfig", seederConfig)
	Application.Set("globalLogger", SeederLogger.NewLogger4g(log4go.ERROR, seederConfig))
	segment := generator.NewIDBufferSegment("uts", Application)

	fmt.Printf("segment=%p\n", segment)
	var id uint64

	i := func(pb *testing.PB) {

		for pb.Next() {
			id = segment.GetId()
			//fmt.Println("id", id)

		}
	}
	b.RunParallel(i)
}

func TestDB(t *testing.T) {
	// Different allocations should not be equal.
	Application := bootstrap.NewApplication()
	seederConfig := config.NewSeederConfig("../seeder.yaml")
	Application.Set("globalSeederConfig", seederConfig)
	Application.Set("globalLogger", SeederLogger.NewLogger4g(log4go.DEBUG, seederConfig))

	bizTag := "test5"
	segment := generator.NewIDBufferSegment(bizTag, Application)
	segment.CreateMasterIDBuffer(bizTag)
	db := idgen.NewDBGen(bizTag, Application)

	fmt.Println(db.Find(bizTag))
}
