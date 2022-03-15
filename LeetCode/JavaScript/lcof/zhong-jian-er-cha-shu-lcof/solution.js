/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {number[]} preorder
 * @param {number[]} inorder
 * @return {TreeNode}
 */
var buildTree = function (preorder, inorder) {
    if (preorder.length === 0) {
        return null;
    }
    const val = preorder[0];
    const root = new TreeNode(preorder[0]);
    const valIndex = inorder.findIndex(v => v === val);
    root.left = buildTree(preorder.slice(1, valIndex + 1), inorder.slice(0, valIndex));
    root.right = buildTree(preorder.slice(valIndex + 1), inorder.slice(valIndex + 1));
    return root;
};