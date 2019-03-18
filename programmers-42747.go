// https://programmers.co.kr/learn/courses/30/lessons/42747

import (
    "sort"
)

func solution(citations []int) int {
    sort.Slice(citations, func(i, k int) bool {
        return citations[i] > citations[k]
    })

    for h := len(citations); h > 0; h-- {
        if sort.Search(len(citations), func(i int) bool { return citations[i] < h }) >= h {
            return h
        }
    }

    return 0
}
