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
        ListNode* lastNodeOfPrevPart = previous; // 1
        ListNode* lastNodeOfSubList = current; // 3

        ListNode* next = nullptr;
        int i = 0;
        if (getRemainingLen(current) < k) { // if we comment this part , remaining nodes will also be reversed if len of them is not multiple of k
            break;
        }
        while (current && i < k) {
            next = current->next;
            current->next = previous;
            previous = current;
            current = next;
            i++;
        }
        // 1->2 , 2->1 , prev = 2 , curr = 3 ,lNS = 1
        if (lastNodeOfPrevPart == nullptr) {
            head = previous;
        } else {
            lastNodeOfPrevPart->next = previous;
        }
        lastNodeOfSubList->next = current; // 2->1->3->4

        if (current == nullptr)
            break;

        previous = lastNodeOfSubList; // prev = 1
    }
    return head;
}

// Main function to test the reverseKGroup function.
int main() {
    int arr[] = {1, 2, 3, 4, 5};
    int k = 3;
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
