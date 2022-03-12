/**
 * // Definition for a Node.
 * function Node(val,children) {
 *    this.val = val;
 *    this.children = children;
 * };
 */

/**
 * @param {Node|null} root
 * @return {number[]}
 */
 var postorder = function(root) {
    var result = [];
    if (!root) {
        return result;
    }
    root.children.forEach((c) => {
        result = result.concat(postorder(c));
    })
    result.push(root.val);
    return result;
};