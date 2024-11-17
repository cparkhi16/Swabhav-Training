// Online C++ compiler to run C++ program online
#include <iostream>
#include <bits/stdc++.h>
using namespace std;

void findUniqueSubsequences(int ind , vector<int> nums , 
vector<vector<int>> &res , vector<int> &temp){
    
    res.push_back(temp);
    
    for(int i = ind ; i < nums.size() ; i ++){
        
        if(i != ind && nums[i] == nums[i-1]) continue;
        
        temp.push_back(nums[i]);
        findUniqueSubsequences( i+ 1, nums , res , temp);
        temp.pop_back();
    }
}


int main() {
    vector<int> nums = {1,2,2,2,3};
    sort(nums.begin(), nums.end());
    vector<vector<int>> res;
    vector<int> temp;
    findUniqueSubsequences(0 ,nums , res , temp );
   // cout<<" res len "<< res.size();
    for(auto row : res){
        if(row.size() == 0){
            cout<<" {} ";
        }
        for(int i : row){
            cout<<" "<<i<<" ";
        }
        cout<<endl;
    }
    return 0;
}