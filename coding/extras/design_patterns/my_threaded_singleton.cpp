#include<iostream>
#include<bits/stdc++.h>
using namespace std;

template<typename T>
class Singleton{
    public:
        static T& get_instance(){
            if(!instance){
                lock_guard<mutex> lock(mtx);
                if(!instance){
                    instance = make_unique<T>();
                }
            }
            return *instance;
        }
    private:
        static unique_ptr<T> instance;
        static std::mutex mtx;
};

template<typename T>
unique_ptr<T> Singleton<T>::instance = nullptr;

template<typename T>
mutex Singleton<T>::mtx;

class Logger{
    public:
        Logger(){
            cout<<" Logger instance created "<<endl;
        }
        void log(){
            cout<<" add log "<<endl;
        }
};


void thread_func(int id){
    cout<<"Thread id "<<id<<" creating or accessing singleton "<<endl;
    auto instance = Singleton<Logger>::get_instance();
    instance.log();
}
int main(){
    std::vector<thread> threads;

    for(int i = 0 ; i <5 ; i++){
        threads.push_back(thread(thread_func , i));
    }

    for(int i = 0 ; i <5 ; i++){
        threads[i].join();
    }

    return 0;
}