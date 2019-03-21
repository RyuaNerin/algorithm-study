// https://www.acmicpc.net/problem/2512

#include <iostream>
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

int main()
{
    int count;
    cin >> count;
    vector<int> budgets(count);

    int bg;
    for (int i = 0; i < count; ++i)
    {
        cin >> bg;
        budgets[i] = bg;
    }

    int k;
    cin >> k;

    cout << solution(budgets, k);
}
