// Online C++ compiler to run C++ program online
#include <iostream>
#include <bits/stdc++.h>
using namespace std;

void findPermu(int freq[], vector<int> nums , vector<int> &ds , vector<vector<int>> &res){

    if( ds.size() == nums.size() ){
        res.push_back(ds);
    }
    
    for(int i = 0 ; i < nums.size() ; i ++){
        if(!freq[i]){
            ds.push_back(nums[i]);
            freq[i] = 1;
            findPermu( freq , nums , ds , res );
            freq[i] = 0;
            ds.pop_back();
        }
    }
}


int main() {
    vector<int> nums = {1,2,3};
    vector<vector<int>> res;
    vector<int> ds;
    int freq[nums.size()];
    for(int i =0 ; i<nums.size() ; i++) freq[i] = 0;
    findPermu(freq ,nums , ds , res );
   // cout<<" res len "<< res.size();
    for(auto row : res){
    
        for(int i : row){
            cout<<" "<<i<<" ";
        }
        cout<<endl;
    }
    return 0;
}