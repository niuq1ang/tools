package utils

import (
	"strings"

	pinyin "github.com/mozillazg/go-pinyin"
)

var (
	PinyinArgs = pinyin.NewArgs()
)

func init() {
	PinyinArgs.Style = pinyin.Tone3
	PinyinArgs.Fallback = func(r rune, args pinyin.Args) []string {
		return []string{string(r)}
	}
}

func ToPinyin(input string) string {
	data := pinyin.Pinyin(input, PinyinArgs)
	output := ""
	for _, word := range data {
		output += strings.Join(word, "")
	}
	return output
}
