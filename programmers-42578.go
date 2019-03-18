func solution(clothes [][]string) int {
    m := make(map[string]int)
    for _, v := range clothes {
        m[v[1]]++
    }
    r := 1
    for _, c := range m {
        r *= (c + 1)
    }
    return r - 1
}
