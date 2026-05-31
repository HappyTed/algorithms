from typing import List


class Solution:
    def summaryRanges(self, nums: List[int]) -> List[str]:

        outputs = []
        if not nums:
            return outputs

        start_idx = 0
        for i in range(1, len(nums) + 1):
            if i == len(nums) or nums[i] != nums[i - 1] + 1:
                if start_idx == i - 1:
                    outputs.append(f"{nums[start_idx]}")
                else:
                    outputs.append(f"{nums[start_idx]}->{nums[i - 1]}")
                start_idx = i
                continue
        return outputs


if __name__ == "__main__":
    print(Solution().summaryRanges([0, 1, 2, 4, 5, 7]))
    print(Solution().summaryRanges([0, 2, 3, 4, 6, 8, 9]))
