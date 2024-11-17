// C++ code for the above approach:
#include <bits/stdc++.h>
using namespace std;

vector<int> majorityElement(vector<int>& nums, int k)
{
	int n = nums.size();

	// Initialize an array of pairs to store k-1
	// candidates and their counts
	vector<pair<int, int> > candidates(k - 1);

	// Initialize candidate elements
	// and their counts to 0
	for (int i = 0; i < k - 1; i++) {
		candidates[i] = { 0, 0 };
	}

	// Step 1: Scanning the Array
	for (int num : nums) {
		bool found = false;
		for (int j = 0; j < k - 1; j++) {

			// If the element exists in
			// candidates
			if (candidates[j].first == num) {

				// Increment its count
				candidates[j].second++;
				found = true;
				break;
			}
		}
		// If the element is not in candidates
		if (!found) {
			for (int j = 0; j < k - 1; j++) {

				// Find an empty slot
				if (candidates[j].second == 0) {

					// Insert as a new candidate
					candidates[j] = { num, 1 };
					found = true;
					break;
				}
			}
		}

		// If all slots are occupied
		if (!found) {

			// Decrement counts
			// of all candidates
			for (int j = 0; j < k - 1; j++) {
				candidates[j].second--;
			}
		}
	}

	// Initialize a vector to store
	// the final results
	vector<int> ans;

	// Step 2: Verification
	for (int i = 0; i < k - 1; i++) {
		int count = 0;
		for (int j = 0; j < n; j++) {
			if (nums[j] == candidates[i].first) {
				count++;
			}
		}

		// If the count is greater than n/k
		if (count > n / k) {

			// Add the candidate
			// to the result
			ans.push_back(candidates[i].first);
		}
	}

	return ans;
}

// Drivers code
int main()
{
	vector<int> nums = { 3, 2, 3,3,3};
	int k = 3;
	vector<int> result = majorityElement(nums, k);

	cout << "Elements occurring more than "
		<< nums.size() / k << " times: ";
	for (int num : result) {
		cout << num << " ";
	}
	cout << endl;

	return 0;
}
