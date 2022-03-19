/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 * @return {string}
 */
var tree2str = function (root) {
    let result = "";
    if (!root) {
        return result;
    }
    result = result + root.val;
    const left = tree2str(root.left);
    const right = tree2str(root.right);
    if (left) {
        result = result + '(' + left + ')';
    } else if (right) {
        result += '()';
    }
    if (right) {
        result = result + '(' + right + ')';
    }
    return result;
};