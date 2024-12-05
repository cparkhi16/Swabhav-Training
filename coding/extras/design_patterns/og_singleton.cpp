#include <iostream>
#include <memory>

class MyClass {
public:
    // Delete copy constructor and assignment operator to prevent cloning
    MyClass(const MyClass& obj) = delete;
    MyClass& operator=(const MyClass& obj) = delete;
    // Static method to access the single instance
    static MyClass& getInstance() {
        static MyClass instance;  // Guaranteed to be created only once
        return instance;
    }

    void display() const {
        std::cout << "MyClass instance in action!" << std::endl;
    }

private:
    // Private constructor prevents direct instantiation
    MyClass() { std::cout << "MyClass instance created\n"; }
};

int main() {
    // Access the singleton instance
    MyClass& instance1 = MyClass::getInstance();
    instance1.display();
    
     MyClass& instance2 = MyClass::getInstance();
   instance2.display();
    // Uncommenting this will result in a compile-time error:
    // MyClass anotherInstance; 

    return 0;
}
