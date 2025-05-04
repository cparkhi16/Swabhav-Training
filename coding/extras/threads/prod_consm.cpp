#include <iostream>
#include <mutex>
#include <condition_variable>
#include <vector>
#include <bits/stdc++.h>
std::mutex mu;
std::condition_variable cond;
std::vector<int> buffer;
const unsigned int maxBufferSize = 50;

void producer(int val) {
    while (val) {
        std::unique_lock<std::mutex> locker(mu);
        cond.wait(locker, []() { return buffer.size() < maxBufferSize; });

        buffer.push_back(val);
        std::cout << "Produced: " << val << std::endl;

        val--;
        locker.unlock();
        cond.notify_one();
    }
}

void consumer() {
    while (true) {
        std::unique_lock<std::mutex> locker(mu);
        cond.wait(locker, []() { return buffer.size() > 0; });

        int val = buffer.back();
        buffer.pop_back();
        std::cout << "Consumed: " << val << std::endl;

        locker.unlock();
        cond.notify_one();
    }
}

int main() {
    std::thread prodThread(producer, 100);
    std::thread consThread(consumer);

    prodThread.join();
    consThread.join();

    return 0;
}
