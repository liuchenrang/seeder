package stats


type Stats struct{
	stats Stats
	total  int
}
func (stats *Stats) GetTotal() int{
	return stats.total
}
func (stats *Stats) Dig(){
	stats.total++
}
func (stats *Stats) Clear(){
	stats.total = 0
}

