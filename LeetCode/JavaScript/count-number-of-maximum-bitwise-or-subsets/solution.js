/**
 * @param {number[]} nums
 * @return {number}
 */
var countMaxOrSubsets = function (nums) {
    let maxOr = 0;
    for (let i = 0; i < nums.length; i++) {
        maxOr |= nums[i];
    }
    let count = 0;
    for (let i = 0; i < (1 << nums.length); i++) {
        let orValue = 0;
        for (let j = 0; j < nums.length; j++) {
            if ((i >> j) & 1 === 1) {
                orValue |= nums[j];
            }
        }
        if (orValue === maxOr) {
            count++;
        }
    }
    return count;
};