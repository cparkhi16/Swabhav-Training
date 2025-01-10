#include <iostream>
#include <vector>
using namespace std;

// Function definition
int removeDuplicates(vector<int>& nums) {
    int dup = 1;

    for (int i = 1; i < nums.size(); i++) {
        if (nums[i] != nums[dup - 1]) { // Check for duplicates
            nums[dup] = nums[i];       // Update the array
            dup++;
        }
    }
    return dup;
}

// Driver code
int main() {
    vector<int> nums = {1, 1, 2, 3, 3, 4, 5, 5}; // Example sorted array

    cout << "Original array: ";
    for (int num : nums) {
        cout << num << " ";
    }
    cout << endl;

    // Call the function
    int newLength = removeDuplicates(nums);

    // Print the results
    cout << "Modified array: ";
    for (int i = 0; i < newLength; i++) {
        cout << nums[i] << " ";
    }
    cout << endl;

    cout << "New length of the array: " << newLength << endl;

    return 0;
}
