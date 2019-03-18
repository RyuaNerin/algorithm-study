// https://programmers.co.kr/learn/courses/30/lessons/42895
// 뺑뺑이조아

func solution(N int, number int) (r int) {
    r = 10
    recursive(int64(N), int64(number), 0, 0, &r)
    if r == 10 {
        r = -1
    }
    return
}

func recursive(N int64, dnum int64, cnum int64, count int, result *int) {
    if cnum == dnum {
        if count < *result {
            *result = count
        }
        return
    }
    if count == 8 {
        return
    }

    for i := 0; i < (8 - count); i++ {
        subNum := makeNumber(N, i + 1)

        recursive(N, dnum, cnum + subNum, count + i + 1, result)
        recursive(N, dnum, cnum - subNum, count + i + 1, result)
        recursive(N, dnum, cnum * subNum, count + i + 1, result)

        if cnum % subNum == 0 {
            recursive(N, dnum, cnum / subNum, count + i + 1, result)
        }
    }
}

func makeNumber(N int64, c int) (r int64) {
    var v int64
    r = 0
    for i := 0; i < c; i++ {
        v = 1
        for k := 0; k < i; k++ { v *= 10 }
        r += v * N
    }

    return
}
