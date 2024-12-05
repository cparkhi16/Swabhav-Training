#include <iostream>
#include <vector>
#include <string>

using namespace std;

// Trie Node structure
struct TrieNode {
    TrieNode* children[26]; // Assuming lowercase English letters
    bool isEndOfWord;       // Marks the end of a word
    
    TrieNode() {
        for (int i = 0; i < 26; ++i) {
            children[i] = nullptr;
        }
        isEndOfWord = false;
    }
};

// Trie class
class Trie {
public:
    TrieNode* root;
    
    Trie() {
        root = new TrieNode();
    }

    // Insert a word into the trie
    void insert(const string& word) {
        TrieNode* node = root;
        int index;
        for (char ch : word) {
             index = ch - 'a';
            if (node->children[index] == nullptr) {
                node->children[index] = new TrieNode();
            }
            node = node->children[index];
        }
        cout<<" index is "<<index<<endl;
        node->isEndOfWord = true;
         //cout<<" index is "<<index<<node->ise
    }

    // Find the longest common prefix
    string longestCommonPrefix() {
        string prefix;
        TrieNode* node = root;
        
        while (node) {
            int childCount = 0;
            int nextIndex = -1;

            // Count the number of children and identify the next character
            int i =0;
            for (i = 0; i < 26; ++i) {
                if (node->children[i]) {
                    char res = i +'a';
                    cout<<" i is "<<i<<" char is "<<res<<"is end "<<node->children[i]->isEndOfWord<<endl;
                    ++childCount;
                    nextIndex = i;
                }
            }
            cout<<" i is "<<i<<endl;
            // If there's more than one child or we've reached the end of a word, stop
            if (childCount != 1 || node->isEndOfWord) {
                break;
            }

            // Add the next character to the prefix
            prefix += ('a' + nextIndex);
            node = node->children[nextIndex];
        }

        return prefix;
    }
};

// Function to find the longest common prefix using a Trie
string findLongestCommonPrefix(vector<string>& strs) {
    if (strs.empty()) {
        return "";
    }

    Trie trie;
    for (const string& word : strs) {
        trie.insert(word);
    }

    return trie.longestCommonPrefix();
}

// Main function
int main() {
    vector<string> strs = {"fl","fl"};
    
    string lcp = findLongestCommonPrefix(strs);
    cout << "Longest Common Prefix: " << lcp << endl;

    return 0;
}
