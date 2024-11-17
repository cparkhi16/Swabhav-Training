#include<iostream>
#include <mutex>
#include <condition_variable>
#include<thread>
using namespace std;
class Semaphore{
    public:
        Semaphore(int c) : count(c){}
        void acquire(){
            std::unique_lock<mutex> lock(mtx);
            //cv.wait(lock , [this](){ return count > 0 ;});
            cv.wait(lock, [this]() { return count > 0; });
            count++;
        }

        void release(){
            std::unique_lock<mutex> lock(mtx);
            count--;
            cv.notify_one();
        }
    private:
        std::mutex mtx;
        std::condition_variable cv;
        int count;
};

void th_func(int id, Semaphore& sem){
    sem.acquire();
    cout<<" thread id "<<id<<" entering critcal section "<<endl;
    std::this_thread::sleep_for(chrono::seconds(1));
    cout<<" Threead id "<<id<<" leaving criticla section "<<endl;
    sem.release();
}

int main(){
    Semaphore sem(3);

    thread t1(th_func, 1 , std::ref(sem));
    thread t2(th_func, 2, std::ref(sem));
    thread t3(th_func, 3,  std::ref(sem));
    thread t4(th_func, 4,  std::ref(sem));

    t1.join();
    t2.join();
    t3.join();
    t4.join();
}