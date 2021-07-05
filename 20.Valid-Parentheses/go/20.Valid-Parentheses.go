package leetcode

import "container/list"

// 时间复杂度: O(n)
// 空间复杂度: O(n)
func isValid(s string) bool {
	l := list.New()
	for _, b := range s {
		switch b {
		case 32:
		case 40, 91, 123:
			l.PushBack(b)
		case 41, 93, 125:
			value := pop(l)
			if value == 0 {
				return false
			}
			if value != b-1 && value != b-2 {
				return false
			}
		default:
			return false
		}
	}
	if l.Len() != 0 {
		return false
	}
	return true
}

func pop(l *list.List) int32 {
	el := l.Back()
	for el != nil {
		cn, ok := el.Value.(int32)
		if !ok {
			return 0
		}
		l.Remove(el)
		return cn
	}
	return 0
}
