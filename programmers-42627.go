// https://programmers.co.kr/learn/courses/30/lessons/42627
// golang 기본 패키지엔 PriorityQueue 가 없어서 정렬을 이용함

import "sort"

func solution(jobs [][]int) (int) {
    var ts uint64

    sort.Slice(jobs, func (i, k int) bool {
        if jobs[i][1] != jobs[k][1] {
            return jobs[i][1] < jobs[k][1]
        } else {
            return jobs[i][0] < jobs[k][0]
        }
    })

    count := uint64(len(jobs))
    sec := 0
    for len(jobs) > 0 {
        found := false
        for i, job := range jobs {
            if job[0] <= sec {
                sec += job[1]
                ts  += uint64(sec - job[0])

                copy(jobs[i:], jobs[i + 1:])
                jobs = jobs[:len(jobs) - 1]

                found = true
                break
            }
        }

        if !found {
            sec++
        }
    }

    return int(ts / count)
}
