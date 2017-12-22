package idgen

import (
	"fmt"
	"testing"
)

func TestNewNode(t *testing.T) {
	node, _ := NewNode(2, 4)
	id := node.Generate()
	step := id & (-1 ^ (-1 << 12))
	id1 := node.Generate()
	step1 := id1 & (-1 ^ (-1 << 12))

	id2 := node.Generate()
	step2 := id2 & (-1 ^ (-1 << 12))
	fmt.Println(step, step1, step2)
	serverID := (id & (-1 ^ (-1<<7)<<12)) >> 12
	idcID := (id & (-1 ^ (-1<<3)<<19)) >> 19
	fmt.Println(serverID, idcID)

	idstr := id1.String()
	fmt.Printf(idstr)
}

func BenchmarkGenerate(b *testing.B) {

	node, _ := NewNode(1, 3)

	b.ReportAllocs()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_ = node.Generate()
	}
}
