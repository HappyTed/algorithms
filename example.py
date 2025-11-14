import re


txt = 'aa aba abba abbba abca abea'

print(re.sub("ab*a", "!", txt))