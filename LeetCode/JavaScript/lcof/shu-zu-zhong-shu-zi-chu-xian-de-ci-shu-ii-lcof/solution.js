/**
 * @param {number[]} nums
 * @return {number}
 */
var singleNumber = function (nums) {
	let one = 0;
	let two = 0;
	nums.forEach((num) => {
		one = one ^ num & ~two;
		two = two ^ num & ~one;
	});

	return one;
};