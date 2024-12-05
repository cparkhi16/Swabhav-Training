#include <iostream>
#include <bits/stdc++.h>
using namespace std;

class DoublyLL{
    public:
    int data;
    DoublyLL* prev;
    DoublyLL* next;
    
    DoublyLL(int d ): data(d),prev(nullptr),next(nullptr){}
};

void sortedInsert(DoublyLL* head, DoublyLL* newnode){
    if( !newnode ) return;
    
    if( head == nullptr ){
        head = newnode;
        return;
    }
    
    if( head->data > newnode->data){
        newnode->next = head->next;
        head->prev = newnode;
        head = newnode;
        return;
    }
    
    DoublyLL* curr = head;
    while(curr->next && curr->next->data < newnode->data){
        curr=curr->next;
    }
    newnode->next = curr->next;
    if(curr->next)
        curr->next->prev = newnode;
    curr->next = newnode;
    newnode->prev = curr;
}

void removeNode(DoublyLL* head, DoublyLL* node){
    if(!head || !node) return;
    
    if(head == node){
        head = node->next;
    }
    
    if(node->next)
        node->next->prev = node->prev;
        
    if(node->prev)
        node->prev->next = node->next;
        
    node->prev = node->next = nullptr;
}

void modifyList(DoublyLL* head , int c, int k){
    DoublyLL* curr = head;
    
    while(curr != nullptr && curr->data != k)
        curr = curr->next;
    
    curr->data = curr->data - c;
    
    removeNode(head , curr);
    
    sortedInsert(head , curr);
}

void printList(DoublyLL* head) {
    DoublyLL* current = head;
    while (current) {
        cout << current->data << " ";
        current = current->next;
    }
    cout << endl;
}

int main() {
    DoublyLL* head = new DoublyLL(10);
    sortedInsert(head, new DoublyLL(20));
    sortedInsert(head, new DoublyLL(30));
    sortedInsert(head, new DoublyLL(40));

    cout << "Original List: ";
    printList(head);

    int C = 15, K = 30;
    modifyList(head, C, K);

    cout << "Modified List: ";
    printList(head);

    return 0;
}
