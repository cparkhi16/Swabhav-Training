// Online C++ compiler to run C++ program online
#include <iostream>
#include <bits/stdc++.h>

using namespace std;

void permute(vector<int> nums , int ind , vector<vector<int>> &res){
    
    if(ind == nums.size()){
        //cout<<" here "<<ds.size()<<endl;
        res.push_back(nums);
        return;
    }
    
    for(int i = ind ; i < nums.size() ; i++){
        swap(nums[ind] , nums[i]);
        permute( nums , ind + 1 , res);
        swap(nums[ind] , nums[i]);
    }
}

int main() {
    vector<int> nums = {1,2,3};
    vector<vector<int>> res;
    permute( nums , 0 , res);
    
    for(auto row : res){
        for(int i : row){
            cout<<" "<<i<<" ";
        }
        cout<<endl;
    }

    return 0;
}