// Online C++ compiler to run C++ program online
#include <iostream>
#include<bits/stdc++.h>

using namespace std;
struct TrieNode{
    TrieNode* child[26];
    bool isWordEnd;
    
    TrieNode(){
        isWordEnd = false;
        
        for(int i = 0 ; i < 26 ; i++){
            child[i] = NULL;
        }
    }
};

void insertKey(TrieNode* root ,string key){
    if(!root) return;
    
    TrieNode* curr = root;
    
    for(char c :  key){
        
        if(curr->child[c - 'a'] == nullptr){
            TrieNode* r = new TrieNode();
            curr->child[c-'a'] = r;
        }
        curr = curr->child[c-'a'];
    }
    curr->isWordEnd = true;
}

bool searchKey(TrieNode* root ,string key){
    if(!root) return false;
    
    TrieNode* curr = root;
    
    for(char c :  key){
        
        if(curr->child[c - 'a'] == nullptr){
          return false;
        }
        curr = curr->child[c-'a'];
    }
    return curr->isWordEnd;
}
string findLCP(TrieNode* root){
    string res;
    
    while(root){
        int childCount = 0;
        int nextIndex = -1;
        for(int i = 0 ; i < 26 ; i++){
            if(root->child[i]){
                childCount++;
                nextIndex = i;
            }
        }
         
            if(childCount != 1 || root->isWordEnd){
                break;
            }
            
            res += ('a' + nextIndex);
            root = root->child[nextIndex];
    }
    return res;
}
int main() {
    TrieNode* r = new TrieNode();
    // vector<string> arr = {"car" , "cat" , "and" , "dog"};
    // for(auto s : arr){
    //     insertKey(r, s);
    // }
    // vector<string> se = {"car" ,"dot"};
    // for(auto s : se){
    //     cout<<"search string "<<s<<" "<<searchKey(r, s)<<endl;
    // }
    
    vector<string> arr = {"flower" , "flow"};
      for(auto s : arr){
        insertKey(r, s);
    }
    cout<<" LCP is "<<findLCP(r)<<endl;
    // Write C++ code here
    //std::cout << "Try programiz.pro";
    
    return 0;
}