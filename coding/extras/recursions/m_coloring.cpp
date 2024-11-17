// Online C++ compiler to run C++ program online
#include <iostream>
#include<bits/stdc++.h>
using namespace std;

bool isValidColor(bool graph[101][101], int N , int color[], int colr ,int node){
    
   for(int k = 0 ; k < N; k ++){
       
       if( k != node && graph[k][node] == 1 && color[k] == colr){
           return false;
       }
   }
   return true;
}
bool solve(int node ,bool graph[101][101] , int m , int N, int color[] ){
    
    if(node == N){
        return true;
    }
    for(int colr = 1 ; colr <=m ; colr++){
        
        if(isValidColor(graph, N , color,colr , node)){
            color[node] = colr;
            if(solve(node + 1, graph, m , N, color) == true){
                return true;
            }
            color[node] = 0;
        }
    }
    return false;
    
}

bool graphColoring(bool graph[101][101] , int m , int N){
    int color[N] = {0};
    
    if( solve (0, graph , m , N, color))
    return true;
    else
    return false;
   
}
int main() {
  int N = 4;
  int m = 3;

  bool graph[101][101];
  memset(graph, false, sizeof graph);

  // Edges are (0, 1), (1, 2), (2, 3), (3, 0), (0, 2)
  graph[0][1] = 1; graph[1][0] = 1;
  graph[1][2] = 1; graph[2][1] = 1;
  graph[2][3] = 1; graph[3][2] = 1;
  graph[3][0] = 1; graph[0][3] = 1;
  graph[0][2] = 1; graph[2][0] = 1;
  
  cout << graphColoring(graph, m, N);

    return 0;
}