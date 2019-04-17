/*

some interesting problems about N!

*/


const factorial = (n) => n === 1 ? 1 : n * factorial(n - 1);

function testFractorial() {
    console.log('MAX Safe Integer: ', Number.MAX_SAFE_INTEGER)
    for (let i = 1; ; i++) {
        const result = factorial(i);
        if (result > Number.MAX_SAFE_INTEGER) {
            console.log('Max N find: ', i);
            return;
        }
        console.log(`${i}! = ${result}`);
    }
}

testFractorial()


function tailZero(n) {
    let total = 0;
    for (let i = 1; i <=n; i++) {
        let t = i;
        while (t % 5 === 0) {
            total++;
            t /= 5;
        }
    }
    console.log('tail zeros:', total)
}

tailZero(10);