#include <iostream>
#include <vector>
#include <cmath> // for std::abs

int main() {
    std::vector<int> nums = {-9, -6, 4}; //given a sorted array , give output array in ascending order (with squares of each no)
    int n = nums.size();
    std::vector<int> res(n);
    int L = 0;
    int R = n - 1;

    // Fill the result vector in reverse order
    for (int index = n - 1; index >= 0; --index) {
        if (std::abs(nums[L]) > std::abs(nums[R])) {
            res[index] = nums[L] * nums[L];
            L++;
        } else {
            res[index] = nums[R] * nums[R];
            R--;
        }
    }

    // Print the result
    for (int num : res) {
        std::cout << num << " ";
    }
    std::cout << std::endl;

    return 0;
}
