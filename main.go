package main

import "fmt"

const (
	SEEDER_CONFIG_PATH = "/usr/local/Cellar/go/gopath/src/seeder/seeder.yaml"
)
func main()  {
	fmt.Println(SEEDER_CONFIG_PATH)
	//select {
	//	case ct := <-inchan:
	//			fmt.Printf("1000000  ", ct)
	//	default:
	//			fmt.Printf("2000000 xx")
	//}
	//var input string
	//fmt.Scanln(&input)
	//fmt.Printf("hh %s",input)
}
func GetSeederConfigPath() string {
	return SEEDER_CONFIG_PATH
}