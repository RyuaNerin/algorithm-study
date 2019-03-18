// https://programmers.co.kr/learn/courses/30/lessons/12931

func solution(n int) (r int) {
    for n > 0 {
        r, n = r + n % 10, n / 10
    }
    return
}
