#include <iostream>
#include <vector>
using namespace std;

// Function definition
int removeElement(vector<int>& nums, int val) {
    int n = 0;
    for (int i = 0; i < nums.size(); i++) {
        if (nums[i] != val) {
            nums[n] = nums[i];
            n++;
        }
    }
    return n;
}

// Driver code
int main() {
    vector<int> nums = {3, 2, 2, 3, 4, 3}; // Example array
    int val = 3; // Value to remove

    cout << "Original array: ";
    for (int num : nums) {
        cout << num << " ";
    }
    cout << endl;

    // Call the function
    int newLength = removeElement(nums, val);

    // Print the results
    cout << "Modified array: ";
    for (int i = 0; i < newLength; i++) {
        cout << nums[i] << " ";
    }
    cout << endl;

    cout << "New length of the array: " << newLength << endl;

    return 0;
}
