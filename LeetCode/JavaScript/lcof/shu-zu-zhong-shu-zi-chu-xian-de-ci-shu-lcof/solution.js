/**
 * @param {number[]} nums
 * @return {number[]}
 */
var singleNumbers = function (nums) {
    let r = 0;
    nums.forEach(n => {
        r = r ^ n;
    })
    let d = 1;
    while ((r & d) === 0) {
        d = d << 1;
    }
    let a = 0;
    let b = 0;
    nums.forEach(n => {
        if ((n & d) !== 0) {
            a ^= n;
        } else {
            b ^= n;
        }
    })
    return [a, b];
};