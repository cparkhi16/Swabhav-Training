// Online C++ compiler to run C++ program online
#include <iostream>
#include <bits/stdc++.h>
using namespace std;

int main() {
    std::vector <string> words = {"cat" , "fox"};
    std::string s = "catfoxcat";
    std::vector<int> indices;
    unordered_map <string,int> freqMap;
  
    
    for(int i =0 ; i< words.size() ; i++ ){
        freqMap[words[i]]++;
    }
    // for (const auto& pair : freqMap) {
    //     std::cout << pair.first << ": " << pair.second << std::endl;
    // }
    int wordsCount = words.size();
    int wordLen = words[0].length();
    cout<<" boundary is "<<s.length() - wordsCount * wordLen<<endl;
    for(int i = 0 ; i < (s.length() - wordsCount * wordLen ) + 1 ; i ++){
        cout<<" i is "<<i<<endl;
        unordered_map <string,int> wordsSeen;
        for( int j = 0 ; j < wordsCount ; j++){
            int nextWordIndex = i + j * wordLen ;
            string subs = s.substr(nextWordIndex , wordLen );
            cout<<" word gen is "<<subs<<endl;
            if( freqMap.find(subs) == freqMap.end()){
                break;
            }
            
            wordsSeen[subs]++;
            
            if(wordsSeen[subs] > freqMap[subs]){
                break;
            }
            if(j + 1 == wordsCount){
                cout<< "pushing i "<<i<<endl;
                indices.push_back(i);
            }
            
        }
    }
    for(int j = 0 ; j < indices.size() ; j++){
        cout<<" index at which the given words combo match in og string is "<<indices[j]<<endl;
    }
    
    return 0;
}