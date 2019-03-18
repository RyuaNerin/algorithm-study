// https://programmers.co.kr/learn/courses/30/lessons/42587

func solution(priorities []int, location int) (index int) {
    index++
    type s struct {
        p int
        l bool
    }
    q := make([]s, len(priorities))
    for i, v := range priorities {
        q[i] = s { v, i == location }
    }
    for len(q) > 0 {
        f := q[0]
        copy(q[0:], q[1:])
        q = q[:len(q) - 1]
        
        pushBack := false
        for _, v := range q {
            if v.p > f.p {
                pushBack = true
                break
            }
        }
        if pushBack {
            q = append(q, f)
        } else {
            if f.l {
                return
            } else {
                index++
            }
        }
    }
    return
}
