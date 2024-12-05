// Online C++ compiler to run C++ program online
#include <iostream>
#include<bits/stdc++.h>

using namespace std;

class DLL{
    public:
    int key;
    int value;
    DLL* next;
    DLL* prev;
    DLL(int k , int v): key(k), value(v), next(nullptr), prev(nullptr){}
};
class LRUCache{
    private:
    DLL* head;
    DLL* tail;
    unordered_map<int , DLL*> map;
    int cap;
    
    void moveToFront(DLL* node){
        node->prev->next = node->next;
       if(node->next){
           node->next->prev = node->prev;
       }
       
       node->next = head;
       node->prev = nullptr;
       if(head)
          head->prev = node;
        head= node;
        
        if(!tail)
            tail = head;
    }
    
    void removeTail(){
        if(!tail) return;
        DLL* temp = tail;
        tail->prev->next = nullptr;
        tail = tail->prev;
        delete temp;
    }
    
    public:
    LRUCache(int c ): cap(c), head(nullptr), tail(nullptr){}
    
    int get( int k){
        if(map.find(k) != map.end()){
            moveToFront(map[k]);
            return map[k]->value;
        }
        return -1;
    }
    
    void put(int k , int v){
        if(map.find(k) != map.end()){
            auto val = map[k];
            val->value = v;
            moveToFront(val);
            return;
        }
        else{
            if(map.size() == cap){
                map.erase(tail->key);
                removeTail();
            }
            DLL* new_node = new DLL(k,v);
            //new_node->key = k;
            //new_node->value = v;
            map[k] = new_node;
            if(!head){
                head = tail = new_node;
            }else{
            new_node->next = head;
            head->prev = new_node;
            head = new_node;
            }
            
        }
    }
        
};
int main() {
    LRUCache lruCache(3); // Cache capacity is 3

    lruCache.put(1, 1); // Cache: [1]
    lruCache.put(2, 2); // Cache: [2, 1]
    cout << lruCache.get(1) << endl; // Returns 1, Cache: [1, 2]
    lruCache.put(3, 3); // Cache: [3, 1, 2]
    cout << lruCache.get(2) << endl; // Returns 2, Cache: [2, 3, 1]
    lruCache.put(4, 4); // Cache: [4, 2, 3]
    cout << lruCache.get(1) << endl; // Returns -1 (not found)
    cout << lruCache.get(3) << endl; // Returns 3, Cache: [3, 4, 2]
    return 0;
}