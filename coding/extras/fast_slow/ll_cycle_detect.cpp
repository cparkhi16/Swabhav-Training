#include <iostream>
using namespace std;

// Definition for singly-linked list.
struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(nullptr) {}
};

// Function to detect cycle
ListNode* detectCycle(ListNode* head) {
    ListNode* slow = head;
    ListNode* fast = head;

    while (fast && fast->next) {
        slow = slow->next;
        fast = fast->next->next;

        if (slow == fast) {
            ListNode* temp = head;
            while (temp != fast) {
                temp = temp->next;
                fast = fast->next;
            }
            return temp;
        }
    }
    return nullptr;
}

// Helper function to create a linked list with a cycle
ListNode* createLinkedListWithCycle(int arr[], int n, int pos) {
    if (n == 0) return nullptr;

    ListNode* head = new ListNode(arr[0]);
    ListNode* curr = head;
    ListNode* cycleNode = nullptr;

    for (int i = 1; i < n; ++i) {
        curr->next = new ListNode(arr[i]);
        curr = curr->next;
        if (i == pos) {
            cycleNode = curr;
        }
    }

    if (pos >= 0) {
        curr->next = cycleNode; // Create the cycle
    }

    return head;
}

// Driver code
int main() {
    int arr[] = {3, 2, 0, -4};
    int n = sizeof(arr) / sizeof(arr[0]);
    int cyclePos = 1; // Cycle starts at index 1 (value 2)

    ListNode* head = createLinkedListWithCycle(arr, n, cyclePos);
    ListNode* cycleStart = detectCycle(head);

    if (cycleStart) {
        cout << "Cycle detected at node with value: " << cycleStart->val << endl;
    } else {
        cout << "No cycle detected." << endl;
    }

    return 0;
}
