// Online C++ compiler to run C++ program online
#include <iostream>
#include <bits/stdc++.h>
int main() {
    // Write C++ code here

    int windowStart = 0;
    std::unordered_map<char, int> freqMap;
    std::string str = "cbbebi";
    int maxRange =0;
    int k = 3;
    for( int windowEnd = 0; windowEnd < str.length() ; windowEnd ++){
        char s = str[windowEnd];
        freqMap[s]++;
        while( freqMap.size() > k){
            char e = str[windowStart];
            freqMap[e]--;
            if(freqMap[e] == 0){
                freqMap.erase(e);
            }
            windowStart ++;
        }
        maxRange = std::max(maxRange,windowEnd - windowStart+1);
    }
        std::cout << "max subarrray len with k distinct char  is "<<maxRange;
    return 0;
}