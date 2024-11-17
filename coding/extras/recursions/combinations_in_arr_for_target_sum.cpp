// Online C++ compiler to run C++ program online
#include <iostream>
#include <bits/stdc++.h>
using namespace std;

void findCombinations(vector<int> nums , vector<vector<int>> &res , int index , int target , vector<int>& ds){
    if( index == nums.size()){
        if(target == 0){
            //cout<<" here ";
            res.push_back(ds);
        }
        return;
    }
    
    if( nums[index] <= target){
        ds.push_back(nums[index]);
        // for(auto d : ds){
        //     cout<<" elems in ds "<<d<<" ";
        // }
        cout<<endl;
        findCombinations(nums , res , index , target- nums[index] , ds);
        ds.pop_back();
    }
    findCombinations(nums , res , index+1 , target , ds);
}

int main() {
    vector<int> nums = {2,3,6,7};
    int target = 7 ;
    vector<vector<int>> res;
    vector<int> ds; 
    findCombinations( nums , res , 0 , target , ds);
    //cout<<" res size "<<res.size();
    for(auto row : res){
        for(int i : row){
            cout<<" "<<i<<" ";
        }
        cout<<endl;
    }
}