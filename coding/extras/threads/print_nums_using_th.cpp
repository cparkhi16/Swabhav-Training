#include <iostream>
#include <thread>
#include <vector>
using namespace std;
void print_range(int start, int end) {
    for (int i = start; i <= end; ++i) {
        std::cout << i << " ";
    }
    std::cout << std::endl;
}

int main() {
    int total_numbers = 100;
    int num_threads = 4;  // Number of threads to use
    
    // Calculate range size for each thread
    int range_size = total_numbers / num_threads;
    std::vector<std::thread> threads;

    for (int i = 0; i < num_threads; ++i) {
        int start = i * range_size + 1;
        int end = (i + 1) * range_size;
         
        std::cout<<" start "<<start<<" end "<<end<<" thread id "<<i<<endl;
        // Ensure the last thread covers any remaining numbers
        if (i == num_threads - 1) {
            end = total_numbers;
        }

        //threads.push_back(std::thread(print_range, start, end));
    }

    // Join all threads
    // for (auto& t : threads) {
    //     t.join();
    // }

    return 0;
}
