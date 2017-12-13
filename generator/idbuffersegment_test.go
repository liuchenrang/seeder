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
)

func TestGenID(t *testing.T) {


	// Different allocations should not be equal.
	Application := bootstrap.NewApplication()
	seederConfig := config.NewSeederConfig("../seeder.yaml")
	Application.Set("globalSeederConfig", seederConfig)
	Application.Set("globalLogger", SeederLogger.NewLogger4g(log4go.DEBUG, seederConfig))
	segment := generator.NewIDBufferSegment("uts", Application)

	var id uint64
	logger := SeederLogger.NewLogger(seederConfig)
	var i uint64
	for i < 300 {

		id = segment.GetId()
		nextId := segment.GetId()
		logger.Debug("id ", id, "nextId", nextId)
		fmt.Printf("xxxx")
		if id+1 != nextId {
			t.Error("id error")
			break
		}
		fmt.Println()
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
