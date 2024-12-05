#include <iostream>
#include <vector>

using namespace std;

// Definition for singly-linked list.
struct ListNode {
    int val;
    ListNode* next;
    ListNode(int x) : val(x), next(nullptr) {}
};

// Function to split the linked list into k parts
vector<ListNode*> splitListToParts(ListNode* head, int k) {
    vector<ListNode*> result(k, nullptr);
    int length = 0;

    // Calculate the total length of the linked list
    ListNode* temp = head;
    while (temp) {
        length++;
        temp = temp->next;
    }

    // Determine the size of each part
    int partSize = length / k;
    int extra = length % k; // Extra nodes to distribute

    ListNode* current = head;
    ListNode* prev = nullptr;

    for (int i = 0; i < k && current; i++) {
        result[i] = current; // Start of the current part
        int currentPartSize = partSize + (extra > 0 ? 1 : 0);
        extra--;

        // Move the pointer to the end of the current part
        for (int j = 0; j < currentPartSize; j++) {
            prev = current;
            current = current->next;
        }

        // Disconnect the current part from the next
        if (prev) prev->next = nullptr;
    }

    return result;
}

// Helper function to create a linked list from a vector
ListNode* createLinkedList(const vector<int>& values) {
    if (values.empty()) return nullptr;

    ListNode* head = new ListNode(values[0]);
    ListNode* current = head;

    for (size_t i = 1; i < values.size(); i++) {
        current->next = new ListNode(values[i]);
        current = current->next;
    }

    return head;
}

// Helper function to print the linked list
void printLinkedList(ListNode* head) {
    while (head) {
        cout << head->val << " ";
        head = head->next;
    }
    cout << endl;
}

// Test the function
int main() {
    vector<int> values = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10};
    ListNode* head = createLinkedList(values);

    int k = 3;
    vector<ListNode*> parts = splitListToParts(head, k);

    for (int i = 0; i < k; i++) {
        cout << "Part " << i + 1 << ": ";
        printLinkedList(parts[i]);
    }

    return 0;
}
