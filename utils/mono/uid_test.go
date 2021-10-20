package mono

import (
	"fmt"
	"testing"
)

func TestUid(t *testing.T) {
	//uuid := make(map[string]string)
	mono := make(map[string]string)
	Init()

	for i := 0; i <= 100000000000000; i++ {
		gen3 := GetMONOID()
		if find, ok := mono[gen3]; ok {
			t.Error(fmt.Sprintf("MONO find: %v gen: %v", find, gen3))
			t.Fail()
			return
		} else {
			mono[gen3] = gen3
		}
	}
}
