// C++ program to Merge Two Binary Trees 
// using iteration

#include <bits/stdc++.h>
using namespace std;

class Node {
  public:
    int data;
    Node *left, *right;
    Node(int x) {
        data = x;
        left = right = nullptr;
    }
};

// Function to perform an inorder traversal 
// of the binary tree
void inorder(Node *node) {
  
    // Base case: if the node is null, return
    if (node == nullptr)
        return;

    // Recursively traverse the left subtree
    inorder(node->left);

    // Print the current node's data
    cout<< node->data << " ";

    // Recursively traverse the right subtree
    inorder(node->right);
}

// Function to merge two binary trees iteratively
// using pair<Node*, Node*>
Node *mergeTrees(Node *t1, Node *t2) {
  
    // If the first tree is null, return the second tree
    if (t1 == nullptr)
        return t2;

    // If the second tree is null, return the first tree
    if (t2 == nullptr)
        return t1;

    // Stack to store pairs of nodes to be processed
    stack<pair<Node *, Node *>> s;

    // Initialize the stack with the root
      // nodes of both trees
    s.push({t1, t2});

    while (!s.empty()) {
      
        // Get the top pair of nodes from the stack
        pair<Node *, Node *> n = s.top();
        s.pop();

        // If either node is null, skip this pair
        if (n.first == nullptr || n.second == nullptr)
            continue;

        // Add the data of the second tree's node to 
          // the first tree's node
        n.first->data += n.second->data;

        // Process the left children
        if (n.first->left == nullptr)
            n.first->left = n.second->left;
        else {
            s.push({n.first->left, n.second->left});
        }

        // Process the right children
        if (n.first->right == nullptr)
            n.first->right = n.second->right;
        else {
            s.push({n.first->right, n.second->right});
        }
    }

    // Return the root of the merged tree
    return t1;
}

int main() {
  
    // Construct the first Binary Tree
    //        1
    //      /   \
    //     2     3
    //    / \     \
    //   4   5     6

    Node *root1 = new Node(1);
    root1->left = new Node(2);
    root1->right = new Node(3);
    root1->left->left = new Node(4);
    root1->left->right = new Node(5);
    root1->right->right = new Node(6);

    // Construct the second Binary Tree
    //      4
    //    /   \
    //   1     7
    //  /     /  \
    // 3     2    6

    Node *root2 = new Node(4);
    root2->left = new Node(1);
    root2->right = new Node(7);
    root2->left->left = new Node(3);
    root2->right->left = new Node(2);
    root2->right->right = new Node(6);

    Node *root = mergeTrees(root1, root2);
    inorder(root);

    return 0;
}
