// Online C++ compiler to run C++ program online
#include <iostream>
#include <bits/stdc++.h>
using namespace std;

void findPaths(int row , int col , string move , vector<string> &res, int n, vector<int>& di , vector<int>& dj, vector<vector<int>> &maze, vector<vector<int>>& vis ){
    
    if( row == n-1 && col == n-1 ){
        // cout<<" addinng move"<<move<<endl;
        res.push_back(move);
        return;
    }
    string dir = "DLRU";
    for(int i = 0 ; i < 4; i++){
        int nexti = row + di[i];
        int nextj = col + dj[i];
        //cout<<" nexti "<<nexti<< "row "<< row <<" col "<<col<<endl;
        //cout<<" nextj "<<nextj<< "row "<< row <<" col "<<col<<endl;
        if( nexti >= 0 && nextj >= 0 && nexti < n && nextj < n && vis[nexti][nextj]!= 1 && maze[nexti][nextj] == 1){
            vis[nexti][nextj] = 1;
            findPaths(nexti , nextj , move + dir[i] , res , n , di , dj , maze,vis);
            vis[nexti][nextj] = 0;
        }
    }
    
    
}


int main() {
   vector < vector < int >> maze = {{1,0,0,0},{1,1,0,1},{1,1,0,0},{0,1,1,1}};
    int n =4;
    vector<int> di = {1,0,0,-1};
    vector<int> dj = {0,-1,1,0};
    vector<string> res;
    vector<vector<int>> vis(n , vector<int> (n, 0));
    findPaths( 0 , 0 ,"",res ,n,di , dj ,maze ,vis);
    for( auto i : res){
        cout<<" "<<i<<endl;
    }
    
    return 0;
}