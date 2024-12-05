#include<bits/stdc++.h>

using namespace std;

struct Node{
    int data;
    Node* next;

    Node(int d): data(d), next(nullptr){}
};

void split(Node* head , vector<Node*>& node_parts, int k){
    int len = 0;
    Node* curr = head;
    while(curr){
        len++;
        curr = curr->next;
    }
    cout<<" len is "<<len<<endl;
    int partition = len/k;
    int extra = len%k;

    Node* prev = nullptr;
    Node* c = head;
    for(int i = 0 ; i < k ; i++){
        cout<<" i is "<<i<<endl;
        node_parts[i] = c;
        int final_size = partition + (extra > 0 ? 1 : 0);
        extra--;
        cout<<" final size "<<final_size<<endl;
        for(int j = 0 ; j < final_size ; j++){
            cout<<" in f size "<<j<<endl;
            prev= c;
            c= c->next;
        }

        if(prev) prev->next = nullptr;

    }
}

void printParts(Node* head, int k){
    vector<Node*> res(k);
    split(head , res , k);
    for(int i = 0 ; i < k ; i++){
        cout<<"Part :: "<<i<<endl;
        Node* cur = res[i];
        while(cur){
            cout<<" "<<cur->data<<" ";
            cur = cur->next;
        }
        cout<<endl;
    }
}
Node* createLinkedList(vector<int> val){
    Node* head = new Node(val[0]);
    Node* curr = head;

    for(int i = 1 ; i < val.size() ; i++){
        curr->next = new Node(val[i]);
        curr = curr->next;
    }
    return head;
}
int main(){
    vector<int> values = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10};
    Node* head = createLinkedList(values);
    cout<<" got head "<<endl;
    int k = 3;
    printParts(head,k);
}