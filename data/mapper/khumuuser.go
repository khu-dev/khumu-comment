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

	dest.Username = src.ID
	dest.Nickname = src.Nickname
	dest.State = "enabled"

	return dest
}
