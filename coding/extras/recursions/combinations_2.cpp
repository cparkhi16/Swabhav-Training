// Online C++ compiler to run C++ program online
#include <iostream>
#include <bits/stdc++.h>
using namespace std;

 void findCombination(int ind, int target, vector<int> &arr, vector<vector<int>> &ans, vector<int>&ds) {
        if(target==0) {
            ans.push_back(ds);
            return;
        }        
        for(int i = ind;i<arr.size();i++) {
            if(i>ind && arr[i]==arr[i-1]) continue; 
            if(arr[i]>target) break; 
            ds.push_back(arr[i]);
            findCombination(i+1, target - arr[i], arr, ans, ds); 
            ds.pop_back(); 
        }
    }

int main() {
    vector<int> nums = {1,1,1,2,2};
    int target = 4 ;
    vector<vector<int>> res;
    vector<int> ds; 
    sort(nums.begin(), nums.end());
    findCombination( 0 , target , nums , res , ds);
    for(auto row : res){
        for(int i : row){
            cout<<" "<<i<<" ";
        }
        cout<<endl;
    }
}