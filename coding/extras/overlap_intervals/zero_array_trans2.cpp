#include <iostream>
#include <vector>
using namespace std;

class Solution {
public:
    // Check if first k queries can turn nums into a zero array
    bool isValid(int k, const vector<int>& nums, const vector<vector<int>>& queries) {
        int n = nums.size();
        vector<int> diff(n + 1, 0);

        for (int i = 0; i < k; ++i) {
            int l = queries[i][0];
            int r = queries[i][1];
            int val = queries[i][2];
            diff[l] += val;
            if (r + 1 < n) diff[r + 1] -= val;
        }

        vector<int> available(n);
        available[0] = diff[0];
        for (int i = 1; i < n; ++i) {
            available[i] = available[i - 1] + diff[i];
        }

        for (int i = 0; i < n; ++i) {
            if (available[i] < nums[i]) return false;
        }
        return true;
    }

    int minimumQueriesToZeroArray(vector<int>& nums, vector<vector<int>>& queries) {
        int low = 0, high = queries.size(), ans = -1;

        while (low <= high) {
            int mid = (low + high) / 2;
            if (isValid(mid, nums, queries)) {
                ans = mid;
                high = mid - 1;
            } else {
                low = mid + 1;
            }
        }

        return ans;
    }
};

// 🧪 Driver Code
int main() {
    Solution sol;

    // Test 1
    vector<int> nums1 = {2, 0, 2};
    vector<vector<int>> queries1 = {{0, 2, 1}, {0, 2, 1}, {1, 1, 3}};
    int result1 = sol.minimumQueriesToZeroArray(nums1, queries1);
    cout << "Test 1 Output: " << result1 << endl; // Expected: 2

    // Test 2
    vector<int> nums2 = {4, 3, 2, 1};
    vector<vector<int>> queries2 = {{1, 3, 2}, {0, 2, 1}};
    int result2 = sol.minimumQueriesToZeroArray(nums2, queries2);
    cout << "Test 2 Output: " << result2 << endl; // Expected: -1

    return 0;
}
