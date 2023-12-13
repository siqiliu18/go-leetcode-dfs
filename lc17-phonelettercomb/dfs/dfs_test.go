package dfs

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLetterCombinations(t *testing.T) {
	Convey("1", t, func() {
		res := LetterCombinations("234")
		exp := []string{"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi", "bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi", "cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi"}
		So(res, ShouldResemble, exp)
	})
}

func TestPermute(t *testing.T) {
	Convey("1", t, func() {
		input := []int{1, 2, 3}
		res := Permute(input)
		exp := [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}
		So(res, ShouldResemble, exp)
	})
}