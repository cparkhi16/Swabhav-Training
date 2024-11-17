#include <iostream>
#include <thread>
#include <mutex>

std::mutex mtx; // mutex for critical section

void print_number(int id) {
    std::lock_guard<std::mutex> lock(mtx); // lock the mutex
    std::cout << "Thread " << id << " is in the critical section.\n";
    std::this_thread::sleep_for(std::chrono::seconds(1));
    std::cout << "Thread " << id << " is leaving the critical section.\n";
}

int main() {
    std::thread t1(print_number, 1);
    std::thread t2(print_number, 2);

    t1.join();
    t2.join();

    return 0;
}
