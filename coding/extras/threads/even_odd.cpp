#include <iostream>
#include <thread>
#include <mutex>
#include <condition_variable>

std::mutex mtx;
std::condition_variable cv;
bool evenTurn = false; // Flag to track turn

void printEven() {
    for (int i = 2; i <= 10; i += 2) {
        std::unique_lock<std::mutex> lock(mtx);
        cv.wait(lock, [] { return evenTurn; }); // Wait until it's even's turn
        std::cout << i << " ";
        evenTurn = false; // Switch to odd turn
        cv.notify_one();  // Notify odd thread
    }
}

void printOdd() {
    for (int i = 1; i <= 9; i += 2) {
        std::unique_lock<std::mutex> lock(mtx);
        cv.wait(lock, [] { return !evenTurn; }); // Wait until it's odd's turn
        std::cout << i << " ";
        evenTurn = true; // Switch to even turn
        cv.notify_one(); // Notify even thread
    }
}

int main() {
    std::thread t1(printOdd);
    std::thread t2(printEven);

    t1.join();
    t2.join();

    return 0;
}
