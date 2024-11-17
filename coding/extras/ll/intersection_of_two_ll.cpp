#include <iostream>

struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

int getLength(ListNode* head) {
    int length = 0;
    while (head) {
        length++;
        head = head->next;
    }
    return length;
}

ListNode *getIntersectionNode(ListNode *headA, ListNode *headB) {
    int lenA = getLength(headA);
    int lenB = getLength(headB);
    
    // Align the start of both linked lists
    while (lenA > lenB) {
        headA = headA->next;
        lenA--;
    }
    while (lenB > lenA) {
        headB = headB->next;
        lenB--;
    }
    
    // Traverse both lists together until we find the intersection
    while (headA && headB) {
        if (headA == headB) {
            return headA;
        }
        headA = headA->next;
        headB = headB->next;
    }
    
    return NULL; // No intersection found
}

int main() {
    // Example usage:
    ListNode *listA = new ListNode(1);
    listA->next = new ListNode(2);
    ListNode *intersection = new ListNode(3);
    listA->next->next = intersection;
    listA->next->next->next = new ListNode(4);
    // 1 -> 2 -> 3 -> 4
    // 5 -> 3
    ListNode *listB = new ListNode(5);
    listB->next = intersection;
    
    ListNode *result = getIntersectionNode(listA, listB);
    if (result) {
        std::cout << "Intersection at node with value: " << result->val << std::endl;
    } else {
        std::cout << "No intersection found." << std::endl;
    }
    
    // Clean up memory
    delete listA->next->next->next; // node with value 4
    delete listA->next->next; // node with value 3 (intersection)
    delete listA->next; // node with value 2
    delete listA; // node with value 1
    delete listB; // node with value 5
    
    return 0;
}
