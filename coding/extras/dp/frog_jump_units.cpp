#include <iostream>
#include <unordered_map>
#include <vector>

using namespace std;

bool canCross(vector<int>& stones) {
    unordered_map<int, unordered_map<int, bool>> dp;
    dp[0][0] = true; // The frog starts on the first stone and the first jump is 0 (i.e., no jump yet)

    for (int i = 0; i < stones.size(); i++) {
        for (auto& kv : dp[stones[i]]) {
            int jump = kv.first;
            for (int nextJump = jump - 1; nextJump <= jump + 1; nextJump++) {
                if (nextJump > 0) {
                    int nextStone = stones[i] + nextJump;
                    if (nextStone == stones.back()) return true;
                    dp[nextStone][nextJump] = true;
                }
            }
        }
    }

    return false;
}

int main() {
    vector<int> stones1 = {0,1,3,5,6,8,12,17};
    vector<int> stones2 = {0,1,2,3,4,8,9,11};
    
    cout << (canCross(stones1) ? "true" : "false") << endl; // Output: true
    cout << (canCross(stones2) ? "true" : "false") << endl; // Output: false

    return 0;
}
