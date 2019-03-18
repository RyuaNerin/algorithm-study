// https://programmers.co.kr/learn/courses/30/lessons/42841

func solution(baseball [][]int) (r int) {
    for i := 123; i <= 999; i++ {
        c0 :=  i / 100
        c1 := (i / 10) % 10
        c2 :=  i % 10

        if c0 == 0 || c1 == 0 || c2 == 0 || c0 == c1 || c0 == c2 || c1 == c2 {
            continue
        }

        ok := true

        for _, bbCase := range baseball {
            bb0 :=  bbCase[0] / 100
            bb1 := (bbCase[0] / 10) % 10
            bb2 :=  bbCase[0] % 10

            // check strike
            count := 0
            if c0 == bb0 { count++ }
            if c1 == bb1 { count++ }
            if c2 == bb2 { count++ }

            if count != bbCase[1] {
                ok = false
                break
            }

            // check ball
            count = 0
            if c0 == bb1 || c0 == bb2 { count++ }
            if c1 == bb0 || c1 == bb2 { count++ }
            if c2 == bb0 || c2 == bb1 { count++ }

            if count != bbCase[2] {
                ok = false
                break
            }
        }

        if ok {
            r++
        }
    }

    return
}
