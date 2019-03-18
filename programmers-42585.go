// https://programmers.co.kr/learn/courses/30/lessons/42585?language=go

func solution(arrangement string) (r int) {
	var stack []int
	for _, v := range arrangement {
		switch v {
		case '(':
			stack = append(stack, 1)
		case ')':
			last_count := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if last_count == 1 {
				for i, _ := range stack {
					stack[i]++
				}
			} else {
				r += last_count
			}
		}
	}
	return
}
