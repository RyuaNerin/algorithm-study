// https://programmers.co.kr/learn/courses/30/lessons/42628

import "strconv"

func solution(operations []string) []int {
    q := make([]int, 0, len(operations))

    find := func(findMax bool) (index int) {
        value := q[0]
        for i, v := range q {
            if (findMax && v > value) || (!findMax && v < value) {
                index = i
                value = v
            }
        }
        return
    }

    for _, v := range operations {
        op := v[0]
        nu, _ := strconv.Atoi(v[2:])
        if op == 'I' {
            q = append(q, nu)
        } else {
            if len(q) > 0 {
                i := find(nu == 1)
		q = append(q[:i], q[i+1:]...)
            }
        }
    }

    if len(q) == 0 {
        return []int { 0, 0 }
    } else {
        max := find(true)
        min := find(false)

        return []int { q[max], q[min] }
    }
}
