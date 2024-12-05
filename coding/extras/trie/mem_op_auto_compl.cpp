// A memory optimized CPP implementation of trie
// using unordered_map
#include <iostream>
#include <unordered_map>
using namespace std;
#include<bits/stdc++.h>
struct Trie {

	// isEndOfWord is true if the node
	// represents end of a word
	bool isEndOfWord;

	/* nodes store a map to child node */
	unordered_map<char, Trie*> map;
};

/*function to make a new trie*/
Trie* getNewTrieNode()
{
	Trie* node = new Trie;
	node->isEndOfWord = false;
	return node;
}

/*function to insert in trie*/
void insert(Trie*& root, const string& str)
{
	if (root == nullptr)
		root = getNewTrieNode();

	Trie* temp = root;
	for (int i = 0; i < str.length(); i++) {
		char x = str[i];

		/* make a new node if there is no path */
		if (temp->map.find(x) == temp->map.end())
			temp->map[x] = getNewTrieNode();

		temp = temp->map[x];
	}

	temp->isEndOfWord = true;
}

/*function to search in trie*/
bool search(Trie* root, const string& str)
{
	/*return false if Trie is empty*/
	if (root == nullptr)
		return false;

	Trie* temp = root;
	for (int i = 0; i < str.length(); i++) {

		/* go to next node*/
// 		temp = temp->map[str[i]];

// 		if (temp == nullptr)
// 			return false;
    if(temp->map.find(str[i]) == temp->map.end()){
        return false;
    }
    temp = temp->map[str[i]];
	}

	return temp->isEndOfWord;
}
void findAllWordsDFS(Trie* root , string prefix , vector<string>& result){
    if(root->isEndOfWord){
        result.push_back(prefix);
    }
    
    for(auto m: root->map){
        findAllWordsDFS(m.second , prefix+m.first ,result);
    }
    
}
void findAllWordsWithPrefix(Trie* root , string prefix , vector<string>& result){
    Trie* curr = root;
    
    for(char p: prefix){
        if(curr->map.find(p)== curr->map.end()){
            return;
        }
        curr= curr->map[p];
    }
    findAllWordsDFS(curr , prefix ,result);
}
/*Driver function*/
int main()
{
	Trie* root = nullptr;

	insert(root, "geeks");
	cout << search(root, "geeks") << " ";

	insert(root, "for");
	cout << search(root, "for") << " ";

	cout << search(root, "geekk") << " ";

	insert(root, "gee");
	cout << search(root, "gee") << " ";

	insert(root, "science");
	insert(root, "scie");
	cout << search(root, "science") << endl;
     vector<string> res;
    findAllWordsWithPrefix(root , "sc" ,res);
    for(auto s : res){
        cout<<" prefix with ca are "<<s<<endl;
    }
	return 0;
}
