package model

// FollowWriteKafkaMsg 对应kafka接收的数据对象
type FollowWriteKafkaMsg struct {
	// Timestamp 时间戳
	Timestamp int64 `json:"timestamp"`
	// FromVuid 关注人vuid
	FromVuid int64 `json:"from_vuid"`
	// FromOmgid 关注人omgid
	FromOmgid string `json:"from_omgid"`
	// ToVuid 被关注人vuid
	ToVuid int64 `json:"to_vuid"`
	// IsFake 是否fake关注，0否，1是
	IsFake int32 `json:"is_fake"`
	// FollowAction 关注动作，0取消，1关注
	FollowAction int32 `json:"follow_action"`
}
