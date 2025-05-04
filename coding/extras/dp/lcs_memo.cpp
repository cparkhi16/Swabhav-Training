// Online C++ compiler to run C++ program online
#include <iostream>
#include <bits/stdc++.h>
using namespace std;

int findLCS(string s, string t,int i , int j,vector<vector<int>> &dp){
    if(i<0 || j<0) return 0;
    if(dp[i][j] !=-1 ) return dp[i][j];
    if(s[i]==t[j]){
        //cout<<"match"<<endl;
        return dp[i][j] = 1 + findLCS(s,t, i-1,j-1,dp);
    }else{
        return dp[i][j]= max(findLCS(s,t,i,j-1,dp), findLCS(s,t,i-1,j,dp));
    }
}
int main() {
  string s ="acd";
  string t = "ced";
  int n =  s.length();
  int m = t.length();
  vector<vector<int>> dp(n, vector<int> (m,-1));
  cout<<" len of lcs is "<<findLCS(s,t, s.length()-1,t.length()-1,dp);

    return 0;
}