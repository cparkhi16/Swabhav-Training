#include <iostream>
using namespace std;

// Definition for singly-linked list.
struct ListNode {
    int val;
    ListNode* next;
    ListNode(int x) : val(x), next(nullptr) {}
};

// Function to reverse a sublist of a linked list
ListNode* reverseBetween(ListNode* head, int left, int right) {
    ListNode* dummy = new ListNode(0);
    dummy->next = head;
    
    ListNode* prev = dummy;
    ListNode* next = nullptr;
    
    for(int i = 1 ; i < left ; i++){
        prev = prev->next;
    }
    ListNode* curr = prev->next;
    
    for(int j = 0 ; j < right-left ; j++){
        next = curr->next;
        curr->next = next->next;
        next->next = prev->next;
        prev->next = next;
    }
    return dummy->next;
}

// Function to print a linked list
void printList(ListNode* head) {
    while (head) {
        cout << head->val << " ";
        head = head->next;
    }
    cout << endl;
}

// Function to create a linked list from an array
ListNode* createList(const int arr[], int size) {
    if (size == 0) return nullptr;
    ListNode* head = new ListNode(arr[0]);
    ListNode* current = head;
    for (int i = 1; i < size; ++i) {
        current->next = new ListNode(arr[i]);
        current = current->next;
    }
    return head;
}

// Driver code
int main() {
    // Input linked list
    int arr[] = {1, 2, 3, 4,5};
    int n = sizeof(arr) / sizeof(arr[0]);
    ListNode* head = createList(arr, n);

    cout << "Original List: ";
    printList(head);

    // Reverse the sublist from position 1 to 4
    int left = 2, right = 4;
    head = reverseBetween(head, left, right);

    cout << "Reversed List: ";
    printList(head);

    return 0;
}
