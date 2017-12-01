package generator


import(
	"../generator"
	"../monitor"
)


func main(){
	
	 tag := "tag"
	 segment := generator.NewIDBufferSegment(tag)
	 masterBuffer := segment.CreateMasterIDBuffer()

	 go func(){
		for{
			monitor := monitor.NewMonitor(segment)
			monitor.SetVigilantValue(200)
			vigilant := monitor.IsOutVigilantValue()
			if vigilant {
				segment.CreateSlaveIDBuffer(tag)
			}	
		}
	 }

	 id = segment.GetId()
	 if id {
		 return id
	 }else{
		 segment.ChangeSlaveToMaster()
		 return segment.GetId()
	 }
	 
}
