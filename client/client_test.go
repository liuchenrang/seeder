package client

import (
	"testing"
	thriftGenerator "seeder/thrift/packages/generator"
	"fmt"
	"log"
)

func TestNewClient(t *testing.T) {

	client := NewClient()

	idp, _  := client.Ping()
	i := 0
	for i < 100 {
		id, error := client.GetId(&thriftGenerator.TGetIdParams{Tag: "uts", GeneratorType: 1})
		if error != nil {
			log.Fatal(error)
		}
		fmt.Println("ping ", idp)
		fmt.Println("id", id)
		i++
	}

}
