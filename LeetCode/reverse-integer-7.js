/*
    Title:7. Reverse Integer
    Description:
    Given a 32-bit signed integer, reverse digits of an integer.
    Example:
    Input: 123
    Output: 321
    Input: -123
    Output: -321
    Input: 120
    Output: 21

    Language: Javascript
*/

/**
 * @param {number} x
 * @return {number}
 */
var reverse = function (x) {
    var xString = x.toString();
    var xArray = xString.split('');
    var sign = xArray[0] === '-';
    if (sign) {
        xArray = xArray.slice(1);
    }
    var reverseArray = xArray.reverse();
    if (sign) {
        reverseArray.unshift('-');
    }
    var result = parseInt(reverseArray.join(''));
    if (result > 2147483647 || result < -2147483648) {
        return 0
    }
    return result;
};

function testMain() {
    console.log(reverse(-123))
}

testMain()