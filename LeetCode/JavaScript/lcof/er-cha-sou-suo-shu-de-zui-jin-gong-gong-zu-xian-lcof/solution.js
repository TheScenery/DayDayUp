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
    if (p.val > q.val) {
        [p, q] = [q, p];
    }
    if (p.val > root.val) {
        return lowestCommonAncestor(root.right, p, q);
    }
    if (q.val < root.val) {
        return lowestCommonAncestor(root.left, p, q);
    }
    if (p.val <= root.val && q.val >= root.val) {
        return root;
    }
};