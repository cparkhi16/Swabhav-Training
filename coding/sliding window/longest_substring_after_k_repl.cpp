// Online C++ compiler to run C++ program online
#include <iostream>
using namespace std;
#include <bits/stdc++.h>
int main() {
    std::string s = "abccde";
    int k = 1;
    int windowStart = 0;
    unordered_map<char ,int> freqMap;
    int maxRepeatingChar = 0;
    int maxLen = 0;
    for(int windowEnd = 0 ; windowEnd < s.length() ; windowEnd++){
        freqMap[s[windowEnd]] ++;
        maxRepeatingChar = max (maxRepeatingChar ,freqMap[s[windowEnd]]);
        
        if((windowEnd - windowStart + 1 - maxRepeatingChar) > k ){
            char m = s[windowStart];
            freqMap[m]--;
            windowStart++;
        }
        maxLen = max(maxLen , windowEnd - windowStart + 1);
    }
    cout<<" lengthOfLongestSubstring after k replacements is "<<maxLen;

    return 0;
}