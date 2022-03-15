/**
 * @param {number[]} postorder
 * @return {boolean}
 */
var verifyPostorder = function (postorder) {
    if (postorder.length === 0) {
        return true;
    }
    let index = -1;
    const end = postorder[postorder.length - 1];
    for (let i = postorder.length - 1; i >= 0; i--) {
        if (postorder[i] < end) {
            index = i;
            break;
        }
    }
    for (let i = 0; i < index; i++) {
        if (postorder[i] > end) {
            return false;
        }
    }
    return verifyPostorder(postorder.slice(0, index + 1)) && verifyPostorder(postorder.slice(index + 1, postorder.length - 1));
};