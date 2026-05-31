from typing import List


class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        if len(nums) < 1:
            return -1
        left, right = 0, len(nums) - 1
        while left <= right:
            if nums[right] == val:
                right -= 1
                continue
            if nums[left] == val:
                nums[left], nums[right] = nums[right], nums[left]
            left += 1
        return left


if __name__ == "__main__":
    # print(Solution().)
    pass
