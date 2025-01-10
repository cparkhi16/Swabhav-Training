#include <iostream>
using namespace std;

struct Node {
    int data;
    Node* next;
    Node(int val) : data(val), next(nullptr) {}
};

// Function to rearrange the linked list
Node* rearrangeLinkedList(Node* head) {
    if (!head || !head->next) return head;

    // Step 1: Use fast and slow pointers to find the middle of the LinkedList
    Node* slow = head;
    Node* fast = head;

    while (fast && fast->next) {
        fast = fast->next->next;
        slow = slow->next;
    }

    // Step 2: Reverse the second half of the LinkedList
    Node* prev = nullptr;
    Node* curr = slow;
    Node* temp;
    while (curr) {
        temp = curr->next;
        curr->next = prev;
        prev = curr;
        curr = temp;
    }

    // Step 3: Interleave the first half with the reversed second half
    Node* first = head;
    Node* second = prev;

    while (second->next) { // Interleave nodes
        temp = first->next;
        first->next = second;
        first = temp;

        temp = second->next;
        second->next = first;
        second = temp;
    }

    return head;
}

// Helper function to print the linked list
void printList(Node* head) {
    Node* current = head;
    while (current) {
        cout << current->data << " -> ";
        current = current->next;
    }
    cout << "null" << endl;
}

// Helper function to create a linked list from an array
Node* createList(const int arr[], int size) {
    if (size == 0) return nullptr;

    Node* head = new Node(arr[0]);
    Node* tail = head;

    for (int i = 1; i < size; i++) {
        tail->next = new Node(arr[i]);
        tail = tail->next;
    }

    return head;
}

int main() {
    int arr[] = {1, 2, 3, 4, 5, 6};
    int size = sizeof(arr) / sizeof(arr[0]);

    // Create the linked list
    Node* head = createList(arr, size);

    cout << "Original Linked List: ";
    printList(head);

    // Rearrange the linked list
    head = rearrangeLinkedList(head);

    cout << "Rearranged Linked List: ";
    printList(head);

    return 0;
}
