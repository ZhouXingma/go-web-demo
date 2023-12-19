package idgenerater

import (
	"github.com/oklog/ulid"
	"math/rand"
	"time"
)

// ulid生成器
// 返回：
// string: ulid
func CreateUlid() string {
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return id.String()
}
