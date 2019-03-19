// https://programmers.co.kr/learn/courses/30/lessons/42584

#include <vector>

using namespace std;

vector<int> solution(vector<int> prices)
{
	vector<int> r;
	int len = prices.size() - 1;

	for (int i = 0; i < len; i++)
	{
		int v = 1;
		for (int k = i + 1; k < len; k++) {
			if (prices[i] <= prices[k])
				v++;
			else
				break;
		}

		r.push_back(v);
	}
    r.push_back(0);

	return r;
}
