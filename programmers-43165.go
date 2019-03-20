// https://programmers.co.kr/learn/courses/30/lessons/43165

func solution(numbers []int, target int) (count int) {
    var recursive func(index int, num int)
    recursive = func(index int, num int) {
        if index == len(numbers) {
            if num == target {
                count++
            }
            return
        }

        recursive(index + 1, num + numbers[index])
        recursive(index + 1, num - numbers[index])
    }

    recursive(0, 0)

    return
}
