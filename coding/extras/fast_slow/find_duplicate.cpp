#include <iostream>
#include <vector>

using namespace std;

// Function to find the duplicate number
int findDuplicate(vector<int>& nums) {
    int fast = nums[0];
    int slow = nums[0];

    do {
        slow = nums[slow];
        fast = nums[nums[fast]];
    } while (fast != slow);

    slow = nums[0];
    while (slow != fast) {
        slow = nums[slow];
        fast = nums[fast];
    }
    return slow;
}

int main() {
    // Test case
    vector<int> nums = {3, 1, 3, 4, 2};  // Example input where 3 is duplicated

    cout << "Duplicate number: " << findDuplicate(nums) << endl;

    return 0;
}
