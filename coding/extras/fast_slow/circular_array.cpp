#include <iostream>
#include <vector>
using namespace std;
//https://leetcode.com/problems/circular-array-loop/description/
class Solution {
public:
    bool circularArrayLoop(vector<int>& nums) {
        int n = nums.size();

        auto nextIndex = [&](int i) {
            int nt = (nums[i] + i) % n;
            return nt >= 0 ? nt : nt + n;
        };

        for (int i = 0; i < n; i++) {
            if (nums[i] == 0) continue; // Already visited
            
            int fast = i, slow = i;

            // Detect cycle
            while ((nums[slow] * nums[nextIndex(slow)]) > 0 &&
                   (nums[fast] * nums[nextIndex(fast)]) > 0 &&
                   (nums[fast] * nums[nextIndex(nextIndex(fast))]) > 0) {

                slow = nextIndex(slow);
                fast = nextIndex(nextIndex(fast));

                if (slow == fast) { // Potential cycle detected
                    if (slow == nextIndex(slow)) break; // Single-element loop, invalid
                    return true;
                }
            }

            // Mark elements as visited after finishing the cycle check
            int sign = nums[i];
            int current = i;
            while (nums[current] * sign > 0) {
                int next = nextIndex(current);
                nums[current] = 0; // Mark as visited
                current = next;
            }
        }
        return false;
    }
};

int main() {
    Solution solution;

    // Test cases
    vector<vector<int>> testCases = {
        {2, -1, 1, 2, 2},  // Expected: true (valid cycle exists: 0 -> 2 -> 3 -> 0)
        {-1, 2},           // Expected: false (no valid cycle)
        {-2, 1, -1, -2, -2}, // Expected: false (no valid cycle)
        {1, 1, 1, 1, 1},   // Expected: true (valid cycle exists: any index forms a cycle)
        {1, -1, 1, -1},    // Expected: false (no valid cycle)
        {0},               // Expected: false (only one element, no cycle possible)
    };

    // Execute test cases
    for (int i = 0; i < testCases.size(); ++i) {
        cout << "Test Case " << i + 1 << ": ";
        bool result = solution.circularArrayLoop(testCases[i]);
        cout << (result ? "true" : "false") << endl;
    }

    return 0;
}
