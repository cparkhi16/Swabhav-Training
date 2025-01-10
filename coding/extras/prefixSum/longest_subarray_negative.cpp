#include <iostream>
#include <vector>
#include <unordered_map>

using namespace std;

int longestSubarrayWithSumK(vector<int>& nums, int k) {
    int maxLength = 0;
    int currentSum = 0;
    unordered_map<int, int> sumIndexMap;

    sumIndexMap[0] = -1; // To handle the case when the sum becomes k itself

    for (int i = 0; i < nums.size(); ++i) {
        currentSum += nums[i];

        if (sumIndexMap.find(currentSum - k) != sumIndexMap.end()) {
            maxLength = max(maxLength, i - sumIndexMap[currentSum - k]);
        }

        if (sumIndexMap.find(currentSum) == sumIndexMap.end()) {
            sumIndexMap[currentSum] = i;
        }
    }

    return maxLength;
}

int main() {
    vector<int> nums = {1, -1, 5, -2, 3};
    int k = 3; // ans is 4
    cout << "Length of the longest subarray with sum " << k << ": " << longestSubarrayWithSumK(nums, k) << endl;
    return 0;
}
