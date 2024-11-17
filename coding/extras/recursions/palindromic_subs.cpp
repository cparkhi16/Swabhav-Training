// Online C++ compiler to run C++ program online
#include <iostream>
#include <bits/stdc++.h>
using namespace std;

bool isPalindrome( string s , int beg , int end){
    while(beg <= end){
        if(s[beg++] != s[end--]){
            return false;
        }
    }
    return true;
}

void findPalindromicSubs(int ind , string s, vector<string>& path, vector<vector<string>> &res){
    
    if( ind == s.size()){
        res.push_back(path);
        return;
    }
    for(int i = ind ; i < s.size() ; i++){
        if(isPalindrome( s , ind , i )){
            path.push_back(s.substr(ind , i - ind + 1));
            findPalindromicSubs(i+1 , s , path , res);
            path.pop_back();
        }
    }
}

int main() {
    string s = "aabb";
    vector<string> path;
    vector<vector<string>> res;
    findPalindromicSubs( 0 , s , path , res);
    for(auto row : res){
        for(string i : row){
            cout<<" "<<i<<" ";
        }
        cout<<endl;
    }

    return 0;
}