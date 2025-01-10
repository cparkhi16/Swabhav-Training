// Online C++ compiler to run C++ program online
#include <iostream>
#include <bits/stdc++.h>
using namespace std;

int threeSumClosest(vector<int> nums , int target){
    sort(nums.begin(), nums.end());
    int triplets = 0;
    //cout<<" nums size "<<nums.size()-2<<endl;
    if(nums.size()<3){
        return 0;
    }
    for(int i = 0 ; i < nums.size()-2; i++){
        int left = i+1;
        int right = nums.size()-1;
        
        while(left<right){
            int curr = nums[i] + nums[left]+ nums[right];
            
            // if(abs(curr-target) < abs(curr-closestSum)){
            //     closestSum = curr;
            // }
            if(curr < target){
                triplets+= (right-left);
                left++;
            }else{
                right--;
            }
        }
       
        
    }
     return triplets;
}


int main() {
    
    vector<int> nums = {-1, 0, 2, 3};
    int k = 3;
    cout<<" total triplets whose sum is closer to the target is "<<threeSumClosest(nums,k);
    return 0;
}