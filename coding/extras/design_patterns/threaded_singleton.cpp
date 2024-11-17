#include <iostream>
#include <thread>
#include <vector>
#include <mutex>
#include <memory>

// Singleton class template with thread-safety and lazy initialization
template <typename T>
class Singleton {
public:
    static T& Instance() {
        if (!instance) {
            std::lock_guard<std::mutex> lock(mutex);
            if (!instance) {
                instance = std::make_unique<T>();
            }
        }
        return *instance;
    }

private:
    static std::unique_ptr<T> instance;
    static std::mutex mutex;
};

// Definition of the static member variables
template <typename T>
std::unique_ptr<T> Singleton<T>::instance = nullptr;

template <typename T>
std::mutex Singleton<T>::mutex;

// Test class to use with Singleton
class MyClass {
public:
    MyClass() {
        std::cout << "MyClass instance created\n";
    }

    void showMessage() {
        std::cout << "Hello from MyClass Singleton!\n";
    }
};

// Driver function for threaded Singleton
void threadFunction(int threadID) {
    std::cout << "Thread " << threadID << " accessing Singleton instance...\n";
    MyClass& instance = Singleton<MyClass>::Instance();
    instance.showMessage();
}

int main() {
    // Create a vector of threads
    std::vector<std::thread> threads;

    // Launch 5 threads to demonstrate thread-safe Singleton
    for (int i = 1; i <= 5; ++i) {
        threads.push_back(std::thread(threadFunction, i));
    }

    // Join all threads
    for (auto& t : threads) {
        t.join();
    }

    return 0;
}
