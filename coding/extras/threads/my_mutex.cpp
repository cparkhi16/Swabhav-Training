#include<iostream>
#include<bits/stdc++.h>

using namespace std;
std::mutex mtx;
void thread_func(int id ){
    std::lock_guard<std::mutex> lock (mtx);
    cout<<" Thread id "<<id<<" entering critical section "<<endl;
    std::this_thread::sleep_for(chrono::seconds(1));
    cout<<" Thread id "<<id<<" done critical section "<<endl;
}

int main(){
    thread t1 (thread_func,1);
    thread t2 (thread_func , 2);

    t1.join();
    t2.join();
}