// Online C++ compiler to run C++ program online
#include <iostream>
#include <bits/stdc++.h>
using namespace std;
int main() {
   string s = "acd";
   string t = "ced";
   int n = s.length();
   int m = t.length();
   vector<vector<int>> dp(n+1, vector<int> (m+1,0));
   for(int i = 0 ;i < n+1 ; i++) dp[i][0] = 0;
   for(int j = 0 ; j < m+1 ; j++ )dp[0][j] = 0;
   
   for(int i = 1 ; i < n+1 ; i++){
       for(int j = 1 ; j < m+1 ; j++){
           if(s[i-1]==t[j-1]){
               dp[i][j]= 1+ dp[i-1][j-1];
           }else{
               dp[i][j] = max(dp[i][j-1], dp[i-1][j]);
           }
       }
   }
   cout<<" lcs is "<<dp[n][m];
}