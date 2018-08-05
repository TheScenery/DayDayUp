/*
    Title:4. Median of Two Sorted Arrays
    Description:
    There are two sorted arrays nums1 and nums2 of size m and n respectively.
    Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).
    You may assume nums1 and nums2 cannot be both empty.
    Example:
    Input: 
    nums1 = [1, 3]
    nums2 = [2]
    Output: 
    2.0
    Input: 
    nums1 = [1, 2]
    nums2 = [3, 4]
    Output: 
    2.5

    Language: Javascript
*/


// solution 1: Why it works? I dont know.
/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number}
 */

// const findMedianSortedArrays = (nums1, nums2) => {
//     const nums = nums1.concat(nums2).sort((n1, n2) => n1 - n2);
//     const medianIndex = (nums1.length + nums2.length) / 2;
//     if (medianIndex === parseInt(medianIndex)) {
//         return (nums[medianIndex] + nums[medianIndex - 1]) / 2;
//     } else {
//         return nums[Math.floor(medianIndex)];
//     }
// };


// solution 2:
/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number}
 */

const findMedianSortedArrays = (nums1, nums2) => {
    if (nums1.length > nums2.length) {
        const temp = nums1;
        nums1 = nums2;
        nums2 = temp;
    }

    const m = nums1.length;
    const n = nums2.length;
    let iMin = 0;
    let iMax = m;
    let i;
    let j;
    while (iMin <= iMax) {
        i = Math.floor((iMin + iMax) / 2);
        j = Math.floor((m + n + 1) / 2 - i);
        if (i < m && nums2[j - 1] > nums1[i]) {
            iMin += 1;
        } else if (i > 0 && nums1[i - 1] > nums2[j]) {
            iMax -= 1;
        } else {
            let maxLeft;
            if (i === 0) {
                maxLeft = nums2[j - 1];
            } else if (j === 0) {
                maxLeft = nums1[i - 1];
            } else {
                maxLeft = Math.max(nums2[j - 1], nums1[i - 1]);
            }

            if ((m + n) % 2) {
                return maxLeft;
            }

            let minRight;
            if (i === m) {
                minRight = nums2[j];
            } else if (j === n) {
                minRight = nums1[i];
            } else {
                minRight = Math.min(nums2[j], nums1[i]);
            }
            return (maxLeft + minRight) / 2;
        }
    }
};

function main() {
    console.log(findMedianSortedArrays([1, 2], [3, 4]));
}

main()