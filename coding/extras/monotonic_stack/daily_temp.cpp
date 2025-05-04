#include <iostream>
#include <vector>
#include <stack>

using namespace std;

// Function to solve the Daily Temperatures problem
vector<int> dailyTemperatures(vector<int>& temperatures) {
    vector<int> ans(temperatures.size(), 0);
    stack<int> st;
    
    // Iterate through all temperatures
    for (int i = 0; i < temperatures.size(); i++) {
        // Pop elements from the stack while the current temperature is higher
        // than the temperature at the index stored at the top of the stack
        while (!st.empty() && temperatures[st.top()] < temperatures[i]) {
            int prevInd = st.top();
            ans[prevInd] = i - prevInd;  // Calculate the difference in days
            st.pop();  // Pop the index from the stack
        }
        st.push(i);  // Push the current index onto the stack
    }
    
    return ans;  // Return the answer
}

int main() {
    // Test case 1
    vector<int> temperatures = {73, 74, 75, 71, 69, 72, 76, 73};
    
    vector<int> result = dailyTemperatures(temperatures);
    
    cout << "Days to wait for warmer temperatures: ";
    for (int val : result) {
        cout << val << " ";
    }
    cout << endl;

    return 0;
}
