// https://programmers.co.kr/learn/courses/30/lessons/42626
// c++ 시러 go 조아

#include <vector>
#include <queue>

using namespace std;

int solution(vector<int> scoville, int K) {
	priority_queue<int, vector<int>, greater<int>> q;

	for (auto v = scoville.begin(); v != scoville.end(); ++v)
		q.push(*v);


	int count = 0;

	int a0, a1;
	while (q.top() < K && q.size() > 1) {
		a0 = q.top(); q.pop();
		a1 = q.top(); q.pop();

		q.push(a0 + a1 * 2);
		count++;
	}

	if (q.top() < K)
		return -1;
	else
		return count;
}
