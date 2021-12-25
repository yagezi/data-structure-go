/*
*@Time        : 2021/12/16
*@Creator     :
*@File        : kmp.go
*@Description : 子串匹配 kmp 算法
 */

package ds_string


type KmpPattern struct {
	ptn string
	nxt []int
}

func NewKmpPattern(pattern string) *KmpPattern {
	p := &KmpPattern{
		ptn: pattern,
		nxt: make([]int, len(pattern)),
	}
	p.setNext()
	return p
}

func (p *KmpPattern) setNext() {
	p.nxt[0] = 0
	now := 0
	for i := 1; i < len(p.nxt); {
		if p.ptn[now] == p.ptn[i] {
			now++
			p.nxt[i] = now
			i++
		} else if now > 0 {
			now = p.nxt[now-1]
		} else {
			p.nxt[i] = 0
			i++
		}
	}
}

func (p KmpPattern) Match(str string) int {
	// i 目标串索引，j 模式串索引
	var i, j int
	for i < len(str) {
		if p.ptn[j] == str[i] {
			i++
			j++
		} else if j > 0 {
			j = p.nxt[j-1]
		} else {
			i++
		}

		if j == len(p.ptn) {
			return i - j
		}
	}
	return -1
}

func (p KmpPattern) MatchAll(str string) (res []int) {
	// i 目标串索引，j 模式串索引
	var i, j int
	for i < len(str) {
		if p.ptn[j] == str[i] {
			i++
			j++
		} else if j > 0 {
			j = p.nxt[j-1]
		} else {
			i++
		}

		if j == len(p.ptn) {
			res = append(res, i-j)
			j = 0
		}
	}
	return
}
