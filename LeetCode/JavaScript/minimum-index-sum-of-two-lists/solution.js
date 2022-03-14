/**
 * @param {string[]} list1
 * @param {string[]} list2
 * @return {string[]}
 */
var findRestaurant = function (list1, list2) {
    let m = {};
    for (let i = 0; i < list1.length; i++) {
        m[list1[i]] = i;
    }
    let minIndexSum = list1.length + list2.length;
    let res = [];
    for (let i = 0; i < list2.length; i++) {
        let restaurant = list2[i];
        if (m[restaurant] !== undefined) {
            let indexSum = i + m[restaurant];
            if (indexSum < minIndexSum) {
                res = [restaurant];
                minIndexSum = indexSum;
            } else if (indexSum === minIndexSum) {
                res.push(restaurant);
            }
        }
    }
    return res;
};