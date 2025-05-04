#include <iostream>
using namespace std;

// Definition for singly-linked list.
struct ListNode {
    int val;
    ListNode* next;
    ListNode(int x) : val(x), next(nullptr) {}
};

// Function to create a linked list from an array.
ListNode* createLinkedList(int arr[], int size) {
    if (size == 0) return nullptr;
    ListNode* head = new ListNode(arr[0]);
    ListNode* current = head;
    for (int i = 1; i < size; i++) {
        current->next = new ListNode(arr[i]);
        current = current->next;
    }
    return head;
}

// Function to print a linked list.
void printLinkedList(ListNode* head) {
    ListNode* current = head;
    while (current != nullptr) {
        cout << current->val << " ";
        current = current->next;
    }
    cout << endl;
}

// Function to get the remaining length of the linked list.
int getRemainingLen(ListNode* ptr) {
    int len = 0;
    while (ptr) {
        ptr = ptr->next;
        len++;
    }
    return len;
}

// Function to reverse nodes in k-group.
ListNode* reverseKGroup(ListNode* head, int k) {
    ListNode* previous = nullptr;
    ListNode* current = head;
    while (true) {
        ListNode* lastNodeOfPrevPart = previous;
        ListNode* lastNodeOfSubList = current;
        
        ListNode* next = nullptr;
        int i = 0;
        if (getRemainingLen(current) < k) {
            break;
        }
        while (current && i < k) {
            next = current->next;
            current->next = previous;
            previous = current;
            current = next;
            i++;
        }

        if (lastNodeOfPrevPart == nullptr) {
            head = previous;
        } else {
            lastNodeOfPrevPart->next = previous;
        }
        lastNodeOfSubList->next = current;
        
        if (current == nullptr)
            break;

        i = 0;
        while(current && i < k){
            previous = current;
            current = current->next;
            i++;
        }
        
    }
    return head;
}

// Main function to test the reverseKGroup function.
int main() {
    int arr[] = {1, 2, 3, 4, 5, 6,7,8};
    int k = 2;
    int size = sizeof(arr) / sizeof(arr[0]);

    // Create the linked list.
    ListNode* head = createLinkedList(arr, size);

    cout << "Original Linked List: ";
    printLinkedList(head);

    // Reverse in k-groups.
    head = reverseKGroup(head, k);

    cout << "Reversed in k-groups: ";
    printLinkedList(head);

    return 0;
}
