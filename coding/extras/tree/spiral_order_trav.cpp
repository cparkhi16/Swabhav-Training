#include <iostream>
#include <vector>
#include <stack>
using namespace std;

// Binary tree node structure
struct Node {
    int data;
    Node* left;
    Node* right;

    Node(int x) {
        data = x;
        left = right = nullptr;
    }
};

// Function to return a list containing the level order traversal in spiral form.
vector<int> findSpiral(Node* root) {
    vector<int> result;
    if (root == nullptr) return result;
    if (root->left == nullptr && root->right == nullptr) {
        result.push_back(root->data);
        return result;
    }

    stack<Node*> s1;
    stack<Node*> s2;

    s1.push(root);
    while (!s1.empty() || !s2.empty()) {
        while (!s1.empty()) {
            Node* temp = s1.top();
            s1.pop();
            result.push_back(temp->data);

            if (temp->right != nullptr) s2.push(temp->right);
            if (temp->left != nullptr) s2.push(temp->left);
        }
        while (!s2.empty()) {
            Node* temp = s2.top();
            s2.pop();
            result.push_back(temp->data);

            if (temp->left != nullptr) s1.push(temp->left);
            if (temp->right != nullptr) s1.push(temp->right);
        }
    }
    return result;
}

// Utility to build a sample binary tree for testing
Node* buildSampleTree() {
    // Sample Tree:
    //        1
    //      /   \
    //     2     3
    //    / \   / \
    //   7   6 5   4

    Node* root = new Node(1);
    root->left = new Node(2);
    root->right = new Node(3);
    root->left->left = new Node(7);
    root->left->right = new Node(6);
    root->right->left = new Node(5);
    root->right->right = new Node(4);
    return root;
}

// Driver Code
int main() {
    Node* root = buildSampleTree();

    vector<int> spiralOrder = findSpiral(root);

    cout << "Spiral Level Order Traversal: ";
    for (int val : spiralOrder) {
        cout << val << " ";
    }
    cout << endl;

    return 0;
}
