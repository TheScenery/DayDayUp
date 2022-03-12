function findIndex(arr, key, start, end) {
    if (start > end) {
        return end + 1;
    }
    var mid = Math.floor(start + (end - start) / 2);
    if (arr[mid] < key) {
        return findIndex(arr, key, mid + 1, end);
    }
    if (arr[mid] > key) {
        return findIndex(arr, key, start, mid - 1);
    }
    return mid
}
/**
 * initialize your data structure here.
 */
var MedianFinder = function () {
    this.data = [];
};

/** 
 * @param {number} num
 * @return {void}
 */
MedianFinder.prototype.addNum = function (num) {
    var index = findIndex(this.data, num, 0, this.data.length - 1);
    this.data.splice(index, 0, num);
};

/**
 * @return {number}
 */
MedianFinder.prototype.findMedian = function () {
    var l = this.data.length;
    if (l===0) {
        return null;
    }
    if ((l & 1) == 0) {
        return (this.data[l / 2] + this.data[l / 2 - 1]) / 2;
    }
    return this.data[Math.floor(l / 2)];
};

var obj = new MedianFinder();
obj.addNum(-1);
obj.addNum(-2);
obj.addNum(-3);
obj.addNum(-4);
obj.addNum(-5);
console.log(obj.findMedian());
