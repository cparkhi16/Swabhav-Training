#include <iostream>
#include <vector>
#include <deque>
using namespace std;
std::vector<int> maxSlidingWindow(const std::vector<int>& nums, int k) {
    std::deque<int> deq;
    std::vector<int> result;

    for (int i = 0; i < nums.size(); ++i) {
        // Remove indices that are out of the current window
        if (!deq.empty() && deq.front() == i - k){
            std::cout <<" front popping out "<< deq.front()<<endl;
            deq.pop_front();
        }

        // Remove smaller elements in k range as they are useless
        while (!deq.empty() && nums[deq.back()] < nums[i]){
            std::cout <<" back popping out "<< deq.back()<<endl;
            deq.pop_back();
        }
        // Add current element at the back of the deque
        deq.push_back(i);
        cout<<" pushing index in deque "<<i<<endl;
        // Starting from the (k-1)th element, store the max in the result
        if (i >= k - 1){
            std::cout<<" push ans in res -------------------->>"<<nums[deq.front()]<<"  i is"<<i<<endl;
            result.push_back(nums[deq.front()]);
        }
    }

    return result;
}

int main() {
    std::vector<int> nums = {7, 3, -1, -3, 5, 3, 6, 7};
    int k = 3;

    std::vector<int> result = maxSlidingWindow(nums, k);

    // Print the result
    for (int maxElem : result)
        std::cout << maxElem << " ";

    return 0;
}
