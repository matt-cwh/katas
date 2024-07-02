import unittest
from palindromic_str_finder import *


class TestPalindromicStrFinder(unittest.TestCase):

    def test_find_longest_palindromic_str(self):
        inputStr = "findtheehtlongesttsegnolsubstringwhichhcihwisapalindrome"
        expected = "longesttsegnol"
        expectedLength = len(expected)

        (longest_palindromic_str, length) = find_longest_palindromic_str(inputStr)

        self.assertEqual(longest_palindromic_str, expected)
        self.assertEqual(length, expectedLength)

    def test_find_longest_palindromic_str_odd(self):
        inputStr = "findtheehtlongestetsegnolsubstringwhichhcihwisapalindrome"
        expected = "longestetsegnol"
        expectedLength = len(expected)

        (longest_palindromic_str, length) = find_longest_palindromic_str(inputStr)

        self.assertEqual(longest_palindromic_str, expected)
        self.assertEqual(length, expectedLength)

if __name__ == '__main__':
    unittest.main()