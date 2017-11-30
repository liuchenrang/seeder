package monitor

import (
	. "seeder/stats"
)

type Monitor struct{
	stats *Stats
}
func (m *Monitor) SetVigilantValue(){

}
func (m *Monitor) IsOutVigilantValue(){

}
func (m *Monitor) Event(tag <-chan string){

}
func (m *Monitor) GetStats() *Stats {
	return m.stats
}


func NewMonitor() *Monitor {
	stats := &Stats{}
	return  &Monitor{stats: stats}
}