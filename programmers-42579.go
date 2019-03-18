// https://programmers.co.kr/learn/courses/30/lessons/42579

import "sort"

type songPlay struct {
    index int
    play  int
}

func solution(genres []string, plays []int) (r []int) {
    playCountMap := make(map[string]int)
    playIndexMap := make(map[string][]songPlay)
    for i := range genres {
        playCountMap[genres[i]] += plays[i]
        playIndexMap[genres[i]] = append(playIndexMap[genres[i]], songPlay { i, plays[i] })
    }
    
    for _, genre := range genreOrder(playCountMap) {
        r = append(r, getOrderedIndexes(playIndexMap[genre])...)
    }
    
    return
}

func getOrderedIndexes(arr []songPlay) (r []int) {
    sort.Slice(arr, func(i, k int) bool {
        if arr[i].play == arr[k].play {
            return arr[i].index < arr[k].index
        }
        return arr[i].play > arr[k].play
    })
    if len(arr) > 2 {
        arr = arr[:2]
    }
    for _, v := range arr {
        r = append(r, v.index)
    }
    return
}

func genreOrder(m map[string]int) (r []string) {
    type kv struct {
        k string
        v int
    }
    
    var arr []kv
    for k, v := range m {
        arr = append(arr, kv { k, v })
    }
    
    sort.Slice(arr, func(i, k int) bool {
        return arr[i].v > arr[k].v
    })

    for _, v := range arr {
        r = append(r, v.k)
    }
    return
}
