// Online C++ compiler to run C++ program online
#include <iostream>
#include <bits/stdc++.h>
int main() {

    int windowStart = 0;
    std::vector<float> avg;
    int windowSum =0;
    std::vector arr = {1, 3, 2, 6, -1, 4, 1, 8, 2};
    int k = 5;
    for( int windowEnd = 0; windowEnd < arr.size() ; windowEnd ++){
        windowSum = windowSum + arr[windowEnd];
        if( windowEnd >= k-1){
            avg.push_back(static_cast<float>(windowSum) /k);
            windowSum = windowSum - arr[windowStart];
            windowStart ++;
        }
    }
    for ( int i =0 ; i < avg.size() ; i++){
        std::cout<<" avg vec is "<<avg[i];
    }
    return 0;
}