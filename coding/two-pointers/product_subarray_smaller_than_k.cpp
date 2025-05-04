#include <iostream>
#include <vector>
using namespace std;

vector<vector<int>> subarrayProductLessThanK(vector<int>& nums, int k) {
    int start = 0, product = 1, count = 0;
    vector<vector<int>> result; // To store the valid subarrays
    vector<vector<int>> tp;
    for (int end = 0; end < nums.size(); end++) {
        product *= nums[end];

        // Shrink the window if the product exceeds or equals `k`
        while (start <= end && product >= k) {
            product /= nums[start];
            start++;
        }
        for(int i = start ; i <= end ; i++){
            vector<int> subarray(nums.begin() + i, nums.begin() + end + 1);
            tp.push_back(subarray);
        }
        // Add all valid subarrays ending at `end`
        for (int i = end; i >= start; i--) {
            cout<<" end "<<end<<" start "<<start<<endl;
            vector<int> subarray(nums.begin() + i, nums.begin() + end + 1);
            result.push_back(subarray);
        }

        // Update count (not strictly needed for the task, but for reference)
        count += (end - start + 1);
    }

    cout << "Valid subarrays with product less than " << k << ":\n";
    for (const auto& subarray : tp) {
        cout << "[";
        for (int num : subarray) {
            cout << num << " ";
        }
        cout << "]\n";
    }

    return result;
}

int main() {
    vector<int> nums = {10, 5, 2, 6};
    int k = 100;

    vector<vector<int>> validSubarrays = subarrayProductLessThanK(nums, k);

    // Print the valid subarrays
    cout << "Valid subarrays with product less than " << k << ":\n";
    for (const auto& subarray : validSubarrays) {
        cout << "[";
        for (int num : subarray) {
            cout << num << " ";
        }
        cout << "]\n";
    }

    return 0;
}
