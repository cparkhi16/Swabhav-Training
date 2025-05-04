#include <iostream>
#include <unordered_map>
#include <unordered_set>
#include <vector>
#include <string>

using namespace std;

string ArrayChallenge(string strArr[], int size) {
    unordered_map<int, vector<int>> parentToChildren;
    unordered_map<int, int> childToParent;
    unordered_set<int> nodes;

    // Parse input strings
    for (int i = 0; i < size; i++) {
        int child, parent;
        sscanf(strArr[i].c_str(), "(%d,%d)", &child, &parent);
        
        nodes.insert(child);
        nodes.insert(parent);

        // Check if a child already has a parent (should not)
        if (childToParent.find(child) != childToParent.end()) {
            return "false";
        }
        childToParent[child] = parent;

        // Add child to parent's children list
        parentToChildren[parent].push_back(child);
        
        // Check if a parent has more than 2 children
        if (parentToChildren[parent].size() > 2) {
            return "false";
        }
    }

    // Find the root (a node that is never a child)
    int root = -1;
    for (int node : nodes) {
        if (childToParent.find(node) == childToParent.end()) {
            if (root == -1) {
                root = node;
            } else {
                return "false"; // More than one root
            }
        }
    }

    // If no root found, return false
    if (root == -1) return "false";

    return "true";
}

// Main function
int main() {
    string strArr[] = {"(1,2)", "(2,4)", "(7,2)"};
    int size = sizeof(strArr) / sizeof(strArr[0]);

    cout << ArrayChallenge(strArr, size) << endl;
    return 0;
}
