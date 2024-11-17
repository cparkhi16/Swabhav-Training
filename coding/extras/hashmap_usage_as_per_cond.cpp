#include <iostream>
#include <vector>
#include <unordered_map>

using namespace std;

bool containsNearbyDuplicate(vector<int>& nums, int k) {
    unordered_map<int, int> indexMap;  // Map to store the element and its index
    
    for (int i = 0; i < nums.size(); ++i) {
        if (indexMap.find(nums[i]) != indexMap.end()) {
            // Check if the previous index of this element is within the distance k
            if (i - indexMap[nums[i]] <= k) {
                return true;
            }
        }
        // Update the index of the current element in the map
        indexMap[nums[i]] = i;
    }
    
    return false;  // No such pair found
}

int main() {
    vector<int> nums = {1, 2, 3, 1};
    int k = 3;
    bool result = containsNearbyDuplicate(nums, k);
    cout << (result ? "True" : "False") << endl;
    return 0;
}
