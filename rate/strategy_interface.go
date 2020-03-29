package rate

type StrategyIf interface {
	IsReset() bool               // 判断是否重置剩余样本
	Reset(sample map[string]int) // 重置剩余样本, 首次必须重置

	GetRemainSample() map[string]int // 获取剩余样本
	HitAfter(key string)                                  // 命中后处理

	Lock()   // 加锁
	Unlock() // 解锁
}
