//https://programmers.co.kr/learn/courses/30/lessons/42883

func solution(number string, k int) string {
    //  0 1 2 3 4 5 6 7 8 9     < index     Length = `nLen` = 10
    // [4 1 7 7 2 5 2 8 4 1]    Result      Length = `nLen` - k = `rLen`
    // [4 1 7 7 2]              7.....
    //  .[. . 7 2 5]            77....
    //  . .[. . 2 5 2]          775...
    //  . . .[. . . 2 8]        7758..
    //  . . . .[. . . . 4]      77584.
    //  . . . . .[. . . . 1]    775841
    //            1 2 3 4 5            Part Length = `nLen` - (`rLen` - 1) = `nLen` - (`nLen` - k - 1) = `k` + 1

    nLen := len(number)
    r    := make([]byte, nLen - k)
    pLen   := k + 1
    
    lastIndex := 0
    for i, _ := range r {
        maxIndex := lastIndex
        maxValue := number[lastIndex]
        for m := lastIndex; m < i + pLen && m < nLen; m++ {
            if number[m] > maxValue {
                maxIndex = m
                maxValue = number[m]
            }
        }
        
        r[i] = number[maxIndex]
        lastIndex = maxIndex + 1
    }

    return string(r)
}
