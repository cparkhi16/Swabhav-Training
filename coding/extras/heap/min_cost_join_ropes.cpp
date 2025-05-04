#include <iostream>
#include <queue>
#include <vector>

using namespace std;

int getMinimumCost(vector<int>& nums) {
    priority_queue<int, vector<int>, greater<int>> minHeap;  // Min-heap

    // Insert all ropes into the min heap
    for (int i : nums) {
        minHeap.push(i);
    }

    int cost = 0;

    // Process while more than one rope remains
    while (minHeap.size() > 1) {  
        // Extract the two smallest ropes
        int top = minHeap.top();
        minHeap.pop();
        int nxt = minHeap.top();
        minHeap.pop();

        // Compute the cost and push the new rope
        cost += top + nxt;
        minHeap.push(top + nxt);
    }

    return cost;
}

int main() {
    vector<int> nums = {4, 3, 2, 6};
    int c = getMinimumCost(nums);
    cout << "Minimum cost to join the ropes: " << c << endl;
    return 0;
}
