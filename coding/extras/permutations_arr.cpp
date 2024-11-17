#include <iostream>
#include <vector>
#include <string>
#include <algorithm>

// Template function to swap elements in a container
template <typename T>
void swap(std::vector<T>& vec, int i, int j) {
    T temp = vec[i];
    vec[i] = vec[j];
    vec[j] = temp;
}

// Template function to generate permutations
template <typename T>
void permute(std::vector<T>& vec, int start, int end) {
    if (start == end) {
        // Print the permutation
        for (const T& item : vec) {
            std::cout << item << " ";
        }
        std::cout << std::endl;
    } else {
        for (int i = start; i <= end; ++i) {
            swap(vec, start, i);
            permute(vec, start + 1, end);
            swap(vec, start, i); // Backtrack
        }
    }
}

int main() {
    // Example with integers
    std::vector<int> intVec = {1, 2, 3};
    std::cout << "Permutations of integers:" << std::endl;
    permute(intVec, 0, intVec.size() - 1);

    // Example with characters (string)
    std::vector<char> charVec = {'A', 'B', 'C'};
    std::cout << "Permutations of characters:" << std::endl;
    permute(charVec, 0, charVec.size() - 1);

    return 0;
}
