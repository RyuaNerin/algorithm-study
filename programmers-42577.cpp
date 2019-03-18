// https://programmers.co.kr/learn/courses/30/lessons/42577

#include <string>
#include <vector>

using namespace std;

bool solution(vector<string> pb) {
    for (auto i = 0; i < pb.size(); ++i) {
        for (auto k = 0; k < pb.size(); ++k) {
            if (i == k) continue;
            if (pb[k].find(pb[i]) == 0)
                return false;
        }
    }
    return true;
}
