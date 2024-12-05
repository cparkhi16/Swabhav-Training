#include <iostream>
#include <vector>
using namespace std;

struct Node {
    int data;
    vector<Node*> children;
};

int maxSumRootToLeaf(Node* root) {
    if (!root) return 0;
    if (root->children.empty()) return root->data;

    int maxSum = 0;
    for (Node* child : root->children) {
        maxSum = max(maxSum, maxSumRootToLeaf(child));
    }

    return root->data + maxSum;
}

Node* createNode(int data) {
    Node* newNode = new Node();
    newNode->data = data;
    return newNode;
}

int main() {
    // Creating an example n-ary tree
    Node* root = createNode(1);
    root->children.push_back(createNode(2));
    root->children.push_back(createNode(3));
    root->children[0]->children.push_back(createNode(4));
    root->children[0]->children.push_back(createNode(5));
    root->children[1]->children.push_back(createNode(6));
    root->children[1]->children.push_back(createNode(7));

    cout << "Maximum Sum from Root to Leaf: " << maxSumRootToLeaf(root) << endl;

    return 0;
}
