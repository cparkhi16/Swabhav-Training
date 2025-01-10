#include <iostream>
#include <vector>

using namespace std;

class NumArray {
private:
    vector<int> arr;
    vector<int> buildPrefix() {
        vector<int> prefixSum(arr.size());
        prefixSum[0] = arr[0];
        for (int i = 1; i < arr.size(); i++) {
            prefixSum[i] = prefixSum[i - 1] + arr[i];
        }
        return prefixSum;
    }
public:
    NumArray(vector<int>& nums) {
        arr = nums;
    }
    
    int sumRange(int left, int right) {
        vector<int> prefixSum = buildPrefix();
        if (left == 0) {
            return prefixSum[right];
        }
        return prefixSum[right] - prefixSum[left - 1];
    }
};

int main() {
    // Input array
    vector<int> nums = {1, 2, 3, 4, 5, 6};

    // Instantiate the NumArray object
    NumArray* obj = new NumArray(nums);

    // Example queries
    vector<pair<int, int>> queries = {{0, 3}, {1, 4}, {2, 5}};

    // Output results of the queries
    for (const auto& query : queries) {
        int left = query.first;
        int right = query.second;
        cout << "Sum of range [" << left << ", " << right << "] = " 
             << obj->sumRange(left, right) << endl;
    }

    // Free up allocated memory
    delete obj;

    return 0;
}
