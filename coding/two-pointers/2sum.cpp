#include <iostream>
#include <unordered_map>
#include <vector>
using namespace std;

vector<int> twoSum(vector<int>& nums, int target) {
    unordered_map<int, int> numToIndex; // Map to store {value -> index}
    for (int i = 0; i < nums.size(); i++) {
        int complement = target - nums[i];
        cout<<" complement "<<complement<<" num to index count "<<numToIndex.count(complement)<<endl;
        if (numToIndex.count(complement)) { // Check if complement exists
            return vector<int>{numToIndex[complement], i}; // Return indices
        }
        numToIndex[nums[i]] = i; // Store the index of the current number
    }
    return vector<int>{-1, -1}; // No solution found
}

// Driver Code
int main() {
    vector<int> nums = {2, 7, 11, 15};
    int target = 9;

    vector<int> result = twoSum(nums, target);
    cout << "Indices: [" << result[0] << ", " << result[1] << "]" << endl;
    return 0;
}
