/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @return {number}
 */
 var maxDepth = function(root) {
    let max = 0;
    let depth = 0;

    let dfs = function(tn) {
        if (!tn) {
            return;
        }
        depth++; 
        if (depth > max) {
            max = depth;
        }
        dfs(tn.left);
        dfs(tn.right);
        depth--;
    };

    dfs(root);
    return max;
};