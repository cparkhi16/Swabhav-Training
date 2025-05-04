#include <iostream>
#include <vector>
#include <unordered_map>
using namespace std;

int longestSubarrayDivByK(const vector<int>& nums, int K) {
    unordered_map<int, int> modIndexMap;
    modIndexMap[0] = -1;  // Handle subarrays starting from index 0

    int prefix_sum = 0;
    int max_len = 0;

    for (int i = 0; i < nums.size(); ++i) {
        prefix_sum += nums[i];
        int mod = ((prefix_sum % K) + K) % K;  // Normalize negative mods

        if (modIndexMap.find(mod) != modIndexMap.end()) {
            max_len = max(max_len, i - modIndexMap[mod]);
        } else {
            modIndexMap[mod] = i;
        }
    }
    return max_len;
}

int main() {
    vector<int> nums = {2, 7, 6, 1, 4, 5};
    int K = 3;

    int result = longestSubarrayDivByK(nums, K);
    cout << "Length of longest subarray divisible by " << K << " is: " << result << endl;

    return 0;
}
