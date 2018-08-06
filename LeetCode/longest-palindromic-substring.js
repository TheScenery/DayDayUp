/*
    Title:5. Longest Palindromic Substring
    Description:
    Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.
    Example:
    Input: 
    "babad"
    Output: 
    "bab"
    Note:
    "aba" is also a valid answer.
    Input: 
    "cbbd"
    Output: 
    "bb"

    Language: Javascript
*/

/**
 * @param {string} s
 * @return {string}
 */
var longestPalindromeDP = function (s) {
    let longestLength = 1;
    let longestMin = 0;
    for (let i = 0; i < s.length; i++) {
        let iMin = i - 1;
        let iMax = i + 1;
        while (iMin >= 0 && iMax <= s.length) {
            if (s[iMin] === s[iMax]) {
                iMax++;
                iMin--;
            } else {
                break;
            }
        }
        let tempLength = iMax - iMin - 1;
        if (tempLength > longestLength) {
            longestLength = tempLength;
            longestMin = iMin + 1;
        }
        if (s[i] === s[i + 1]) {
            if (longestLength < 2) {
                longestLength = 2;
                longestMin = i;
            }
            iMin = i - 1;
            iMax = i + 1 + 1;
            while (iMin >= 0 && iMax <= s.length) {
                if (s[iMin] === s[iMax]) {
                    iMax++;
                    iMin--;
                } else {
                    break;
                }
            }
            tempLength = iMax - iMin - 1;
            if (tempLength > longestLength) {
                longestLength = tempLength;
                longestMin = iMin + 1;
            }
        }
    }
    return s.substr(longestMin, longestLength);
};

console.log(longestPalindromeDP('babad'));