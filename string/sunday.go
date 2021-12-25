/*
*@Time        : 2021/12/17
*@Creator     :
*@File        : sunday.go
*@Description : 子串匹配sunday算法
 */

package ds_string

type SundayPattern struct {
	ptn    string
	length int
	idx    [256]int
}

func NewSundayPattern(pattern string) *SundayPattern {
	sp := &SundayPattern{
		ptn:    pattern,
		length: len(pattern),
	}
	sp.setIndex()
	return sp
}

func (sp *SundayPattern) setIndex() {
	for i := 0; i < len(sp.idx); i++ {
		sp.idx[i] = -1
	}
	for j := 0; j < len(sp.ptn); j++ {
		sp.idx[sp.ptn[j]] = j
	}
}

func (sp SundayPattern) Match(str string) int {
	var i, j int
	for i < len(str) {
		if str[i] == sp.ptn[j] {
			i++
			j++
		} else if i-j+sp.length < len(str) {
			// 好字符(good char)不在模式串中时，x - l = x + 1,
			l := sp.idx[str[i-j+sp.length]]
			i = i - j + sp.length - l
			j = 0
		} else {
			break
		}

		if j == len(sp.ptn) {
			return i - j
		}
	}
	return -1
}

func (sp SundayPattern) MatchAll(str string) (res []int) {
	var i, j int
	for i < len(str) {
		if str[i] == sp.ptn[j] {
			i++
			j++
		} else if i-j+sp.length < len(str) {
			// 好字符不在模式串中时，x - l = x + 1,
			l := sp.idx[str[i-j+sp.length]]
			i = i - j + sp.length - l
			j = 0
		} else {
			// 主串的后缀失配
			break
		}

		if j == len(sp.ptn) {
			res = append(res, i-j)
			j = 0
		}
	}
	return
}
