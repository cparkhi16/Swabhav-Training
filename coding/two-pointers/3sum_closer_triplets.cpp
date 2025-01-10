// Online C++ compiler to run C++ program online
#include <iostream>
#include <bits/stdc++.h>
using namespace std;

vector<vector<int>> threeSumClosest(vector<int> nums , int target){
    sort(nums.begin(), nums.end());
    int triplets = 0;
    vector<vector<int>> res;
    //cout<<" nums size "<<nums.size()-2<<endl;
    if(nums.size()<3){
        return res;
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
                for(int j = right ; j > left; j--){
                    res.push_back({nums[i], nums[left], nums[j]});
                }
                left++;
            }else{
                right--;
            }
        }
       
        
    }
    cout<<" total triplets "<<triplets<<endl;
     return res;
}


int main() {
    
    vector<int> nums = {-1, 0, 2, 3};
    int k = 3;
    vector<vector<int>> res = threeSumClosest(nums ,k);
    for(auto a : res){
        for(int i : a){
            cout<<i<<" ";
        }
        cout<<endl;
    }
    return 0;
}