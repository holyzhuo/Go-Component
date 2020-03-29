package rate

import (
	"math/rand"
	"time"
	"sort"
)

const (
	SUM_SAMPLE = 100 // 默认样本总数100
)

type RateHandle struct {
	Sample    map[string]int
	EnterRate StrategyIf
}

func NewRateHandle(sample map[string]int, enterRate StrategyIf) *RateHandle {
	if !checkSample(sample) {
		panic("SUM_SAMPLE IS NOT right")
	}

	ph := &RateHandle{}
	ph.Sample = sample
	ph.EnterRate = enterRate

	// 生成随机数种子，保证随机数随机
	rand.Seed(int64(time.Now().UnixNano()))

	return ph
}

// 检查样本总数
func checkSample(sample map[string]int) bool {
	sum := 0
	for _, v := range sample {
		sum += v
	}

	if sum != SUM_SAMPLE {
		return false
	}
	return true
}

// 随机选择流程
func (this *RateHandle) HappenSelect() (string, bool) {
	this.EnterRate.Lock()
	defer this.EnterRate.Unlock()

	if this.EnterRate.IsReset() {
		this.EnterRate.Reset(this.Sample)
	}

	var isRoundEnd bool
	var hitKey string
	var randNum, nextNum int

	remainCount, remainSample := this.GetRemain()
	if remainCount != 1 {
		randNum = rand.Intn(remainCount - 1)
	}

	for _, v := range this.GetOrderedKey() {
		nextNum += remainSample[v]
		if randNum < nextNum {
			hitKey = v
			this.EnterRate.HitAfter(v)
			break
		}
	}

	if remainCount == 1 {
		isRoundEnd = true
		this.EnterRate.Reset(this.Sample)
	}

	return hitKey, isRoundEnd
}

// 获取剩余样本数据
func (this *RateHandle) GetRemain() (total int, remain map[string]int) {
	remain = this.EnterRate.GetRemainSample()
	for _, value := range remain {
		total += value
	}
	return
}

// 排序
func (this *RateHandle) GetOrderedKey() (sort.StringSlice) {
	keySlice := sort.StringSlice{}
	for key, _ := range this.Sample {
		keySlice = append(keySlice, key)
	}

	sort.Sort(keySlice)
	return keySlice
}
