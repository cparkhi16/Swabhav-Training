// Online C++ compiler to run C++ program online
#include <iostream>
#include<bits/stdc++.h>
using namespace std;

void findSubsetSum( int ind , vector<int>& nums , int sum , vector<int>& ans){
    if( ind == nums.size()){
        ans.push_back(sum);
        return;
    }
    findSubsetSum( ind+1 , nums , sum + nums[ind] , ans);
    findSubsetSum( ind+1 , nums , sum  , ans);
}
int main() {
    vector<int> nums = {3,2,1};
    
    vector<int> ans;
    int sum = 0;
    
    findSubsetSum(0 ,nums  , sum  , ans);
    sort(ans.begin(), ans.end());
    for(int i: ans){
        cout<<" "<<i<<" ";
    }

    return 0;
}