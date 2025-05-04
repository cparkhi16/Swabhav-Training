#include <iostream>
#include <vector>
#include <unordered_map>
#include <stack>

using namespace std;

// Your nextGreaterElement function
vector<int> nextGreaterElement(vector<int>& nums1, vector<int>& nums2) {
    unordered_map<int, int> nge;
    stack<int> st;

    for (int n : nums2) {
        while (!st.empty() && st.top() < n) {
            nge[st.top()] = n;
            st.pop();
        }
        st.push(n);
    }

    while (!st.empty()) {
        nge[st.top()] = -1;
        st.pop();
    }

    vector<int> ans;
    for (int b : nums1) {
        ans.push_back(nge[b]);
    }
    return ans;
}

int main() {
    // Test case 1
    vector<int> nums1 = {4, 1, 2};
    vector<int> nums2 = {1, 3, 4, 2};
    
    vector<int> result = nextGreaterElement(nums1, nums2);
    
    cout << "Result: ";
    for (int val : result) {
        cout << val << " ";
    }
    cout << endl;

    return 0;
}
