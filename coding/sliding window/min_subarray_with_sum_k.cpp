// Online C++ compiler to run C++ program online
#include <iostream>
#include <bits/stdc++.h>
int main() {
    // Write C++ code here

    int windowStart = 0;
    int minRange =INT_MAX;
    int windowSum = 0;
    std::vector arr = {2, 1, 5, 2, 3, 2};
    int s = 7;
    for( int windowEnd = 0; windowEnd < arr.size() ; windowEnd ++){
        windowSum = windowSum + arr[windowEnd];
        while( windowSum >= s){
            minRange = std::min(minRange,windowEnd - windowStart+1);
            windowSum = windowSum - arr[windowStart];
            windowStart ++;
        }
        
    }
        std::cout << "min subarrray len with given sum is "<<minRange;
    return 0;
}