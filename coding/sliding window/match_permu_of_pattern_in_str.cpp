// Online C++ compiler to run C++ program online
#include <iostream>
#include <bits/stdc++.h>
using namespace std;
int main() {
    std::string str = "odicf";
    std::string pattern = "dc";
    int windowStart = 0;
    unordered_map<char , int> freqMap;
    bool isFound = false;
    int isMatch = 0;
    for(int i = 0 ; i < pattern.length() ; i++){
        freqMap[pattern[i]]++;
    }
    for(int windowEnd = 0 ; windowEnd < str.length() ; windowEnd++){
        char s = str[windowEnd];
        if( freqMap.find(s) != freqMap.end()){
            freqMap[s] --;
            if(freqMap[s] == 0){
                isMatch ++;
            }
        }
        cout<<" char processing "<<s<<" ismatch val "<<isMatch<<" freqMap size "<<freqMap.size()<<endl;
        if( isMatch == freqMap.size()){
             cout<<" given string has a permutation of pattern  ";
             isFound = true;
             break;
        }
        if( windowEnd >= pattern.length()-1){
            char m = str[windowStart];
            //freqMap[m]--;
            windowStart++;
            if( freqMap.find(m) != freqMap.end()){
                if(freqMap[m] == 0){
                    isMatch--;
                }
                freqMap[m]++;
            }
        }
    }
    if(isFound == false){
    cout<<" given string doesn't have a permutation of pattern  ";
    }
    return 0;
}