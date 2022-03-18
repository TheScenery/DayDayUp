/**
 * @param {number[]} nums
 * @return {number}
 */
var majorityElement = function (nums) {
    const count = nums.length / 2;
    const countMap = {};
    for (let i = 0; i < nums.length; i++) {
        const num = nums[i];
        if (!countMap[num]) {
            countMap[num] = 0;
        }
        countMap[num]++;
        if (countMap[num] > count) {
            return num;
        }
    }
    return null;
};