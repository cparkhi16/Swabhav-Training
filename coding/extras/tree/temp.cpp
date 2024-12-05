#include<bits/stdc++.h>
using namespace std;

struct Tree{
    int data;
    Tree* left;
    Tree* right;

    Tree(int d): data(d), left(nullptr), right(nullptr){}
};

void traverse(Tree* root, int hd, map<int , vector<int>>& vertical_trav_map){
    if(!root) return ;

    vertical_trav_map[hd].push_back(root->data);

    traverse(root->left , hd-1 , vertical_trav_map);

    traverse(root->right, hd+1 , vertical_trav_map);
}

void printTraversal(Tree* root){
    map<int,vector<int>> tree_trav;

    traverse(root , 0 , tree_trav);

    for(auto d : tree_trav){
        for(int val : d.second){
            cout<<" "<<val<<" ";
        }
        cout<<endl;
    }
}

int main(){

    Tree* root = new Tree(1);
    root->left = new Tree(2);
    root->right = new Tree(3);
    root->left->left = new Tree(4);
    root->left->right = new Tree(5);
    root->right->left = new Tree(6);
    root->right->right = new Tree(7);
    root->right->left->right = new Tree(8);
    root->right->right->right = new Tree(9);
    printTraversal(root);
}