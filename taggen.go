package taggen

import (
	"fmt"
	"math/rand"
	"time"
)

var chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

type Strategy int

const (
	None Strategy = iota
	CalVer
	SemVer
	FullCommitHash
	ShortCommit
	UniqTag
)

func (d Strategy) String() string {
	return [...]string{"None", "CalVer", "SemVer", "FullCommitHash", "ShortCommit", "UniqTag"}[d]
}

type Tag struct {
	CommitHash  string
	Strategy    Strategy // Tag generation strategy
	Sequence    int      // Build sequence
	GitTag      string
	IncludeArch bool
	Arch        string
}

func (T *Tag) Generate() string {
	now := time.Now()
	tag := ""
	switch T.Strategy {
	case None:
		tag = ""
	case CalVer:
		tag = fmt.Sprintf("%d.%d.%d-%s", now.Year()%100, now.Month(), T.Sequence, T.CommitHash[:8])
	case SemVer:
		tag = T.GitTag
	case FullCommitHash:
		tag = T.CommitHash
	case ShortCommit:
		tag = T.CommitHash[:8]
	case UniqTag:
		tag = shortID(16)

	}

	if T.IncludeArch {
		tag = tag + "-" + T.Arch
	}

	return tag
}

func shortID(length int) string {
	ll := len(chars)
	b := make([]byte, length)
	rand.Read(b) // generates len(b) random bytes
	for i := 0; i < length; i++ {
		b[i] = chars[int(b[i])%ll]
	}
	return string(b)
}
