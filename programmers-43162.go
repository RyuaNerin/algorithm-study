// https://programmers.co.kr/learn/courses/30/lessons/43162

func solution(n int, computers [][]int) (count int) {
    visitFlag := make([]bool, n)

    queue := make([]int, 0, n)

    for {
        vindex := -1
        for i, v := range visitFlag {
            if !v {
                vindex = i
                break
            }
        }
        if vindex == -1 {
            break
        }
        count++

        queue = append(queue, vindex)
        for len(queue) > 0 {
            fmt.Println(queue)

            // dequeue
            cindex := queue[0]
            copy(queue[0:], queue[1:])
            queue = queue[:len(queue) - 1]

            visitFlag[cindex] = true

            // find
            for nindex, v := range computers[cindex] {
                if v == 1 && nindex != cindex && !visitFlag[nindex] {
                    queue = append(queue, nindex)
                    visitFlag[nindex] = true
                }
            }
        }
    }

    return
}
