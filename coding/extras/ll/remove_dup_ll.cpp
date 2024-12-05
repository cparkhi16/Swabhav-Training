#include <iostream>
#include <unordered_set>

struct Node {
    int data;
    Node* next;
    Node(int val) : data(val), next(nullptr) {}
};

// Function to remove duplicates from the linked list
void removeDuplicates(Node* head) {
    if (!head) return;

    std::unordered_set<int> seen; // To store unique values
    Node* current = head;
    Node* prev = nullptr;

    while (current) {
        if (seen.find(current->data) != seen.end()) {
            // Duplicate found, remove the node
            prev->next = current->next;
             Node* temp = current;
             current = current->next;
              delete temp;
        } else {
            // Value not seen before, add to set
            seen.insert(current->data);
            prev = current;
            current = current->next;
        }
         // Move to the next node
    }
}

// Helper function to print the linked list
void printList(Node* head) {
    Node* current = head;
    while (current) {
        std::cout << current->data << " -> ";
        current = current->next;
    }
    std::cout << "nullptr\n";
}

// Helper function to insert a node at the end of the linked list
void insertNode(Node*& head, int value) {
    if (!head) {
        head = new Node(value);
        return;
    }
    Node* temp = head;
    while (temp->next) {
        temp = temp->next;
    }
    temp->next = new Node(value);
}

int main() {
    Node* head = nullptr;
    insertNode(head, 10);
    insertNode(head, 20);
    insertNode(head, 10);
    insertNode(head, 30);
    insertNode(head, 20);

    std::cout << "Original List:\n";
    printList(head);

    removeDuplicates(head);

    std::cout << "List after removing duplicates:\n";
    printList(head);

    return 0;
}
