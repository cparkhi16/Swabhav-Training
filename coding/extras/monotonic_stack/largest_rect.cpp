#include <iostream>
#include <vector>
#include <stack>
#include <algorithm>

using namespace std;

// Function to calculate the largest rectangle area in a histogram
int largestRectangleArea(vector<int>& heights) {
    stack<int> st;
    int maxArea = 0;
    heights.push_back(0);  // Adding a 0 height at the end to ensure we pop all bars in the stack
    
    for (int i = 0; i < heights.size(); i++) {
        // While the stack is not empty and the current height is smaller than the height at the top of the stack
        while (!st.empty() && heights[i] < heights[st.top()]) {
            int height = heights[st.top()];
            cout<<" height is "<<height<<endl;
            st.pop();
            cout<<" is st empty "<<st.empty()<<endl;
            int width = st.empty() ? i : i - st.top() - 1;
            cout<<" width is "<<width<<endl;
            maxArea = max(maxArea, height * width);  // Update max area
        }
        st.push(i);  // Push the current index to the stack
    }
    
    return maxArea;
}

int main() {
    // Test case 1: Example histogram heights
    vector<int> heights = {2, 1, 5, 6, 2, 3};
    
    int result = largestRectangleArea(heights);
    
    cout << "Largest rectangle area: " << result << endl;

    // Test case 2: Another example
    vector<int> heights2 = {2, 4};
    
    result = largestRectangleArea(heights2);
    
    cout << "Largest rectangle area: " << result << endl;

    // Test case 3: Edge case with a single bar
    vector<int> heights3 = {5};
    
    result = largestRectangleArea(heights3);
    
    cout << "Largest rectangle area: " << result << endl;

    return 0;
}
