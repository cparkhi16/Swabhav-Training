// Online C++ compiler to run C++ program online
#include <iostream>
#include <bits/stdc++.h>
using namespace std;

int f(int ind,vector<int>& heights,vector<int>& dp){
    
    if(ind==0){
        return 0;
    }
   if(dp[ind] !=-1) return dp[ind];
    int l1 = f(ind-1,heights,dp) + abs(heights[ind]- heights[ind-1]);
    int l2 = INT_MAX;
    if(ind > 1){
     l2 = f(ind-2,heights,dp) + abs(heights[ind] - heights[ind-2]);
    }
    
    return dp[ind]= min(l1,l2);
    
}
int main() {
    vector<int> heights = {10,20,30,10};

    //vector<int> dp(heights.size(),-1);
    //cout<<" minimum energy reqd by the frog "<<f(heights.size()-1, heights,dp);
    
    int n = heights.size();
    vector<int> dp(n,0);
    int prev = 0;
    int prev2 = 0;
    //dp[0] = 0;
    int curi = 0;
    for(int i = 1; i < n; i++ ){
        // int l1 = dp[i-1] + abs(heights[i]- heights[i-1]);
        int l1 = prev + abs(heights[i]- heights[i-1]);
        int l2 = INT_MAX;
        if(i > 1){
            // l2 = dp[i-2] + abs(heights[i] - heights[i-2]);
            l2 = prev2 + abs(heights[i] - heights[i-2]);
        }
    
        curi= min(l1,l2);
        prev2 = prev;
        prev = curi;
    }
    // cout<<" minimum energy reqd by the frog "<<dp[n-1]<<endl;
    cout<<" minimum energy reqd by the frog "<<prev<<endl;
    return 0;
}