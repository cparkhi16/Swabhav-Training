// Online C++ compiler to run C++ program online
#include <iostream>
using namespace std;
#include <bits/stdc++.h>
int main() {
    std::vector s = {0, 1, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1, 1};
    int k = 3;
    int windowStart = 0;
    int maxRepeatingOnes = 0;
    int maxLen = 0;
    for(int windowEnd = 0 ; windowEnd < s.size() ; windowEnd++){
        if(s[windowEnd] == 1){
            maxRepeatingOnes++;
        }
        
        if((windowEnd - windowStart + 1 - maxRepeatingOnes) > k ){
            if(s[windowStart] == 1){
                maxRepeatingOnes--;
            }
            windowStart++;
        }
        maxLen = max(maxLen , windowEnd - windowStart + 1);
    }
    cout<<" contiguous subarray having all 1's after k replacements is "<<maxLen;

    return 0;
}