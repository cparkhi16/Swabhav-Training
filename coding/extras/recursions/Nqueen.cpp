// Online C++ compiler to run C++ program online
#include <iostream>
#include<bits/stdc++.h>
using namespace std;

bool checkBoard( int row , int col , vector<string> board, int n){
    //cout<<"start checkBoard row "<<row <<" col "<<col<<endl;
    int tempR = row;
    int tempC = col;
    // left row
    while( tempC >=0 ){
        //cout<<"start 1st checkBoard row "<<row <<" col "<<col<<endl;
        if(board[row][tempC] == 'Q' ){
            return false;
        }
        tempC--;
    }
    
    tempC = col;
    // lower diagonal
    while( tempC>=0 && tempR <n){
        //cout<<"start 2nd checkBoard row "<<row <<" col "<<col<<endl;
        if(board[tempR][tempC] == 'Q'){
            return false;
        }
        tempC--;
        tempR++;
    }
    
    tempR= row;
    tempC = col;
    // upper diagonal
    while( tempR >= 0 && tempC>=0){
        //cout<<"start 3rd checkBoard row "<<row <<" col "<<col<<endl;
          if(board[tempR][tempC] == 'Q'){
            return false;
        }
        tempC--;
        tempR--;
    }
    //cout<<"checkBoard row "<<row <<" col "<<col<<endl;
    return true;
    
}


void PlaceNQueen(vector<vector<string>> &res ,vector<string> &board , int col,int n  ){
    
    if( col == n){
        //cout<<" here ";
        res.push_back(board);
        return;
    }
    
    for(int row = 0; row < n ; row++){
       // cout<<" row "<<row<<"col "<<col<<endl;
        if( checkBoard(row , col , board ,n )){
            board[row][col] = 'Q';
            PlaceNQueen(res , board , col + 1,n);
            board[row][col] = '.';
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
    PlaceNQueen(res , board ,0 , n);
    
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