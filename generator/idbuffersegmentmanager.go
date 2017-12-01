package generator

import "seeder/monitor"

type NewIDBufferSegmentManager struct{

}

func NewIDBufferSegmentManager(bizTag string){
	segment := NewIDBufferSegment(bizTag)
	segment.CreateMasterIDBuffer(bizTag)
	go func(){
		for{
			monitor := monitor.NewMonitor(segment)
			monitor.SetVigilantValue(200)
			vigilant := monitor.IsOutVigilantValue()
			if vigilant {
				segment.CreateSlaveIDBuffer(bizTag)
			}
		}
	}()
	segment.GetId()
}
