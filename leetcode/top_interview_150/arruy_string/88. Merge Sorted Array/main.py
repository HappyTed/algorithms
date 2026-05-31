from typing import List


class Solution:
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        """
        Do not return anything, modify nums1 in-place instead.
        """
        nums1[m:] = nums2[:]
        nums1.sort()


tc_1 = ([1, 2, 3, 0, 0, 0], 3, [2, 5, 6], 3, [1, 2, 2, 3, 5, 6])

if __name__ == "__main__":
    Solution().merge(tc_1[0], tc_1[1], tc_1[2], tc_1[3])
    print(tc_1[0])
    print(tc_1[4])
    assert tc_1[0] == tc_1[4]
