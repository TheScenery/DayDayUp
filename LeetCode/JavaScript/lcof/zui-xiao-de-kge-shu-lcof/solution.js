/**
 * @param {number[]} arr
 * @param {number} k
 * @return {number[]}
 */

function quickSort(arr, start, end) {
    if (start >= end) {
        return;
    }
    var key = arr[start];
    var i = start;
    var j = end;
    while (i < j) {
        while (j > i && arr[j] >= key) {
            j--;
        }
        while (i < j && arr[i] <= key) {
            i++;
        }
        var temp = arr[i];
        arr[i] = arr[j];
        arr[j] = temp;
    }
    arr[start] = arr[i];
    arr[i] = key;
    quickSort(arr, start, i - 1);
    quickSort(arr, i + 1, end);
}
var getLeastNumbers = function (arr, k) {
    quickSort(arr, 0, arr.length - 1);
    return arr.slice(0, k);
};
