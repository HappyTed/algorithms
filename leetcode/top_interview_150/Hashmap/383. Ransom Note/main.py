class Solution(object):
    def canConstruct(self, ransomNote: str, magazine: str) -> str:
        """
        :type ransomNote: str
        :type magazine: str
        :rtype: bool
        """

        mapped = {}
        for el in magazine:
            if mapped.get(el):
                mapped[el] += 1
            else:
                mapped[el] = 1

        for let in ransomNote:
            if mapped.get(let, False) and mapped[let] > 0:
                mapped[let] -= 1
                continue
            else:
                return False
        return True


if __name__ == "__main__":
    print(Solution().canConstruct("a", "b"))
    print(Solution().canConstruct("aa", "ab"))
    print(Solution().canConstruct("aa", "aab"))
