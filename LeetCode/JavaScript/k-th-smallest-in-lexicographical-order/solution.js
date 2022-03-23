/**
 * @param {number} n
 * @param {number} k
 * @return {number}
 */
var findKthNumber = function (n, k) {
    const t = [];
    for (let i = 1; i <= n; i++) {
        t.push(i);
    }
    t.sort();
    return t[k - 1];
};


function getCount(p, n) {
    let cur = p;
    let next = p + 1;
    let count = 0;
    while (cur <= n) {
        count += Math.min(n + 1, next) - cur;
        cur *= 10;
        next *= 10;
    }
    return count;
}

var findKthNumber = function (n, k) {
    let r = 1;
    let p = 1;

    while (r < k) {
        const count = getCount(p, n);
        if (r + count > k) {
            p *= 10;
            r++;
        } else {
            p++;
            r += count;
        }
    }

    return p;
};