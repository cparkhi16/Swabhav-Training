// Online C++ compiler to run C++ program online
#include <iostream>
#include <bits/stdc++.h>
using namespace std;

int getMaxLen(vector<int> nums ){
    
    unordered_map<int , int> prefixSum;
    int pSum = 0;
    prefixSum[0] = -1;
    int maxLen = 0;
    for(int i = 0 ; i < nums.size() ; i++){
        pSum += (nums[i] == 0) ? -1 : 1 ;
        
        if(prefixSum.find(pSum) != prefixSum.end()){
           maxLen = max(maxLen , i - prefixSum[pSum]);
        }else{
            prefixSum[pSum] = i;
        }
    }
    return maxLen;
}
int main() {
    vector<int> nums = {0,1,0,1,1,0 ,0,0};
    cout<<" Max size of subarray with an equal number of 0 and 1 are "<<getMaxLen(nums);

    return 0;
}