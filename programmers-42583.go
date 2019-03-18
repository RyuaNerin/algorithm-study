// https://programmers.co.kr/learn/courses/30/lessons/42583

func solution(bridge_length int, weight int, truck_weights []int) (sec int) {
    type ts struct {
        weight  int
        pos     int
    }
    var onBridge []ts
    for {
        sec++
        i := 0
        for i < len(onBridge) {
            onBridge[i].pos++

            if onBridge[i].pos == bridge_length {
                onBridge = onBridge[1:]
                continue
            }
            i++
        }
        if len(onBridge) == 0 && len(truck_weights) == 0 {
            break
        }
        allowance_weight := 0
        for _, v := range onBridge {
            allowance_weight += v.weight
        }
        allowance_weight = weight - allowance_weight
        if len(truck_weights) > 0 {
            if truck_weights[0] <= allowance_weight {
                onBridge = append(onBridge, ts { weight: truck_weights[0] })
                truck_weights = truck_weights[1:]
            }
        }
    }
    return
}
