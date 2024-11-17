from typing import List

class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        n = len(nums)
        ans, sol = [], []

        def backtrack():
            if len(sol) == n:
                print(" got solution ",sol)
                ans.append(sol[:])
                return

            for x in nums:
                if x not in sol:
                    print(" appending in sol ",x)
                    sol.append(x)
                    backtrack()
                    print(" returned from backtrack , popping from ",sol)
                    sol.pop()

        backtrack()
        return ans

# Create an instance of the Solution class
solution = Solution()

# Invoke the permute method with [1, 2, 3] as the input list
result = solution.permute([1, 2, 3])

# Print the result
print(result)
