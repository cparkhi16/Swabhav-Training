// Online C++ compiler to run C++ program online
#include <iostream>
#include<bits/stdc++.h>
using namespace std;




void PlaceNQueen(vector<vector<string>> &res ,vector<string> &board , int col,int n ,vector<int> &lower_diag, vector<int> &upper_diag, vector<int> &left_row){
    
    if( col == n){
        //cout<<" here ";
        res.push_back(board);
        return;
    }
    
    for(int row = 0; row < n ; row++){
       // cout<<" row "<<row<<"col "<<col<<endl;
        if( left_row[row] == 0 && lower_diag[row + col ]== 0 && upper_diag[ n-1 + col-row] == 0){
            board[row][col] = 'Q';
            left_row[row] = 1;
            lower_diag[row + col] = 1;
            upper_diag[n - 1 + col - row] = 1;
            PlaceNQueen(res , board , col + 1,n,lower_diag , upper_diag , left_row );
            board[row][col] = '.';
            left_row[row] = 0;
            lower_diag[row + col] = 0;
            upper_diag[n - 1 + col - row] = 0;
        }
    }
    
}

int main() {
    int n =4;
    vector<vector<string>> res;
    vector < string > board(n);
      string s(n, '.');
      for (int i = 0; i < n; i++) {
        board[i] = s;
      }
      vector<int> lower_diag (2 * n -1, 0);
      vector<int> upper_diag (2 * n -1, 0);
      vector<int> left_row (n, 0);
    PlaceNQueen(res , board ,0 , n , lower_diag , upper_diag , left_row );
    
  for (int i = 0; i < res.size(); i++) {
    cout << "Arrangement " << i + 1 << "\n";
    for (int j = 0; j < res[0].size(); j++) {
      cout << res[i][j];
      cout << endl;
    }
    cout << endl;
  }
    return 0;
}