// Online C++ compiler to run C++ program online
#include <iostream>
#include <bits/stdc++.h>
using namespace std;

int main() {
    int windowStart = 0;
    int matched = 0;
    vector<int> indices;
    unordered_map <char,int> freqMap;
    
    string s = "ppqp";
    string pattern = "pq";
    
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
        if( matched == freqMap.size()){
            indices.push_back(windowStart);
        }
        
        if(windowEnd >= pattern.length() - 1){
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
    for(int k = 0 ; k < indices.size() ; k ++){
        cout <<" start index of found anagram in original string of given pattern "<<indices[k]<<endl;
    }
    return 0;
}