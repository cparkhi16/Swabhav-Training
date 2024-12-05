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
void findAllWordsDFS(TrieNode* root , string prefix , vector<string>& result){
    if(root->isWordEnd){
        result.push_back(prefix);
    }
    for(int i = 0 ; i<26; i++){
        if(root->child[i]){
            char c = 'a' + i;
            findAllWordsDFS(root->child[i] , prefix + c , result);
        }
    }
    
}
void findAllWordsWithPrefix(TrieNode* root , string prefix , vector<string>& result){
    TrieNode* curr = root;
    
    for(char p: prefix){
        if(curr->child[p-'a'] == nullptr){
            return;
        }
        curr= curr->child[p-'a'];
    }
    findAllWordsDFS(curr , prefix ,result);
}

int main() {
    TrieNode* r = new TrieNode();
    vector<string> arr = {"car" , "cat" , "and" , "dog"};
    for(auto s : arr){
        insertKey(r, s);
    }
    vector<string> se = {"car" ,"dot"};
    for(auto s : se){
        cout<<"search string "<<s<<" "<<searchKey(r, s)<<endl;
    }
    vector<string> res;
    findAllWordsWithPrefix(r , "ca" ,res);
    for(auto s : res){
        cout<<" prefix with ca are "<<s<<endl;
    }
    // Write C++ code here
    //std::cout << "Try programiz.pro";
    
    return 0;
}