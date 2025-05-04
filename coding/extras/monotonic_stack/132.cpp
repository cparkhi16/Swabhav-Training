#include <iostream>
#include <vector>
#include <stack>
#include <climits>

using namespace std;

// Function to check for 132 pattern
bool find132pattern(vector<int>& nums) {
    int third = INT_MIN;
    stack<int> st;

    // Traverse the array from right to left
    for (int i = nums.size() - 1; i >= 0; i--) {
        if (nums[i] < third) {
            return true;  // Found a valid 132 pattern
        }
        
        // Update the third value and pop elements from the stack
        while (!st.empty() && st.top() < nums[i]) {
            third = st.top();
            st.pop();
        }

        // Push the current element onto the stack
        st.push(nums[i]);
    }

    return false;  // No valid 132 pattern found
}

int main() {
    // Test case 1: Example with a valid 132 pattern
    vector<int> nums = {1, 2, 3, 4};
    if (find132pattern(nums)) {
        cout << "Found a 132 pattern!" << endl;
    } else {
        cout << "No 132 pattern found!" << endl;
    }

    // Test case 2: Example without a valid 132 pattern
    vector<int> nums2 = {3, 1, 4, 2};
    if (find132pattern(nums2)) {
        cout << "Found a 132 pattern!" << endl;
    } else {
        cout << "No 132 pattern found!" << endl;
    }

    // Test case 3: Another example with a valid 132 pattern
    vector<int> nums3 = {-1, 3, 2, 0};
    if (find132pattern(nums3)) {
        cout << "Found a 132 pattern!" << endl;
    } else {
        cout << "No 132 pattern found!" << endl;
    }

    return 0;
}
