package generator


import(
	"../generator/"
	"../monitor"
)




func main(){
	
	 segment := generator.NewIDBufferSegment('tag')
	 segment.CreateMasterIDBuffer()

	 monitor := monitor.NewMonitor()


}
