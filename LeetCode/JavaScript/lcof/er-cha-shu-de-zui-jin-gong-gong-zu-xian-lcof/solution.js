/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @param {TreeNode} p
 * @param {TreeNode} q
 * @return {TreeNode}
 */
var lowestCommonAncestor = function (root, p, q) {
    let ans;
    const contains = (node) => {
        if (node === null) {
            return false;
        }
        const leftContains = contains(node.left);
        const rightContains = contains(node.right);
        if ((leftContains && rightContains) || ((node.val === p.val || node.val === q.val) && (leftContains || rightContains))) {
            ans = node;
        }
        return leftContains || rightContains || node.val === p.val || node.val === q.val;
    }
    contains(root);
    return ans;
};