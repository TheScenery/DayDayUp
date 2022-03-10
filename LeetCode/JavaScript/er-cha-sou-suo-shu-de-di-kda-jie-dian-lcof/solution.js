/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @param {number} k
 * @return {number}
 */
 var kthLargest = function(root, k) {
    var a = []
    var dfs = function(tn) {
        if (!tn) {
            return
        }
        dfs(tn.left)
        a.push(tn)
        dfs(tn.right)
    }
    dfs(root)
    var index = a.length - k
    if (index < 0) {
        return
    }
    return a[index].val
};