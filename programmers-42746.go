// https://programmers.co.kr/learn/courses/30/lessons/42746
// string 으로 변환 해서 붙이고 비교하면 편한데 이렇게 함 해보고싶었음

import (
    "math"
    "sort"
    "strconv"
    "strings"
)

func solution(numbers []int) (r string) {
    sort.Slice(numbers, func(i, k int) bool {
        if numbers[i] == 0 { return false }
        if numbers[k] == 0 { return true }

        di := math.Log10(float64(numbers[i]))
        dk := math.Log10(float64(numbers[k]))
        if di == float64(int(di)) { di = di + 1 } else { di = math.Ceil(di) }
        if dk == float64(int(dk)) { dk = dk + 1 } else { dk = math.Ceil(dk) }
        if di == dk {
            return numbers[i] > numbers[k]
        }

        ni := numbers[i] * int(math.Pow(10, dk)) + numbers[k]
        nk := numbers[k] * int(math.Pow(10, di)) + numbers[i]

        return ni > nk
    })

    // numbers의 원소는 0 이상 1,000 이하입니다.
    isNotZero := false

    var sb strings.Builder
    sb.Grow(4 * len(numbers))
    for _, v := range numbers {
        if v != 0 {
            isNotZero = true
        }
        sb.WriteString(strconv.Itoa(v))
    }
    if isNotZero {
        return sb.String()
    } else {
        return ""
    }
}
