/**
 * @param {number[]} nums
 * @param {number} key
 * @return {number}
 */
var mostFrequent = function (nums, key) {
    let targetCountMap = {};
    let i = 0;
    let maxCount = 0;
    let ans;
    for (let i = 0; i < nums.length - 1; i++) {
        if (nums[i] === key) {
            let target = nums[i + 1];
            if (!targetCountMap[target]) {
                targetCountMap[target] = 0;
            }
            targetCountMap[target]++;
            if (targetCountMap[target] > maxCount) {
                maxCount = targetCountMap[target];
                ans = target;
            }
        }
    }
    return ans;
};