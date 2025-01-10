// Online C++ compiler to run C++ program online
#include <iostream>
#include <bits/stdc++.h>
using namespace std;
//https://leetcode.com/problems/shortest-unsorted-continuous-subarray/description/
int shortestSubArrayLenToMakeArraySorted(vector<int> nums){
    int left = 0;
    int right = nums.size()-1;
    
    while(left<right && nums[left]<=nums[left+1]){
        left++;
    }
    
    if(left==right){
        return 0;
    }
    
    while(right>0 && nums[right]>=nums[right-1]){
        right--;
    }
    
    int mini= INT_MAX;
    int maxi = INT_MIN;
    int n = nums.size()-1;
    for(int k = left ; k <= right ; k++){
        mini = min(mini , nums[k]);
        maxi = max(maxi , nums[k]);
    }
   
    while(left>0 && nums[left-1] > mini){
        left--;
    }
    
    while(right<n && nums[right+1]<maxi){
        right++;
    }
    
    return right-left+1;
}
int main() {
    vector<int> nums = {1,3,2,2,2};
    int res = shortestSubArrayLenToMakeArraySorted(nums);
    cout<<" ans is " <<res;
    return 0;
}