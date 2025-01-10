#include <iostream>
using namespace std;

struct ListNode {
    int val;
    ListNode* next;
    ListNode(int x) : val(x), next(nullptr) {}
};

class Solution {
public:
    ListNode* removeNthFromEnd(ListNode* head, int n) {
        // Create a dummy node to simplify edge cases
        ListNode* dummy = new ListNode(0);
        dummy->next = head;

        ListNode* first = dummy;
        ListNode* second = dummy;

        // Step 1: Move the first pointer `n + 1` steps ahead
        for (int i = 0; i <= n; ++i) {
            first = first->next;
        }

        // Step 2: Move both pointers until the first pointer reaches the end
        while (first) {
            first = first->next;
            second = second->next;
        }
        cout<<" second is pointing to "<<second->val<<endl;
        // Step 3: Remove the nth node
        ListNode* nodeToDelete = second->next;
        second->next = second->next->next;
        delete nodeToDelete; // Free memory

        // Step 4: Return the new head of the list
        ListNode* newHead = dummy->next;
        delete dummy; // Free memory for the dummy node
        return newHead;
    }
};

void printList(ListNode* head) {
    while (head) {
        cout << head->val << " -> ";
        head = head->next;
    }
    cout << "NULL" << endl;
}

int main() {
    // Example linked list: 1 -> 2 -> 3 -> 4 -> 5
    ListNode* head = new ListNode(1);
    head->next = new ListNode(2);
    head->next->next = new ListNode(3);
    head->next->next->next = new ListNode(4);
    head->next->next->next->next = new ListNode(5);

    cout << "Original List: ";
    printList(head);

    Solution solution;
    head = solution.removeNthFromEnd(head, 2); // Remove the 2nd node from the end

    cout << "After Removing 2nd Node from End: ";
    printList(head);

    return 0;
}
