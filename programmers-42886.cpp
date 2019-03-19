// https://programmers.co.kr/learn/courses/30/lessons/42886
// gg ez

#include <vector>
#include <algorithm>

using namespace std;

int solution(vector<int> weight) {
	sort(weight.begin(), weight.end());

	int sum = 1;
	for (auto i = weight.begin(); i != weight.end(); ++i)
	{
		if (sum < *i) return sum;
		sum += *i;
	}
	return sum;
}
