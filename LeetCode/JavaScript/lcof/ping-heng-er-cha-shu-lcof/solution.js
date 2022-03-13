/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
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


/**
 * @param {TreeNode} root
 * @return {boolean}
 */
 var isBalanced = function(root) {
    if (!root) {
        return true;
    }
    if (Math.abs(maxDepth(root.left) - maxDepth(root.right)) > 1) {
        return false;
    }
    return isBalanced(root.left) && isBalanced(root.right);
};