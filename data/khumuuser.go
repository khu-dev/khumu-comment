package data

// 댓글 DTO의 작성자로 제공될 타입
type SimpleKhumuUserDto struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Status   string `json:"status"`
}

// command-center 에게 이벤트로서 전달받는 KhumuUser 타입
type CommandCenterKhumuUserDto struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Status   string `json:"status"`
}
