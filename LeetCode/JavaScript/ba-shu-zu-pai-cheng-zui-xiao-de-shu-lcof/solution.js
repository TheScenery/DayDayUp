/**
 * @param {number[]} nums
 * @return {string}
 */
var minNumber = function (nums) {
    nums = nums.sort((a, b) => {
        var sa = a.toString();
        var sb = b.toString();
        if (sa + sb < sb + sa) {
            return -1;
        } else {
            return 1;
        }
    })
    return nums.join("");
};


var r = minNumber([10, 2]);
console.log(r);

var r = minNumber([3, 30, 34, 5, 9]);
console.log(r);