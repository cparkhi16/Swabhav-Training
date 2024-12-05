// Online C++ compiler to run C++ program online
#include <iostream>
#include<bits/stdc++.h>
using namespace std;

int findLongestSubs(string s){
    int windowStart = 0;
    int maxLen = INT_MIN;
    unordered_map<char,int> occMap;
    
    for(int windowEnd = 0 ; windowEnd < s.length() ; windowEnd++){
    
        if(occMap.find(s[windowEnd]) != occMap.end()){
            
            windowStart = max(windowStart , occMap[s[windowEnd]] +1 );
        }
        
        occMap[s[windowEnd]] =windowEnd;
        maxLen = max(maxLen , windowEnd - windowStart + 1);
    }
    return maxLen;
}
int main() {
    // Write C++ code here
    //std::cout << "Try programiz.pro";
    int l = findLongestSubs("abccde");
    cout<<" ans is "<<l;
    return 0;
}