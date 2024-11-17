// Online C++ compiler to run C++ program online
#include <iostream>
#include <bits/stdc++.h>
using namespace std;

struct Tree{
    int data;
    Tree* left;
    Tree* right;
    
    Tree(int d): data(d), left(nullptr),right(nullptr){}
};
Tree* findLCA(Tree* root ,int n1 , int n2 ){
    if(root == nullptr){
        return nullptr;
    }
    if(root->data == n1 || root->data == n2){
        return root;
    }
    
    Tree* left_lca = findLCA( root->left , n1 , n2);
    Tree* right_lca = findLCA(root->right , n1 , n2);
    
    if( left_lca && right_lca ){
        return root;
    }
    
    return left_lca? left_lca : right_lca;
}
int main() {
   Tree* root = new Tree(1);
    root->left = new Tree(2);
    root->right = new Tree(3);
    root->left->left = new Tree(4);
    root->left->right = new Tree(5);
    root->right->left = new Tree(6);
    root->right->right = new Tree(7);

    cout << "LCA(4, 5) = " << findLCA(root, 4, 5)->data;
    cout << "\nLCA(4, 6) = " << findLCA(root, 4, 6)->data;
    cout << "\nLCA(3, 4) = " << findLCA(root, 3, 4)->data;
    cout << "\nLCA(2, 4) = " << findLCA(root, 2, 4)->data;

    return 0;
}