/**
 * @param {string[]} words
 * @return {string}
 */
var longestWord = function (words) {
    const wordMap = words.reduce((p, c) => {
        p[c] = true;
        return p;
    }, {});
    let longest = "";
    for (let i = 0; i < words.length; i++) {
        const word = words[i];
        let isCandidate = true;
        for (let j = 0; j < word.length; j++) {
            if (!wordMap[word.substring(0, j + 1)]) {
                isCandidate = false;
                break;
            }
        }
        if (isCandidate) {
            if (word.length > longest.length) {
                longest = word;
            } else if (word.length === longest.length && word.localeCompare(longest) === -1) {
                longest = word;
            }

        }
    }
    return longest;
};

console.log(longestWord(["a", "banana", "app", "appl", "ap", "apply", "apple"]));