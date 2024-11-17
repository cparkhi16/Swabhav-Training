#include <iostream>
#include <vector>

using namespace std;

void dfs(int current, int n, vector<int>& result) {
    if (current > n) return; // Stop when the current number exceeds n
    
    result.push_back(current); // Add current number to the result
    
    // Try appending digits 0 to 9 to the current number
    for (int i = 0; i <= 9; i++) {
        int next = current * 10 + i;
        //if (next > n) return; // Stop recursion if the next number exceeds n
        dfs(next, n, result); // Continue exploring the next number
    }
}

vector<int> lexicalOrder(int n) {
    vector<int> result;
    
    // Start DFS from numbers 1 to 9
    for (int i = 1; i <= 9; i++) {
        if (i > n) break; // Stop if the starting number exceeds n
        dfs(i, n, result);
    }
    
    return result;
}

int main() {
    int n = 13;
    vector<int> result = lexicalOrder(n);
    
    for (int num : result) {
        cout << num << " ";
    }
    
    return 0;
}
