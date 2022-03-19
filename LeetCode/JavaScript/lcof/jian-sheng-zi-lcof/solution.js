/**
 * @param {number} n
 * @return {number}
 */
var cuttingRope = function (n) {
    if (n <= 3) {
        return n - 1;
    }
    const mod = n % 3;
    if (mod === 0) {
        return Math.pow(3, n / 3);
    } else if (mod === 1) {
        return Math.pow(3, (n - 1) / 3 - 1) * 4;
    } else if (mod === 2) {
        return Math.pow(3, (n - 2) / 3) * 2;
    }
};