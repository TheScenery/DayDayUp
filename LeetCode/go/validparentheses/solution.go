package validparentheses

import "container/list"

func isValid(s string) bool {
	l := list.New()
	parenthessesMap := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	for _, ch := range s {
		back := l.Back()
		if back == nil || ch != parenthessesMap[back.Value.(rune)] {
			l.PushBack(ch)
		} else {
			l.Remove(l.Back())
		}
	}
	return l.Len() == 0
}
