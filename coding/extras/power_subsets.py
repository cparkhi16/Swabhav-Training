from typing import List

class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        n = len(nums)
        ans, sol = [], []

        def backtrack(i):
            if i == n:
                ans.append(sol[:])
                return

            # for x in nums:
            #     if x not in sol:
            #         sol.append(x)
            #         backtrack()
            #         sol.pop()
            sol.append(nums[i])
            backtrack(i+1)
            sol.pop()
            backtrack(i+1)

        backtrack(0)
        return ans

# Create an instance of the Solution class
solution = Solution()
# subsequences of the given array 
# Invoke the permute method with [1, 2, 3] as the input list
result = solution.permute([1, 2, 3])

# Print the result
print(result)
