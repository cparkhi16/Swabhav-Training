#include <iostream>
#include <vector>
#include <stack>

using namespace std;

// Function to find the next greater element in a circular array
vector<int> nextGreaterElements(vector<int>& nums) {
    int n = nums.size();
    vector<int> ans(n, -1);  // Initialize result array with -1
    stack<int> st;

    // Iterate over the array twice to simulate circular array
    for (int i = 0; i < 2 * n; i++) {
        int nm = nums[i % n];  // Get the current element in circular manner
        
        // Pop elements from the stack while the current element is greater
        while (!st.empty() && nums[st.top()] < nm) {
            ans[st.top()] = nm;
            st.pop();
        }
        
        // Push the index to stack only for the first pass
        if (i < n) {
            st.push(i);
        }
    }

    return ans;
}

int main() {
    // Test case 1
    vector<int> nums = {1, 2, 1};
    vector<int> result = nextGreaterElements(nums);
    
    cout << "Next Greater Elements (Circular): ";
    for (int val : result) {
        cout << val << " ";
    }
    cout << endl;

    return 0;
}
