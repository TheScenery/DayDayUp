/*
    Title:6. ZigZag Conversion
    Description:
    The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)
    Example:
    Input: s = "PAYPALISHIRING", numRows = 3
    Output: "PAHNAPLSIIGYIR"
    Explanation: 
    P   A   H   N
    A P L S I I G
    Y   I   R   

    Language: Javascript
*/

/**
 * @param {string} s
 * @param {number} numRows
 * @return {string}
 */
const convert = function (s, numRows) {
    const result = new Array(numRows);
    let resultIndex = 0;
    let step = 1;
    for (let i = 0; i < s.length; i++) {
        let row = result[resultIndex];
        if (!row) {
            result[resultIndex] = row = [];
        }
        row.push(s[i]);
        if (resultIndex === 0) {
            step = 1;
        } else if (resultIndex === numRows - 1) {
            step = -1;
        }
        resultIndex += step;
    }
    return result.map(r => r.join('')).join('');
};

const testMain = () => {
    console.log(convert('PAYPALISHIRING', 3));
}

testMain();