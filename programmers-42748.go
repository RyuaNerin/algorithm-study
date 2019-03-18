// https://programmers.co.kr/learn/courses/30/lessons/42748?language=go

func solution(array []int, commands [][]int) (r []int) {
    for _, command := range commands {
        length := command[1] - command[0] + 1
        part := make([]int, length)
        copy(part, array[command[0] - 1:])
        
        sort(part)
	
        r = append(r, part[command[2] - 1])
    }
    return
}

func sort(array []int) []int {
    var tmp int
    
    for i := 0; i < len(array); i++ {
        for k := i + 1; k < len(array); k++ {
            if array[k] < array[i] {
                tmp = array[i]
                array[i] = array[k]
                array[k] = tmp
            }
        }
    }
    
    return array
}
