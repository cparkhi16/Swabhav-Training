// Online C++ compiler to run C++ program online
#include <iostream>
#include <bits/stdc++.h>
using namespace std;
int main() {
    // Write C++ code here
    //std::cout << "Try programiz.pro";
    std::vector<int> fruits = {1,2,3,2,2}; // ans is 4
    int windowStart = 0;
    int maxFruits = 0;
    unordered_map<int,int> freqMap;
    for(int windowEnd = 0 ; windowEnd < fruits.size() ; windowEnd++){
      freqMap[fruits[windowEnd]]++;
      while (freqMap.size() > 2){
          int f = fruits[windowStart];
          freqMap[f] --;
          if (freqMap[f] == 0){
              freqMap.erase(f);
          }
          windowStart++;
      }
      maxFruits = std::max(maxFruits , windowEnd - windowStart + 1);
    }
    std::cout << "max fruits taken are "<<maxFruits;
    return 0;
}