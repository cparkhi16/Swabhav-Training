// Online C++ compiler to run C++ program online
#include <iostream>
#include <bits/stdc++.h>
using namespace std;

int main() {
    int windowStart = 0;
    int matched = 0;
    int substrStart = 0;
    int minLen = INT_MAX;
    
    unordered_map <char,int> freqMap;
    
    string s = "aabdecaaaaa";
    string pattern = "dbc";
    
    for(int i =0 ; i< pattern.length() ; i++ ){
        freqMap[pattern[i]]++;
    }

    for(int windowEnd = 0 ; windowEnd <s.length() ; windowEnd ++ ){
        char c = s[windowEnd];
        if(freqMap.find(c) != freqMap.end()){
            freqMap[c] --;
            
            if(freqMap[c] == 0){
                matched ++;
            }
        }
        
        while(matched == freqMap.size()){
            if(minLen > windowEnd - windowStart + 1){
                minLen = windowEnd - windowStart + 1;
                substrStart = windowStart;
            }
            char sta = s[windowStart];
            windowStart++;
            
            if(freqMap.find(sta) != freqMap.end()){
                if(freqMap[sta] == 0){
                    matched --;
                }
                freqMap[sta]++;
            }
        }
        
    }
    if(minLen > s.length()){
        cout<<" no substr of pattern in any order present in og string ";
    }else{
        cout<<" substr of pattern in any order present in og string "<<s.substr(substrStart , minLen);
        cout<<" subs "<<substrStart<<endl;
        cout<<" end "<<minLen<<endl;
    }
    return 0;
}