package rate

import (
	"sync"
)

var isReset = true // 默认为true, 为了首次重置剩余样本
var remainSample = make(map[string]int)
var memoryMutex sync.Mutex

type MemoryRateHandle struct {
}

// 重置剩余样本
func (this *MemoryRateHandle) Reset(sample map[string]int) {
	for key, value := range sample {
		remainSample[key] = value
	}
	isReset = false
}

// 判断是否重置样本空间
func (this *MemoryRateHandle) IsReset() bool {
	return isReset
}

// 命中后处理
func (this *MemoryRateHandle) HitAfter(key string) {
	remainSample[key]--
}

// 获取剩余样本
func (this *MemoryRateHandle) GetRemainSample() map[string]int {
	return remainSample
}

func (this *MemoryRateHandle) Lock() {
	memoryMutex.Lock()
}

func (this *MemoryRateHandle) Unlock() {
	memoryMutex.Unlock()
}
