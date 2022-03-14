/**
 * @param {number} n
 * @return {number}
 */
 var sumNums = function(n) {
    return n == 1 ? 1 : n + sumNums(n-1);
};