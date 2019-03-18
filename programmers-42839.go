// https://programmers.co.kr/learn/courses/30/lessons/42839
// 2 ~ 10^l 범위 돌면서 조합 가능한지 확인하여 작업함
// 무식하지만 완전탐색 그 자체

func solution(numbers string) (r int) {
    num := make([]int, len(numbers))
    for i, c := range numbers {
        num[i] = int(c - '0')
    }

    ii := 1
    for range num {
        ii *= 10
    }

    isNotPrime := make([]bool, ii)
    isNotPrime[0] = true
    isNotPrime[1] = true
    for i := 2; i < ii; i++ {
        if isNotPrime[i] {
            continue
        }
        for k := i * 2; k < ii; k += i {
            isNotPrime[k] = true
        }
    }

    flag := make([]bool, len(numbers))
    for i := 2; i < ii; i++ {
        verified := true

        for ni := range flag {
            flag[ni] = true
        }

        k := i
        for k > 0 {
            v := k % 10
            k  = k / 10

            notFound := true
            for ni, nv := range num {
                if nv == v && flag[ni] {
                    notFound = false
                    flag[ni] = false
                    break
                }
            }

            if notFound {
                verified = false
                break
            }
        }

        if verified && !isNotPrime[i] {
            r++
        }
    }

    return
}
