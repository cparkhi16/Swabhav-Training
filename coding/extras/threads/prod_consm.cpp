#include <iostream>
#include <thread>
#include <mutex>
#include <condition_variable>
#include <chrono>

std::mutex mtx;                 // Mutex to protect shared resource
std::condition_variable cv;     // Condition variable for synchronization
bool resourceReady = false;     // Flag indicating whether the resource is ready
int resource = 0;               // Shared resource

void producer() {
    for (int i = 1; i <= 5; ++i) { // Producing 5 resources
        std::unique_lock<std::mutex> lock(mtx);

        // Wait until the consumer consumes the previous resource
        cv.wait(lock, [] { return !resourceReady; });

        // Produce a resource
        resource = i;
        std::cout << "Producer: Produced resource " << resource << "\n";
        resourceReady = true;

        // Notify the consumer
        cv.notify_one();
    }
}

void consumer() {
    for (int i = 1; i <= 5; ++i) { // Consuming 5 resources
        std::unique_lock<std::mutex> lock(mtx);

        // Wait until the producer produces a new resource
        cv.wait(lock, [] { return resourceReady; });

        // Consume the resource
        std::cout << "Consumer: Consumed resource " << resource << "\n";
        resourceReady = false;

        // Notify the producer
        cv.notify_one();
    }
}

int main() {
    std::thread th1(producer);
    std::thread th2(consumer);

    // Join threads
    th1.join();
    th2.join();

    return 0;
}
