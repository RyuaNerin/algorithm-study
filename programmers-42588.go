// https://programmers.co.kr/learn/courses/30/lessons/42588?language=go

func solution(h []int) (r []int) {
    r = make([]int, len(h))
    
    for i := 1; i < len(h); i++ {
        for k := i - 1; k >= 0; k-- {
            if h[k] > h[i] {
                r[i] = k + 1
                break
            }
        }
    }
    return
}
