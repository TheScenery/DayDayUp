//评测题目: // Sample data.
const html = '<div><div>Orange</div>Directory<ul><div>Apple</div><div>Orange</div></ul></div>';
const obj = {
    "tag": "div",
    "children": [
        "Directory",
        {
            "tag": "ul",
            "children": [{
                    "tag": "li",
                    "children": [
                        "Apple"
                    ]
                },
                {
                    "tag": "li",
                    "children": [
                        "Orange"
                    ]
                }
            ]
        }
    ]
};

function stringfySubObject(obj) {
    if (!obj.children) {
        return obj;
    }
    return `<${obj.tag}>${obj.children.map(child => stringfySubObject(child)).join('')}</${obj.tag}>`;
}


// Convert object to html.
function stringify(obj) {
    // Your Code Here
    return stringfySubObject(obj);
}

// Convert html to object.
function getDomElement(charArray) {
    return charArray.join('');
}


function parse(html) {
    // Your Code Here	
    const charStack = [];
    const domStack = [];
    for (let i = 0; i < html.length; i++) {
        const char = html[i];
        if (char === '<') {
            const domElement = getDomElement(charStack);
            charStack.length = 0;
            if (domElement !== "") {
                domStack.push(domElement);
            }
            charStack.push(char);
        } else if (char === '>') {
            charStack.push(char);
            const domElement = getDomElement(charStack);
            if (domElement.indexOf('</') === 0) {
                const children = [];
                const headTag = domElement.replace('/', '');
                let child = domStack.pop();
                while (child !== headTag) {
                    children.unshift(child);
                    child = domStack.pop();
                }
                const node = {
                    tag: headTag.substring(1, headTag.length - 1),
                    children,
                }
                domStack.push(node);
            } else {
                domStack.push(domElement);
            }
            charStack.length = 0;
        } else {
            charStack.push(char)
        }
    }
    return domStack[0];
}


console.log(stringify(obj));

console.log(parse(html));


function stringify_1(obj) {
    if (typeof obj === 'string') {
        return obj;
    }
    const tag = obj.tag;
    const children = obj.children.map(c => stringify_1(c)).join('');
    return `<${tag}>${children}</${tag}>`
}

function parse_1(html) {
    let code = 'return ['
    + html.replace(/>([^<]+?)</g, '>"$1",<')
    .replace(/<(\w+)>/g, '{tag: "$1", children: [')
    .replace(/<\/\w+>/g, ']},')
    + ']'
    return (new Function(code))().pop()
}

console.log(stringify_1(obj));

console.log(parse_1(html));