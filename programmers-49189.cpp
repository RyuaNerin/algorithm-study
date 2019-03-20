// https://programmers.co.kr/learn/courses/30/lessons/49189

#include <vector>
#include <queue>

using namespace std;

int solution(int n, vector<vector<int>> edge) {
	auto table = new vector<int>[n];
    for (auto v : edge) {
        table[v[0] - 1].push_back(v[1] - 1);
        table[v[1] - 1].push_back(v[0] - 1);
    }

    auto worked = new bool[n] { true, false, };
    worked[0] = true;

    auto q  = new queue<int>();
    auto qn = new queue<int>();

    int result;
    q->push(0);
    while (true) {
        if (q->size() == 0) {
            if (qn->size() == 0) {
                break;
            }

            auto t = q;
            q  = qn;
            qn = t;

            result = q->size();
        }

        auto c = q->front();
        q->pop();
        for (auto v : table[c]) {
            if (!worked[v]) {
                worked[v] = true;
                qn->push(v);
            }
        }
    }

    delete[] table;
    delete[] worked;
    delete q;
    delete qn;

    return result;
}
