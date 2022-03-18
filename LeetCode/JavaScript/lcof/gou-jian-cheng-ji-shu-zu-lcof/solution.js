/**
 * @param {number[]} a
 * @return {number[]}
 */
var constructArr = function (a) {
    const before = [];
    const after = [];
    for (let i = 0; i < a.length; i++) {
        if (i > 0) {
            before[i] = before[i - 1] * a[i];
        } else {
            before[i] = a[i];
        }
    }
    for (let i = a.length - 1; i >= 0; i--) {
        if (i === a.length - 1) {
            after[i] = a[i];
        } else {
            after[i] = a[i] * after[i + 1];
        }
    }
    const result = [];
    for (let i = 0; i < a.length; i++) {
        if (i === 0) {
            result[i] = after[i + 1];
        } else if (i === a.length - 1) {
            result[i] = before[i - 1];
        } else {
            result[i] = before[i - 1] * after[i + 1];
        }
    }
    return result;
};