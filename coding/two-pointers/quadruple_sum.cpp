#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

vector<vector<int>> fourSum(vector<int>& nums, int target) {
    vector<vector<int>> result;
    int n = nums.size();
    
    if (n < 4) return result; // Less than 4 elements, no quadruplets possible.

    sort(nums.begin(), nums.end()); // Sort the array.

    // First loop for the first element.
    for (int i = 0; i < n - 3; i++) {
        if (i > 0 && nums[i] == nums[i - 1]) continue; // Skip duplicates for the first element.

        // Second loop for the second element.
        for (int j = i + 1; j < n - 2; j++) {
            if (j > i + 1 && nums[j] == nums[j - 1]) continue; // Skip duplicates for the second element.

            int left = j + 1, right = n - 1;

            // Two-pointer approach for the remaining two elements.
            while (left < right) {
                long long sum = (long long)nums[i] + nums[j] + nums[left] + nums[right]; // Use long long to avoid overflow.

                if (sum == target) {
                    result.push_back({nums[i], nums[j], nums[left], nums[right]});

                    // Skip duplicates for the third and fourth elements.
                    while (left < right && nums[left] == nums[left + 1]) left++;
                    while (left < right && nums[right] == nums[right - 1]) right--;

                    left++;
                    right--;
                } else if (sum < target) {
                    left++;
                } else {
                    right--;
                }
            }
        }
    }

    return result;
}

int main() {
    vector<int> nums = {1, 0, -1, 0, -2, 2};
    int target = 0;

    vector<vector<int>> result = fourSum(nums, target);

    cout << "Quadruplets that sum to " << target << " are:\n";
    for (const auto& quad : result) {
        cout << "[";
        for (int num : quad) {
            cout << num << " ";
        }
        cout << "]\n";
    }

    return 0;
}
