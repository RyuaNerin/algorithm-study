// https://programmers.co.kr/learn/courses/30/lessons/42586?language=go

func solution(progresses []int, speeds []int) (result []int) {
    days := make([]int, len(progresses))
    for i, p := range progresses {
        days[i] = 100 - p
        if days[i] % speeds[i] == 0 {
            days[i] = days[i] / speeds[i]
        } else {
            days[i] = days[i] / speeds[i] + 1
        }
    }
    
    i := 1
    for i < len(days) {
        if days[i] < days[i - 1] {
            days[i] = days[i - 1]
        }
        i++
    }
    
    last := -1
    for _, v := range days {
        if last != v {
            result = append(result, 0)
            last = v
        }
        result[len(result) - 1]++
    }
    return
}
