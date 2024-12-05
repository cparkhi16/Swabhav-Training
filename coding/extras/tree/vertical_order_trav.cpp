#include <bits/stdc++.h>
using namespace std;

struct Node {
    int key;
    Node *left, *right;
    Node(int k) {
        key = k;
        left = nullptr;
        right = nullptr;
    }
};

// Utility function to store vertical order
// in map 'm'. 'hd' is horizontal distance
// of current node from root. 'hd' is
// initially passed as 0
void getOrder(Node* root, int hd,
              map<int, vector<int>>& m) {
    // Base case
    if (root == nullptr)
        return;

    // Store current node in map 'm'
    m[hd].push_back(root->key);

    // Store nodes in left subtree
    getOrder(root->left, hd - 1, m);

    // Store nodes in right subtree
    getOrder(root->right, hd + 1, m);
}

// Main function to print vertical order of
// a binary tree with the given root
void printOrder(Node* root) {
  
    // Create a map and store vertical order
    // in map using function getOrder()
    map<int, vector<int>> m;
    int hd = 0;
  
    getOrder(root, hd, m);

    // Traverse the map and print nodes at
    // every horizontal distance (hd)
    for (auto& entry : m) {
        for (int key : entry.second) {
            cout << key << " ";
        }
        cout << endl;
    }
}

// Driver code
int main() {
    Node* root = new Node(1);
    root->left = new Node(2);
    root->right = new Node(3);
    root->left->left = new Node(4);
    root->left->right = new Node(5);
    root->right->left = new Node(6);
    root->right->right = new Node(7);
    root->right->left->right = new Node(8);
    root->right->right->right = new Node(9);
    printOrder(root);
    return 0;
}