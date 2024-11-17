#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

int minMovesToEqualArray(vector<int>& nums) {
    sort(nums.begin(), nums.end()); // Step 1: Sort the array
    
    int median = nums[nums.size() / 2]; // Step 2: Find the median
    
    int moves = 0;
    for (int num : nums) {
        moves += abs(num - median); // Step 3: Calculate total moves
    }
    
    return moves;
}

int main() {
    vector<int> nums = {1, 2, 3};
    cout << "Minimum number of moves: " << minMovesToEqualArray(nums) << endl;
    return 0;
}
