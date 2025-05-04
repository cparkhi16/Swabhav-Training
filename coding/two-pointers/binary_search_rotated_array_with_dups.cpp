#include <iostream>
#include <vector>

int searchInRotatedArray(std::vector<int>& nums, int target) {
    int left = 0, right = nums.size() - 1;

    while (left <= right) {
        int mid = left + (right - left) / 2;

        if (nums[mid] == target) return mid; // Found target

        // Handling duplicates: Shrink search space
        if (nums[left] == nums[mid] && nums[mid] == nums[right]) {
            left++;
            right--;
            continue;
        }

        // Check if left half is sorted
        if (nums[left] <= nums[mid]) {
            if (nums[left] <= target && target < nums[mid]) {
                right = mid - 1; // Search left
            } else {
                left = mid + 1; // Search right
            }
        }
        // Right half is sorted
        else {
            if (nums[mid] < target && target <= nums[right]) {
                left = mid + 1; // Search right
            } else {
                right = mid - 1; // Search left
            }
        }
    }
    return -1; // Target not found
}

int main() {
    std::vector<int> arr = {2, 5, 6, 0, 0, 1, 2};
    int target = 0;

    int index = searchInRotatedArray(arr, target);
    if (index != -1)
        std::cout << "Element found at index: " << index << std::endl;
    else
        std::cout << "Element not found." << std::endl;

    return 0;
}
