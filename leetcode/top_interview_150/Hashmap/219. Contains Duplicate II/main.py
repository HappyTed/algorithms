from typing import List


class Solution:
    def containsNearbyDuplicate(self, nums: List[int], k: int) -> bool:
        for i in range(len(nums)):
            target_1 = i - k  # j; abs(i-j) = k
            target_2 = i + k

            if target_1 > 0:
                if target_1 < len(nums):
                    if nums[target_1] == nums[i] and i - target_1 <= k:
                        return True

            if target_2 > 0:
                if target_2 < len(nums):
                    if nums[target_2] == nums[i] and i - target_2 <= k:
                        return True

        return False


if __name__ == "__main__":
    print(Solution().containsNearbyDuplicate([1, 2, 3, 1], 3))
