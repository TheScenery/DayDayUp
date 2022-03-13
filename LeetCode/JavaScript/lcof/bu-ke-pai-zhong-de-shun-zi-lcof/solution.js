/**
 * @param {number[]} nums
 * @return {boolean}
 */
var isStraight = function (nums) {
    nums = nums.sort((a, b) => a - b);
    var specialCard = 0;
    var pre = -1;
    for (var i = 0; i < nums.length;) {
        if (nums[i] === 0) {
            specialCard++;
            i++;
        } else {
            if (pre === -1) {
                pre = nums[i];
                i++;
            } else {
                if (nums[i] != pre + 1) {
                    if (specialCard > 0) {
                        specialCard--;
                    } else {
                        return false;
                    }
                } else {
                    i++;
                }
                pre += 1;
            }
        }
    }
    return true;
};

var r = isStraight([1, 2, 3, 4, 5]);
console.log(r);

var r = isStraight([0, 0, 1, 2, 5]);
console.log(r);

var r = isStraight([1, 2, 3, 5, 6]);
console.log(r);

var r = isStraight([0, 0, 2, 2, 5]);
console.log(r);