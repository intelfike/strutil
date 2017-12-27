package strutil

import "math/rand"

// 完全ランダムに文字列をシャッフル
func Shuffle(s string) string {
	rs := []rune(s)
	for n, _ := range rs {
		ran := int(rand.Int31()) % len(rs)
		rs[n], rs[ran] = rs[ran], rs[n]
	}
	return string(rs)
}

// 複数の配列でAnd
func And(ss ...[]string) []string {
	and := make([]string, len(ss[0]))
	copy(and, ss[0])
	if len(ss) != 1 {
		for _, v := range ss[1:] {
			and = and(and, v)
		}
	}
	return and
}

// 2つのスライスでAnd
func and(a, b []string) []string {
	c := make([]string, 0)
	for _, v := range a {
		index := -1
		for n2, v2 := range b {
			if v == v2 {
				index = n2
			}
		}
		if index == -1 {
			continue
		}
		c = append(c, b[index])
	}
	return c
}

func Unique(ss ...string) []string {
	result := make([]string, 0)
	unique := map[string]bool{}
	for _, v := range ss {
		_, ok := unique[v]
		if ok {
			continue
		}
		unique[v] = true
		result = append(result, v)
	}
	return result
}
