#include <iostream>
#include <unordered_set>
using namespace std;

struct Node {
    int data;
    Node* next;
};

// Remove duplicates, keeping one occurrence
void removeDuplicates(Node*& head) {
    if (!head || !head->next) return; // Empty or single-node list

    unordered_set<int> seen; // Track seen values
    Node* current = head;
    Node* prev = nullptr;

    do {
        if (seen.count(current->data)) {
            // Duplicate found; remove current node
            prev->next = current->next;
            Node* temp = current;
            current = current->next;
            delete temp;
        } else {
            // Not a duplicate; add to 'seen' and move forward
            seen.insert(current->data);
            prev = current;
            current = current->next;
        }
    } while (current != head); // Stop when we've come full circle

    // Handle case where the last node connects back to the head and creates a duplicate
    // if (seen.count(head->data) && prev->next == head) {
    //     prev->next = head->next;
    //     Node* temp = head;
    //     head = head->next;
    //     delete temp;
    // }
}

void display(Node* head) {
    if (!head) return;
    Node* current = head;
    do {
        cout << current->data << " ";
        current = current->next;
    } while (current != head);
    cout << endl;
}

Node* createNode(int data) {
    Node* newNode = new Node();
    newNode->data = data;
    newNode->next = nullptr;
    return newNode;
}

void append(Node*& head, int data) {
    Node* newNode = createNode(data);
    if (!head) {
        head = newNode;
        head->next = head; // Point to itself to make circular
    } else {
        Node* temp = head;
        while (temp->next != head) {
            temp = temp->next;
        }
        temp->next = newNode;
        newNode->next = head;
    }
}

int main() {
    Node* head = nullptr;

    // Example circular linked list: 1 → 2 → 2 → 3 → 3 → 1 (circular)
    append(head, 1);
    append(head, 2);
    append(head, 2);
    append(head, 3);
    append(head, 3);
    append(head, 1);

    cout << "Original list: ";
    display(head);

    removeDuplicates(head);

    cout << "After removing duplicates: ";
    display(head);

    return 0;
}
