#include <iostream>
#include <unordered_map>
using namespace std;

// Definition for a binary tree node
struct TreeNode {
    int val;
    TreeNode* left;
    TreeNode* right;
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
};

// Helper function to build the tree
TreeNode* buildTreeHelper(int inorder[], int postorder[], int& postIndex, 
                          int inStart, int inEnd, unordered_map<int, int>& inMap) {
    if (inStart > inEnd) return nullptr;

    // Get the root value from postorder
    int rootVal = postorder[postIndex--];
    TreeNode* root = new TreeNode(rootVal);

    // Find the root index in inorder
    int inIndex = inMap[rootVal];

    // Build right subtree first (postorder processes Right before Left)
    root->right = buildTreeHelper(inorder, postorder, postIndex, inIndex + 1, inEnd, inMap);
    root->left = buildTreeHelper(inorder, postorder, postIndex, inStart, inIndex - 1, inMap);

    return root;
}

// Main function to construct the tree
TreeNode* buildTree(int inorder[], int postorder[], int n) {
    unordered_map<int, int> inMap;
    for (int i = 0; i < n; i++) {
        inMap[inorder[i]] = i;
    }
    int postIndex = n - 1;  // Start from the last index in postorder
    return buildTreeHelper(inorder, postorder, postIndex, 0, n - 1, inMap);
}

// Function to print inorder traversal of the tree (for verification)
void printInorder(TreeNode* root) {
    if (!root) return;
    printInorder(root->left);
    cout << root->val << " ";
    printInorder(root->right);
}

// Example usage
int main() {
    int inorder[] = {9, 3, 15, 20, 7};
    int postorder[] = {9, 15, 7, 20, 3};
    int n = sizeof(inorder) / sizeof(inorder[0]);

    TreeNode* root = buildTree(inorder, postorder, n);

    cout << "Constructed Inorder Traversal: ";
    printInorder(root);
    cout << endl;

    return 0;
}
