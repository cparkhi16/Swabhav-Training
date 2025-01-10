// Online C++ compiler to run C++ program online
#include <iostream>
#include <bits/stdc++.h>
using namespace std;

int getTotalSubarrays(vector<int> nums , int k){
    
    unordered_map<int , int> prefixSumCount;
    int prefixSum = 0;
    prefixSumCount[0] = 1;
    int count = 0;
    for(int i = 0 ; i < nums.size() ; i++){
        prefixSum += nums[i];
        
        if(prefixSumCount.find(prefixSum - k) != prefixSumCount.end()){
            count += prefixSumCount[prefixSum - k];
        }
        prefixSumCount[prefixSum]++;
    }
    return count;
}
int main() {
    vector<int> nums = {1,1,1};
    int k = 2;
    cout<<" total subarrays with sum = k are "<<getTotalSubarrays(nums,k);

    return 0;
}