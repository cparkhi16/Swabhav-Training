#include <iostream>
#include <unordered_map>
#include <vector>
#include <string>

class TrieNode {
public:
    std::unordered_map<char, TrieNode*> children;
    bool is_end_of_word;

    TrieNode() : is_end_of_word(false) {}
};

class Trie {
private:
    TrieNode* root;

    // Helper function for DFS to find suggestions
    void dfs(TrieNode* node, const std::string& prefix, std::vector<std::string>& suggestions) {
        if (node->is_end_of_word) {
            suggestions.push_back(prefix);
        }
        for (const auto& pair : node->children) {
            dfs(pair.second, prefix + pair.first, suggestions);
        }
    }

public:
    Trie() {
        root = new TrieNode();
    }

    // Insert a word into the trie
    void insert(const std::string& word) {
        TrieNode* node = root;
        for (char ch : word) {
            if (!node->children.count(ch)) {
                node->children[ch] = new TrieNode();
            }
            node = node->children[ch];
        }
        node->is_end_of_word = true;
    }

    // Check if a word exists in the trie
    bool search(const std::string& word) {
        TrieNode* node = root;
        for (char ch : word) {
            if (!node->children.count(ch)) {
                return false;
            }
            node = node->children[ch];
        }
        return node->is_end_of_word;
    }

    // Get suggestions for a prefix
    std::vector<std::string> getSuggestions(const std::string& prefix) {
        std::vector<std::string> suggestions;
        TrieNode* node = root;

        // Navigate to the end of the prefix
        for (char ch : prefix) {
            if (!node->children.count(ch)) {
                return suggestions; // No suggestions if prefix not found
            }
            node = node->children[ch];
        }

        // Perform DFS to find all words with the prefix
        dfs(node, prefix, suggestions);
        return suggestions;
    }
};

int main() {
    Trie trie;

    // Insert words into the trie
    std::vector<std::string> dictionary = {
        "apple", "orange", "banana", "grape", "grapefruit", "pineapple", "gravy"
    };
    for (const std::string& word : dictionary) {
        trie.insert(word);
    }

    // Take input word from the user
    std::string input_word;
    std::cout << "Enter a word to check spelling: ";
    std::cin >> input_word;

    // Check for the word and get suggestions
    if (trie.search(input_word)) {
        std::cout << "'" << input_word << "' is correctly spelled." << std::endl;
    } else {
        std::vector<std::string> suggestions = trie.getSuggestions(input_word);
        if (!suggestions.empty()) {
            std::cout << "'" << input_word << "' is misspelled. Did you mean:" << std::endl;
            for (const std::string& suggestion : suggestions) {
                std::cout << "- " << suggestion << std::endl;
            }
        } else {
            std::cout << "No suggestions found for '" << input_word << "'." << std::endl;
        }
    }

    return 0;
}
