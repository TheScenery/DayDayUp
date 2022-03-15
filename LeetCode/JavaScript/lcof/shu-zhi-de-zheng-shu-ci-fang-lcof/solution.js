/**
 * @param {number} x
 * @param {number} n
 * @return {number}
 */
var myPow = function (x, n) {
    if (n === 0) {
        return 1
    }
    let res = 1;
    if (n < 0) {
        x = 1 / x;
        n = -n;
    }
    while (n > 0) {
        if (n & 1 === 1) {
            res *= x;
        }
        x *= x;
        n = Math.floor(n / 2);
    }
    return res;
};

console.log(myPow(2, -2147483648));