package rate

import (
	"fmt"
	"testing"
)

func TestRate(t *testing.T) {
	ps := make(map[string]int)

	testValue1 := 5
	testValue2 := 95
	testKey1 := "A"
	testKey2 := "B"

	ps[testKey1] = testValue1
	ps[testKey2] = testValue2

	nph := NewRateHandle(ps, &MemoryRateHandle{})

	mapKey := make(map[string]int)

	for i := 0; i < 10000; i++ {
		key, isEnd := nph.HappenSelect()

		if _, ok := mapKey[key]; !ok {
			mapKey[key] = 1
		} else {
			mapKey[key]++
		}

		if isEnd {
			if mapKey[testKey1] == testValue1 && mapKey[testKey2] == testValue2 {
			} else {
				println("test is error")
				fmt.Println(mapKey)
			}
			mapKey = map[string]int{}
		}
	}

	println("test success")
}
