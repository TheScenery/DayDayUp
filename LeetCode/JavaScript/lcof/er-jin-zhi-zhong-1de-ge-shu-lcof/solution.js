/**
 * @param {number} n - a positive integer
 * @return {number}
 */
var hammingWeight = function (n) {
    let count = 0;
    while (n > 0) {
        n = n & (n - 1);
        n = n >>> 0;
        count++;
    }
    return count;
};

console.log(hammingWeight(0b11111111111111111111111111111101));