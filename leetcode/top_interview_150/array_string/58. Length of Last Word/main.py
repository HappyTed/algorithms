class Solution:
    def lengthOfLastWord(self, s: str) -> int:
        counter = 0
        slow = 0
        is_substring = False
        for faster in range(len(s)):
            if s[faster] != " " and not is_substring:
                is_substring = True
                slow = faster
            elif (s[faster] == " ") and is_substring:
                counter = faster - slow
                is_substring = False
        if is_substring:
            counter = faster + 1 - slow
        return counter


if __name__ == "__main__":
    print(Solution().lengthOfLastWord("luffy is still joyboy"))
    print(Solution().lengthOfLastWord("Hello World"))
    print(Solution().lengthOfLastWord("   fly me   to   the moon  "))
