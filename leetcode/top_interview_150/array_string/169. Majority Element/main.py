from typing import List
from collections import Counter


class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        freq = dict()
        threshold = len(nums) // 2
        for n in nums:
            freq[n] = freq.get(n, 0) + 1
            if freq[n] > threshold:
                return n
        return nums[0]

    def majorityElementII(self, nums: List[int]) -> int:
        return Counter(nums).most_common(1)[0][0]


if __name__ == "__main__":
    # print(Solution().)
    pass
