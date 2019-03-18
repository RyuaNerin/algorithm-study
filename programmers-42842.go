// https://programmers.co.kr/learn/courses/30/lessons/42842
// 완전탐색 문제니까 완전탐색으로

func solution(brown int, red int) []int {
    xpy := brown / 2 + 2
    for x := 3; x <= xpy; x++ {
        for y := 3; y <= x; y++ {
            if red == (x - 2) * (y - 2) && brown == (x + y) * 2 - 4 {
                return []int { x, y }
            }
        }
    }

    return []int { 0, 0} 
}
