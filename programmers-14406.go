// https://programmers.co.kr/learn/courses/30/lessons/14406
// 유클리드 호제법

long long solution(int N) {
    long long answer = 0;
    
    bool* b = new bool[10000001];
    int i, k;
    for (i = 2; i <= N; i++) {
        if (b[i]) continue;
        for (k = i; k <= N; k += i)
            b[k] = true;
        answer += i;
    }
    delete[] b;
    
    return answer;
}
