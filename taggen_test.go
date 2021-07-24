package taggen

import (
	"fmt"
	"testing"
)

func TestTag_Generate(t *testing.T) {
	tg := Tag{
		CommitHash: "305fcbefbffd3e565f518bc435ea574774db5d88",
		GitTag:     "2.2.5",
		Sequence:   8,
		Strategy:   None,
		Arch:       "arm64",
	}
	fmt.Println(tg.Strategy.String(), tg.Generate())

	tg.Strategy = CalVer
	fmt.Println(tg.Strategy.String(), tg.Generate())

	tg.Strategy = SemVer
	fmt.Println(tg.Strategy.String(), tg.Generate())

	tg.Strategy = FullCommitHash
	fmt.Println(tg.Strategy.String(), tg.Generate())

	tg.Strategy = ShortCommit
	fmt.Println(tg.Strategy.String(), tg.Generate())

	tg.Strategy = UniqTag
	fmt.Println(tg.Strategy.String(), tg.Generate())

	tg.Strategy = CalVer
	tg.IncludeArch = true
	fmt.Println(tg.Strategy.String(), tg.Generate())
}
