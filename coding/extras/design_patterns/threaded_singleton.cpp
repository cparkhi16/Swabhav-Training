#include <iostream>
#include <bits/stdc++.h>
using namespace std;

class Singleton {
public:
    Singleton(const Singleton& o) = delete;
    Singleton operator=(const Singleton& o)= delete;
    static Singleton& getInstance() {
        if (!instance) {
            lock_guard<mutex> lock(mtx);
            if(!instance){
            instance = new Singleton();
            }
        }
        return *instance;
    }

    void getLog() {
        cout << " log from singleton " << endl;
    }

private:
    Singleton() { cout << " singleton instance created " << endl; }
    static Singleton* instance;
    static mutex mtx;
};

// Initialize the static member outside the class
Singleton* Singleton::instance = nullptr;
mutex Singleton::mtx;

void func(int i){
     Singleton& l = Singleton::getInstance();  // Access the Singleton instance
    l.getLog();
}
int main() {
   std::vector<std::thread> threads;
   
   for(int i = 0 ; i < 5 ; i++){
       threads.push_back(thread(func, i));
   }
   
   for(int i = 0 ; i < 5 ; i++){
       threads[i].join();
   }

    return 0;
}
