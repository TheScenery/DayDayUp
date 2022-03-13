function Node(val,left,right) {
    this.val = val;
    this.left = left;
    this.right = right;
 };
/**
 * @param {Node} root
 * @return {Node}
 */
 var treeToDoublyList = function(root) {
    if (!root) {
        return root
    }
    var pre, head
    
    var dfs = function(n) {
        if (!n) {
            return
        }
        dfs(n.left)
        if (!pre) {
            head = n
        } else {
            n.left = pre
            pre.right = n
        }
        pre = n
        dfs(n.right)
    }

    dfs(root)
    head.left = pre
    pre.right = head
    return head
};