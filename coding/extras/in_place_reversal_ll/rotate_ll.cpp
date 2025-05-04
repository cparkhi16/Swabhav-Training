#include <iostream>
using namespace std;

// Definition for singly-linked list
struct ListNode {
    int val;
    ListNode* next;
    ListNode(int x) : val(x), next(nullptr) {}
};

// Function to rotate the linked list
ListNode* rotateRight(ListNode* head, int k) {
    if (!head || !head->next || k == 0) return head; // Handle edge cases

    // Step 1: Find the length of the list and make it circular
    int len = 1; // Start with 1 to count the head
    ListNode* tail = head;
    while (tail->next) {
        tail = tail->next;
        len++;
    }
    tail->next = head; // Make the list circular

    // Step 2: Calculate the effective rotations
    k = k % len;
    int stepsToNewHead = len - k;

    // Step 3: Break the circular list at the new head
    ListNode* newTail = tail;
    while (stepsToNewHead--) {
        newTail = newTail->next;
    }
    ListNode* newHead = newTail->next;
    newTail->next = nullptr;

    return newHead;
}

// Helper function to create a linked list from an array
ListNode* createList(int arr[], int n) {
    if (n == 0) return nullptr;
    ListNode* head = new ListNode(arr[0]);
    ListNode* curr = head;
    for (int i = 1; i < n; i++) {
        curr->next = new ListNode(arr[i]);
        curr = curr->next;
    }
    return head;
}

// Helper function to print a linked list
void printList(ListNode* head) {
    while (head != nullptr) {
        cout << head->val << " ";
        head = head->next;
    }
    cout << endl;
}

int main() {
    // Input linked list
    int arr[] = {1, 2, 3};
    int n = sizeof(arr) / sizeof(arr[0]);
    ListNode* head = createList(arr, n);

    // Rotation value
    int k = 2000000000;

    cout << "Original List: ";
    printList(head);

    // Rotate the list
    head = rotateRight(head, k);

    cout << "Rotated List: ";
    printList(head);

    return 0;
}
