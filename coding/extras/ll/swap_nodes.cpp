#include <iostream>
using namespace std;

struct Node {
    int data;
    Node* next;
    Node(int val) : data(val), next(nullptr) {}
};

Node* swapPairsRecursive(Node* head) {
    if (!head || !head->next) return head; // Base case

    Node* first = head;
    Node* second = head->next;

    // Recursive call to swap the rest of the list
    first->next = swapPairsRecursive(second->next);

    // Swap the current pair
    second->next = first;

    return second;
}

void printList(Node* head) {
    while (head) {
        cout << head->data << " -> ";
        head = head->next;
    }
    cout << "nullptr" << endl;
}

int main() {
    Node* head = new Node(1);
    head->next = new Node(2);
    head->next->next = new Node(3);
    //head->next->next->next = new Node(4);

    cout << "Original List: ";
    printList(head);

    head = swapPairsRecursive(head);

    cout << "Swapped List: ";
    printList(head);

    return 0;
}
