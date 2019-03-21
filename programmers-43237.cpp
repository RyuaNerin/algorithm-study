// https://programmers.co.kr/learn/courses/30/lessons/43237
// 효율성 테스트 2번은 입력된 정답이 오답인거같음

#include <vector>
#include <algorithm>

using namespace std;

int solution(vector<int> budgets, int M)
{
    int len = budgets.size();
    int max = 0;
    int sum = 0;
    for (auto i : budgets)
    {
        sum += i;
        if (i > max)
            max = i;
    }
    if (sum <= M)
        return max;

    sort(budgets.begin(), budgets.end());

    max = M / len;
    sum = budgets[0];
    for (int i = 1; i < len; ++i)
    {
        // 계산
        int newMax = (M - sum) / (len - i);

        if (newMax < budgets[i - 1])
            break;

        if ((sum + (len - i) * newMax) > M)
            break;

        sum += budgets[i];
        max = newMax;
    }

    return max;
}
