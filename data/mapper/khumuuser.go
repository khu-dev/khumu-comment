package mapper

import (
	"github.com/khu-dev/khumu-comment/data"
	"github.com/khu-dev/khumu-comment/ent"
)

// 새로운 SimpleKhumuUserOutput을 생성한다
// 원래 KhumuUser 참조 X
func KhumuUserModelToSimpleOutput(src *ent.KhumuUser, dest *data.SimpleKhumuUserOutput) *data.SimpleKhumuUserOutput {
	if src == nil {
		return nil
	}

	if dest == nil {
		dest = &data.SimpleKhumuUserOutput{}
	}

	if src.Status == "deleted" {
		dest.Username = "탈퇴한 유저"
		dest.Nickname = "탈퇴한 유저"
		dest.Status = "deleted"
	}

	dest.Username = src.ID
	dest.Nickname = src.Nickname
	dest.Status = src.Status

	return dest
}
