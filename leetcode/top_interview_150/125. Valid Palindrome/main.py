import re

remove_non_alch_ch = lambda s: re.sub(r'[^a-zA-Z0-9]', '', s) 

class Solution:
    def isPalindrome(self, s: str) -> bool:
        s = s.lower()
        s = s.strip()
        s = remove_non_alch_ch(s)
        l = 0
        r = len(s) - 1
        while l < len(s) or r > 0:
            if s[l] != s[r]:
                return False
            l+=1
            r-=1
        return True



if __name__ == "__main__":
    print(Solution().isPalindrome("A man, a plan, a canal: Panama"))
    print(Solution().isPalindrome("race a car"))
    print(Solution().isPalindrome(" "))