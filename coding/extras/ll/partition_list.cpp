#include <iostream>
using namespace std;

struct ListNode {
    int val;
    ListNode* next;
    ListNode(int x) : val(x), next(nullptr) {}
};

ListNode* partition(ListNode* head, int x) {
    ListNode* l1 = new ListNode(0);  // Dummy node for the "less than x" list
    ListNode* l2 = new ListNode(0);  // Dummy node for the "greater or equal to x" list
    ListNode* l1Head = l1;
    ListNode* l2Head = l2;
    
    while (head) {
        if (head->val < x) {
            l1->next = head;
            l1 = l1->next;
        } else {
            l2->next = head;
            l2 = l2->next;
        }
        head = head->next;
    }
    
    l2->next = nullptr;  // Important: terminate the second list
    l1->next = l2Head->next;  // Connect the two lists
    return l1Head->next;  // Return the new head (skipping dummy node)
}

void printList(ListNode* head) {
    while (head) {
        cout << head->val << " ";
        head = head->next;
    }
    cout << endl;
}

int main() {
    // Creating the linked list: 1 -> 4 -> 3 -> 2 -> 5 -> 2
    ListNode* head = new ListNode(1);
    head->next = new ListNode(4);
    head->next->next = new ListNode(3);
    head->next->next->next = new ListNode(2);
    head->next->next->next->next = new ListNode(5);
    head->next->next->next->next->next = new ListNode(2);

    int x = 3;
    cout << "Original list: ";
    printList(head);
    
    ListNode* partitioned = partition(head, x);
    
    cout << "Partitioned list: ";
    printList(partitioned);

    return 0;
}
