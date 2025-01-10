// Online C++ compiler to run C++ program online
#include <iostream>
#include<bits/stdc++.h>
using namespace std;

struct Node{
  int data;
  Node* next;
  
  Node(int d):data(d), next(nullptr){}
};
bool isPalindrome(Node* head){
    Node* slow = head;
    Node* fast = head;
    
    while(fast && fast->next){
        fast= fast->next->next;
        slow = slow->next;
    }
    Node* prev = nullptr;
    Node* n;
    while(slow){
        n = slow->next;
        slow->next = prev;
        prev = slow;
        slow = n;
    }
    while(head && prev){
        if(head->data != prev->data){
            return false;
        }
        head=head->next;
        prev = prev->next;
    }
    return true;
}
int main() {
    Node* head = new Node(1);
    head->next = new Node(2);
    head->next->next = new Node(2);
    head->next->next->next = new Node(1);
    
    if(isPalindrome(head)){
        cout<<" LL is a palindrome"<<endl;
    }else{
        cout<<" LL is not a palindrome "<<endl;
    }

    return 0;
}