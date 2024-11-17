#include <iostream>
using namespace std;

// Node structure
struct Node {
    int data;
    Node* next;
    
    Node(int value) : data(value), next(nullptr) {}
};

// Function to calculate the GCD of two numbers
int gcd(int a, int b) {
    // if (a == 0)
    //     return b;
    // return gcd(b % a, a);
    cout<<" a "<<a<<" b "<<b<<endl;
    while (b != 0) {
        int temp = b;
        b = a % b;
        a = temp;
        cout<<" b now "<<b<<endl;
        cout<<" a now "<<a<<endl;
    }
    return a;
}

// Function to insert a node after a given node
void insertAfter(Node* prev_node, int value) {
    if (prev_node == nullptr) {
        cout << "Previous node cannot be NULL" << endl;
        return;
    }
    
    Node* new_node = new Node(value);
    new_node->next = prev_node->next;
    prev_node->next = new_node;
}

// Function to insert GCD between pairs in the linked list
void insertGCDInLinkedList(Node* head) {
    Node* current = head;
    
    // Traverse the linked list until the second last node
    while (current != nullptr && current->next != nullptr) {
        int gcdValue = gcd(current->data, current->next->data);
        insertAfter(current, gcdValue);
        current = current->next->next; // Move to the next original node
    }
}

// Function to print the linked list
void printList(Node* head) {
    Node* temp = head;
    while (temp != nullptr) {
        cout << temp->data << " -> ";
        temp = temp->next;
    }
    cout << "NULL" << endl;
}

// Driver code
int main() {
    // Creating a linked list: 12 -> 15 -> 25 -> NULL
    Node* head = new Node(12);
    head->next = new Node(15);
    head->next->next = new Node(25);

    cout << "Original Linked List: ";
    printList(head);
    
    // Insert GCDs in between
    insertGCDInLinkedList(head);
    
    cout << "Linked List after inserting GCDs: ";
    printList(head);

    return 0;
}
