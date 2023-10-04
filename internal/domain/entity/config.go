package entity

type Config struct {
	// 채널을 통한 응답을 받습니다.
	Channel bool
	// 알림을 통한 응답을 받습니다.
	Notification bool
	// 자동 매매를 사용합니다.
	AutoTrading bool
}
