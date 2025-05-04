// Online C++ compiler to run C++ program online
#include <iostream>
#include <bits/stdc++.h>
using namespace std;

class PriorityQueue{
    private:
        vector<int> heap;
        
        void siftup(int index){
            int parent = (index-1) / 2;
            while (parent > 0 && heap[parent] > heap[index]){
                swap(heap[parent], heap[index]);
                index = parent;
                parent = (index-1) / 2;
            }
        }
        
        void siftdown(int index){
            int left = (2 * index) + 1;
            int right = (2 * index) + 2;
            int smallest = index;
            
            if(left<heap.size() && heap[left] < heap[smallest]){
                smallest = left;
            }
            
            if(right< heap.size() && heap[right] < heap[smallest]){
                smallest = right;
            }
            
            if(smallest != index){
                swap(heap[smallest], heap[index]);
                siftdown(smallest);
            }
        }
    public:
        void heapSortAscending(vector<int> nums){
            heap = nums;
            for(int i = heap.size() -1 ; i >=0 ; i--){
                siftdown(i);
            }
        }
        
        void push(int data){
            heap.push_back(data);
            siftup(heap.size()-1);
        }
        
        void pop(){
            if(heap.empty()){
                return;
            }
            heap[0] = heap.back();
            heap.pop_back();
            siftdown(0);
        }
        
        int top(){
            if(heap.empty()){
                return -1;
            }
            return heap[0];
        }
        
        bool empty() {
        return heap.empty();
        }

    // Get the size of the priority queue
    int size() {
        return heap.size();
    }

    // Display the elements (for debugging purposes)
    void display() {
        for (int val : heap) {
            std::cout << val << " ";
        }
        std::cout << std::endl;
    }
    
    void getElemAfterSort(){
        while(!heap.empty()){
            std::cout<<top()<<endl;
            pop();
        }
    }
};
int main() {
    PriorityQueue pq;
    pq.push(10);
    pq.push(20);
    pq.push(15);
    pq.push(30);

    std::cout << "Priority queue after inserts: ";
    pq.display();

    // Get the top element
    std::cout << "Top element: " << pq.top() << std::endl;

    // Remove the top element
    pq.pop();
    std::cout << "Priority queue after pop: ";
    pq.display();

    // Get the top element again
    std::cout << "Top element: " << pq.top() << std::endl;

    vector<int> nums = {1,7,81,3,0,100};
    pq.heapSortAscending(nums);
    std::cout << "after heap sort in asc order";
    pq.getElemAfterSort();
    return 0;
}