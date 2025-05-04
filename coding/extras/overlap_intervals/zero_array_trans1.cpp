#include <iostream>
#include <vector>
using namespace std;

class Solution {
public:
    bool canTransformToZeroArray(vector<int>& nums, vector<vector<int>>& queries) {
        int n = nums.size();
        vector<int> diff(n + 1, 0);

        // Apply range increments using difference array
        for (const auto& q : queries) {
            int l = q[0];
            int r = q[1];
            diff[l] += 1;
            if (r + 1 < n) diff[r + 1] -= 1;
        }

        // Build prefix sum to compute total available decrements
        vector<int> decrements(n);
        decrements[0] = diff[0];
        for (int i = 1; i < n; ++i) {
            decrements[i] = decrements[i - 1] + diff[i];
        }

        // Check if each element can be reduced to zero
        for (int i = 0; i < n; ++i) {
            if (nums[i] > decrements[i]) return false;
        }

        return true;
    }
};


int main() {
    Solution sol;

    vector<int> nums1 = {1, 0, 1};
    vector<vector<int>> queries1 = {{0, 2}};
    cout << "Test 1: " << (sol.canTransformToZeroArray(nums1, queries1) ? "true" : "false") << endl;

    vector<int> nums2 = {4, 3, 2, 1};
    vector<vector<int>> queries2 = {{1, 3}, {0, 2}};
    cout << "Test 2: " << (sol.canTransformToZeroArray(nums2, queries2) ? "true" : "false") << endl;

    return 0;
}
