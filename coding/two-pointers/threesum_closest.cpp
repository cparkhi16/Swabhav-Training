// Online C++ compiler to run C++ program online
#include <iostream>
#include <bits/stdc++.h>
using namespace std;

int threeSumClosest(vector<int>& nums, int target) {
    if(nums.size() < 3){
        return 0;
    }
    sort(nums.begin(), nums.end());
    int closestSum = nums[0]+nums[1]+nums[2];
    for(int i = 0 ; i < nums.size()-2; i++){
        int left = i+1;
        int right = nums.size()-1;
        
        while(left<right){
            int curr = nums[i] + nums[left]+ nums[right];
            
            if(abs(curr-target) < abs(closestSum-target)){
                closestSum = curr;
            }
            if(curr < target){
                left++;
            }else{
                right--;
            }
        }
       
        
    }
     return closestSum;
}


int main() {
    
    vector<int> nums = {1,1,1,0};
    int k = 100;
    cout<<" Triplets sum closer to the target is "<<threeSumClosest(nums,k);
    return 0;
}