function convertNum(mapping, num) {
    if (num === 0) {
        return mapping[0];
    }
    let res = 0;
    let base = 0;
    while (num > 0) {
        let i = num % 10;
        let v = mapping[i];
        res = res + v * Math.pow(10, base);
        base++;
        num = Math.floor(num / 10);
    }
    return res;
}


/**
 * @param {number[]} mapping
 * @param {number[]} nums
 * @return {number[]}
 */
var sortJumbled = function (mapping, nums) {

    let mappedNumbers = nums.map((num, i) => ({ value: num, mapped: convertNum(mapping, num), index: i }))
    mappedNumbers.sort((a, b) => {
        if (a.mapped < b.mapped) {
            return -1;
        } else if (a.mapped > b.mapped) {
            return 1;
        }
        return a.index - b.index;
    })
    return mappedNumbers.map(a => a.value);
};

console.log(sortJumbled([8, 9, 4, 0, 2, 1, 3, 5, 7, 6], [991, 338, 38]));