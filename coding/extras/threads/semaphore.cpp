#include <iostream>
#include <thread>
#include <mutex>
#include <condition_variable>

class Semaphore {
public:
    Semaphore(int count) : count(count) {}

    void acquire() {
        std::unique_lock<std::mutex> lock(mtx);
        cv.wait(lock, [this]() { return count > 0; }); // wait until the count is positive
        --count;
    }

    void release() {
        std::lock_guard<std::mutex> lock(mtx);
        ++count;
        cv.notify_one(); // notify one waiting thread
    }

private:
    int count;
    std::mutex mtx;
    std::condition_variable cv;
};

void use_resource(Semaphore& sem, int id) {
    sem.acquire();
    std::cout << "Thread " << id << " is using the resource.\n";
    std::this_thread::sleep_for(std::chrono::seconds(1));
    std::cout << "Thread " << id << " is releasing the resource.\n";
    sem.release();
}

int main() {
    Semaphore sem(3); // semaphore with 3 available resources

    std::thread t1(use_resource, std::ref(sem), 1);
    std::thread t2(use_resource, std::ref(sem), 2);
    std::thread t3(use_resource, std::ref(sem), 3);
    std::thread t4(use_resource, std::ref(sem), 4); // This thread will wait for a resource

    t1.join();
    t2.join();
    t3.join();
    t4.join();

    return 0;
}
