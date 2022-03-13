/**
 * @param {number[]} data
 * @return {boolean}
 */
 var validUtf8 = function(data) {
    let i=0;
    while (i < data.length) {
        let num = data[i];
        if ((num & 1<<7) === 0) {
            i++
            continue
        }

        let nBit = 0;
        let mask = 1<<(7-nBit);
        while (nBit < 5 && (num & mask) === mask) {
            nBit++;
            mask = 1<<(7-nBit);
        }
        if (nBit <= 1 ||  nBit >= 5) {
            return false;
        }
        for (let j=1;j<nBit;j++) {
            let index = i+j;
            if (index >= data.length) {
                return false;
            }
            let v = data[index];
            if (((v >> 6) & 0b11) !== 0b10) {
                return false;
            }
        }
        i+=nBit;
    }
    return true;
};

console.log(validUtf8([145]));