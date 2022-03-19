/**
 * @param {number} target
 * @return {number[][]}
 */
var findContinuousSequence = function (target) {
    let l = 1;
    let r = 2;
    const result = [];
    while (l < r) {
        const sum = (r + l) * (r - l + 1) / 2;
        if (sum < target) {
            r++;
        } else if (sum > target) {
            l++;
        } else {
            const temp = [];
            for (let i = l; i <= r; i++) {
                temp.push(i);
            }
            result.push(temp);
            l++;
        }
    }
    return result;
};

console.log(findContinuousSequence(9));