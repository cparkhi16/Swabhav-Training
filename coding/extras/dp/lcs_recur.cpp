// Online C++ compiler to run C++ program online
#include <iostream>
#include <bits/stdc++.h>
using namespace std;

int findLCS(string s, string t,int i , int j){
    if(i<0 || j<0) return 0;
    
    if(s[i]==t[j]){
        cout<<"match"<<endl;
        return 1 + findLCS(s,t, i-1,j-1);
    }else{
        return max(findLCS(s,t,i,j-1), findLCS(s,t,i-1,j));
    }
}
int main() {
  string s ="acd";
  string t = "ced";
  cout<<" len of lcs is "<<findLCS(s,t, s.length()-1,t.length()-1);

    return 0;
}