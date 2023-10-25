package compose

type Game interface {
	Start()
	//暂停
	Pause()
	//继续
	Resume()
	//停止
	Stop()
}
