// Online C++ compiler to run C++ program online
#include <iostream>
#include <bits/stdc++.h>
using namespace std;
int main() {
   string s = "bbabcbcab";
   string t = s;
   reverse(t.begin(), t.end());
   cout<<" t is "<<t<<endl;
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
   int len = dp[n][m];
   cout<<"len of lcs is "<<dp[n][m]<<endl;
 
  string res;
  for(int i = 0 ; i< len ; i++){
      res+='$';
  }
  int i = n;
  int j = m;
  int index = len-1;
  while(i >0 && j >0){
       if(s[i-1] == t[j-1]){
           res[index] = s[i-1];
           index--;
           i--;
           j--;
       }else if (dp[i-1][j] > dp[i][j-1]){
           i=i-1;
       }else{
           j = j-1;
       }
  }
  cout<<"lcs is  "<<res<<endl;
}