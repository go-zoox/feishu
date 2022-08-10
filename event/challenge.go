package event

type ChallengeRequest struct {
	// 未加密
	Challenge string `json:"challenge"`
	Tolen     string `json:"token"`
	Type      string `json:"type"`
	// 已加密
	Encrypt string `json:"encrypt"`
}

type ChallengeResponse struct {
	Challenge string `json:"challenge"`
}
