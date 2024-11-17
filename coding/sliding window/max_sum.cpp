// Online C++ compiler to run C++ program online
#include <iostream>
#include <bits/stdc++.h>
int main() {
    // Write C++ code here

    int windowStart = 0;
    int maxSum =0;
    int windowSum = 0;
    std::vector arr = {2, 3, 4, 1, 5};
    int k = 2;
    for( int windowEnd = 0; windowEnd < arr.size() ; windowEnd ++){
        windowSum = windowSum + arr[windowEnd];
        if( windowEnd >= k-1){
            maxSum = std::max(windowSum,maxSum);
            windowSum = windowSum - arr[windowStart];
            windowStart ++;
        }
    }
        std::cout << "max sum is "<<maxSum;
    return 0;
}