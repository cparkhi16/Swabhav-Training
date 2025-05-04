#include <iostream>
#include <vector>
#include <stack>
using namespace std;

// Function to calculate trapped rainwater
int trap(vector<int>& height) {
    int n = height.size();
    if (n == 0) return 0;

    stack<int> st;  // Stack to store indices
    int water = 0;

    for (int i = 0; i < n; ++i) {
        // While the current bar is taller than the bar at the top of the stack
        while (!st.empty() && height[i] > height[st.top()]) {
            int top = st.top();  // The index of the "bottom" bar
            st.pop();

            if (st.empty()) break;  // If the stack is empty, no left boundary exists

            int left = st.top();       // The index of the "left" boundary
            int width = i - left - 1;  // Distance between the left and current bars
            int boundedHeight = min(height[left], height[i]) - height[top];

            water += width * boundedHeight;  // Calculate trapped water
        }
        st.push(i);  // Push the current index onto the stack
    }

    return water;
}

int main() {
    // Test cases
    vector<vector<int>> testCases = {
        {0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},  // Example 1
        {4, 2, 0, 3, 2, 5},                    // Example 2
        {0, 0, 0, 0},                          // No trapping water
        {3, 0, 3},                             // Simple trap
        {1, 2, 3, 4, 5},                       // Increasing height
        {5, 4, 3, 2, 1},                       // Decreasing height
        {}                                     // Empty input
    };

    for (int i = 0; i < testCases.size(); ++i) {
        cout << "Test Case " << i + 1 << ": ";
        vector<int>& height = testCases[i];

        // Print input
        cout << "Height = [";
        for (size_t j = 0; j < height.size(); ++j) {
            cout << height[j] << (j < height.size() - 1 ? ", " : "");
        }
        cout << "]\n";

        // Calculate trapped water
        int result = trap(height);

        // Print result
        cout << "Trapped Water = " << result << " units\n\n";
    }

    return 0;
}
