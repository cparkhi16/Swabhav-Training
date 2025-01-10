// Online C++ compiler to run C++ program online
#include <iostream>
#include <bits/stdc++.h>
using namespace std;

int shortestSubArrayLenToRemove(vector<int> nums){
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
    cout<<" left is "<<left<<" right is "<<right<<endl;
    int n = nums.size();
    int res = min(n-left-1, right);
    
    int i = 0 , j = right;
    
    while(i<=left && j<n){
        if(nums[i] <= nums[j]){
            res = min(res , j-i-1);// we need to merge prefix with suffix if nums[i] <= nums[j] , the subarray to be removed will be from i to j (excluding i & j , i+1 to j-1) hence res will be min of res and j-i-1
            i++;
            }
        else{
            j++;
        }
    }
    return res;
}
int main() {
    vector<int> nums = {2,2,2,1,1,1};
    int res = shortestSubArrayLenToRemove(nums);
    cout<<" shortestSubArrayLenToRemove is " <<res;
    return 0;
}