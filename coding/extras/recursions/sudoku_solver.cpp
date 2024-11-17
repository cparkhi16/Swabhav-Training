// Online C++ compiler to run C++ program online
#include <iostream>
#include<bits/stdc++.h>
using namespace std;

bool isValid(vector<vector<char>> board , int row , int col, char k){
    
    for(int g = 0 ; g < 9; g++){
        
        if(board[row][g] == k){
            return false;
        }
        
        if(board[g][col] == k){
            return false;
        }
        
        if(board[3 * (row/3) + g /3] [3 * (col /3) + g % 3] == k){
            return false;
        }
    }
    return true;
}

bool solveSudoku(vector<vector<char>> &board){
    
    for(int i = 0; i < board.size() ; i++){
        for(int j = 0 ; j < board[0].size() ; j++){
            //cout<<" i  "<<i<<" j "<<j<<endl;
            if(board[i][j] == '.'){
                for(char k = '1' ;  k <='9'; k++){
                    if(isValid(board , i , j, k)){
                        board[i][j] = k;
                        
                        if(solveSudoku(board) == true){
                            return true;
                        }
                        else{
                            board[i][j] = '.';
                        }
                    }
                }
                return false;
            }
        }
    }
    
    return true;
}
int main() {
    vector<vector<char>>board{
        {'9', '5', '7', '.', '1', '3', '.', '8', '4'},
        {'4', '8', '3', '.', '5', '7', '1', '.', '6'},
        {'.', '1', '2', '.', '4', '9', '5', '3', '7'},
        {'1', '7', '.', '3', '.', '4', '9', '.', '2'},
        {'5', '.', '4', '9', '7', '.', '3', '6', '.'},
        {'3', '.', '9', '5', '.', '8', '7', '.', '1'},
        {'8', '4', '5', '7', '9', '.', '6', '1', '3'},
        {'.', '9', '1', '.', '3', '6', '.', '7', '5'},
        {'7', '.', '6', '1', '8', '5', '4', '.', '9'}
    };
   
    solveSudoku(board);
        	
    for(int i= 0; i< 9; i++){
        for(int j= 0; j< 9; j++)
            cout<<board[i][j]<<" ";
            cout<<"\n";
    }

    return 0;
}