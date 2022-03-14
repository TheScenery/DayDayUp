function getAnces(g, n, result) {
    if (result[n]) {
        return
    }
    const nodes = g[n];
    let res = new Set(nodes);
    nodes.forEach(node => {
        getAnces(g, node, result);
        res = new Set([...res, ...result[node]])
    });
    return result[n] = [...res].sort((a, b) => a - b);
}

/**
* @param {number} n
* @param {number[][]} edges
* @return {number[][]}
*/
var getAncestors = function (n, edges) {
    const res = [];
    const g = [];
    for (let i = 0; i < n; i++) {
        g.push([]);
    }
    for (let i = 0; i < edges.length; i++) {
        const edge = edges[i];
        g[edge[1]].push(edge[0]);
    }
    for (let i = 0; i < g.length; i++) {
        getAnces(g, i, res)
    }
    return res
};

let edgeList = [[0, 3], [0, 4], [1, 3], [2, 4], [2, 7], [3, 5], [3, 6], [3, 7], [4, 6]]
console.log(getAncestors(8, edgeList));