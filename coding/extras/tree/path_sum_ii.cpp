#include <iostream>
#include <vector>
using namespace std;

/**
 * Definition for a binary tree node.
 */
struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode() : val(0), left(nullptr), right(nullptr) {}
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
    TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
};

class Solution {
public:
    void dfs(TreeNode* node, int targetSum, vector<int>& path, vector<vector<int>>& result) {
        if (!node) return; // If the node is null, return
        
        path.push_back(node->val);  // Add the current node to the path
        
        // Check if we are at a leaf node and if the current path sum equals targetSum
        if (!node->left && !node->right && targetSum == node->val) {
            result.push_back(path);  // Add the path to the result
        }
        
        // Recursively search the left and right subtrees with updated targetSum
        dfs(node->left, targetSum - node->val, path, result);
        dfs(node->right, targetSum - node->val, path, result);
        
        path.pop_back();  // Backtrack to explore other paths
        
    }

    vector<vector<int>> pathSum(TreeNode* root, int targetSum) {
        vector<vector<int>> result;  // Store all the valid paths
        vector<int> path;  // Temporary vector to store the current path
        dfs(root, targetSum, path, result);  // Start DFS from the root
        return result;
    }
};

// Helper function to create a new tree node
TreeNode* newNode(int val) {
    return new TreeNode(val);
}

// Driver code
int main() {
    // Construct the following binary tree:
    //         5
    //        / \
    //       4   8
    //      /   / \
    //     11  13  4
    //    /  \    / \
    //   7    2  5   1

    TreeNode* root = newNode(5);
    root->left = newNode(4);
    root->right = newNode(8);
    root->left->left = newNode(11);
    root->left->left->left = newNode(7);
    root->left->left->right = newNode(2);
    root->right->left = newNode(13);
    root->right->right = newNode(4);
    root->right->right->left = newNode(5);
    root->right->right->right = newNode(1);

    Solution sol;
    int targetSum = 22;

    // Get the paths with the given target sum
    vector<vector<int>> result = sol.pathSum(root, targetSum);

    // Print the result
    cout << "Paths with sum " << targetSum << ":" << endl;
    for (const auto& path : result) {
        for (int val : path) {
            cout << val << " ";
        }
        cout << endl;
    }

    return 0;
}
