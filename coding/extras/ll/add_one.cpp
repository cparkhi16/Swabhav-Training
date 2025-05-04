#include <iostream>
using namespace std;

struct ListNode {
    int val;
    ListNode* next;
    ListNode(int x) : val(x), next(nullptr) {}
};

// Helper function to reverse a linked list
ListNode* reverseList(ListNode* head) {
    ListNode* prev = nullptr;
    ListNode* curr = head;
    while (curr) {
        ListNode* nextTemp = curr->next;
        curr->next = prev;
        prev = curr;
        curr = nextTemp;
    }
    return prev;
}

// Function to add 1 to the linked list
ListNode* addOne(ListNode* head) {
    // Step 1: Reverse the list
    head = reverseList(head);

    // Step 2: Add 1 to the number
    ListNode* curr = head;
    int carry = 1;
    while (curr && carry) {
        int sum = curr->val + carry;
        curr->val = sum % 10;
        carry = sum / 10;
        if (!curr->next && carry) {
            // Add a new node if carry is left and we're at the last node
            curr->next = new ListNode(carry);
            carry = 0;
        }
        curr = curr->next;
    }

    // Step 3: Reverse the list back to the original order
    return reverseList(head);
}

// Helper function to print the linked list
void printList(ListNode* head) {
    while (head) {
        cout << head->val;
        if (head->next) cout << " -> ";
        head = head->next;
    }
    cout << endl;
}

int main() {
    // Example: 1 -> 9 -> 9
    ListNode* head = new ListNode(1);
    head->next = new ListNode(9);
    head->next->next = new ListNode(9);

    cout << "Original list: ";
    printList(head);

    head = addOne(head);

    cout << "After adding 1: ";
    printList(head);

    return 0;
}
