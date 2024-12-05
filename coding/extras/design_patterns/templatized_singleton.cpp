#include <iostream>
#include <memory>
using namespace std;

template <typename T>
class Singleton {
public:
    static T& getInstance() {
        if (!ins) {
            ins.reset(new T()); // Directly call new T() instead of std::make_unique
        }
        return *ins;
    }

private:
    static std::unique_ptr<T> ins;
};

template <typename T>
std::unique_ptr<T> Singleton<T>::ins = nullptr;

class Logger : public Singleton<Logger> {
    // Declare Singleton<Logger> as a friend to allow access to the private constructor
    friend class Singleton<Logger>;

private:
    Logger() {
        cout << "Logger created" << endl;
    }

public:
    void showMsg() {
        cout << "Message from Logger" << endl;
    }
};

int main() {
    auto& instance1 = Singleton<Logger>::getInstance(); // Singleton instance
    instance1.showMsg();

    auto& instance2 = Singleton<Logger>::getInstance(); // Same instance
    instance2.showMsg();

    return 0;
}
