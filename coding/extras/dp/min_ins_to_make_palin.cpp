#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

int minInsertions(string s) {
    string t = s;
    reverse(t.begin(), t.end());
    int n = s.length();
    int m = t.length();
    vector<vector<int>> dp(n + 1, vector<int>(m + 1, 0));

    for (int i = 1; i < n + 1; i++) {
        for (int j = 1; j < m + 1; j++) {
            if (s[i - 1] == t[j - 1]) {
                dp[i][j] = 1 + dp[i - 1][j - 1];
            } else {
                dp[i][j] = max(dp[i][j - 1], dp[i - 1][j]);
            }
        }
    }
    return n - dp[n][m];
}

int main() {
    string s="mbadm";
   
    cout << "Minimum insertions needed to make '" << s << "' a palindrome: " << minInsertions(s) << endl;
    return 0;
}
