#include <iostream>
using namespace std;

// Node structure
struct Node {
    int data;
    Node* next;
    Node* prev;

    Node(int val) : data(val), next(nullptr), prev(nullptr) {}
};

// Function to reverse the doubly linked list
void reverseDoublyLinkedList(Node*& head) {
    if (!head) return; // If the list is empty, nothing to reverse

    Node* current = head;
    Node* temp = nullptr;

    // Traverse the list and swap next and prev for each node
    while (current) {
        // Swap the next and prev pointers
        temp = current->prev;
        current->prev = current->next;
        current->next = temp;

        // Move to the next node (previous in original list)
        current = current->prev;
    }

    // After the loop, temp will point to the new head (last node of the original list)
    if (temp) {
        head = temp->prev;
    }
}

// Function to print the doubly linked list
void printList(Node* head) {
    Node* current = head;
    while (current) {
        cout << current->data << " ";
        current = current->next;
    }
    cout << endl;
}

// Function to insert a node at the end of the list
void append(Node*& head, int data) {
    Node* newNode = new Node(data);
    if (!head) {
        head = newNode;
        return;
    }

    Node* current = head;
    while (current->next) {
        current = current->next;
    }
    current->next = newNode;
    newNode->prev = current;
}

// Main function to test the reversal
int main() {
    Node* head = nullptr;

    // Create a doubly linked list
    append(head, 10);
    append(head, 20);
    append(head, 30);
    append(head, 40);

    cout << "Original List: ";
    printList(head);

    // Reverse the doubly linked list
    reverseDoublyLinkedList(head);

    cout << "Reversed List: ";
    printList(head);

    return 0;
}
