/*
    Title: combination two dimension array
    Description:
    combination two dimension array
    Example:
    Input: 
    [[1,2], [3, 4], [5, 6]]
    Output:
    [ 1, 3, 5 ] 
    [ 1, 3, 6 ] 
    [ 1, 4, 5 ] 
    [ 1, 4, 6 ] 
    [ 2, 3, 5 ] 
    [ 2, 3, 6 ] 
    [ 2, 4, 5 ] 
    [ 2, 4, 6 ]

    Language: Javascript
*/

function joinArray(array) {
	const length = array.length;
  	if (length >= 2) {
    	const array1Length = array[0].length;
      	const array2Length = array[1].length;
      	const tempArray = [];
      	for (let i = 0; i < array1Length; i++) {
        	for (let j = 0; j< array2Length; j++) {
            	tempArray.push([array[0][i], array[1][j]].join(', '))
            }
        }
      	const newArray = [];
      	for (let i = 2; i<length;i++) {
        	newArray.push(array[i]);
        }
      	newArray.unshift(tempArray);
      	return joinArray(newArray);
    } else {
    	return array[0];
    }
}

function main(inputValues) {
	if (!inputValues || !inputValues.length) {
        return;
    }
  	const joidedArray = joinArray(inputValues);
  	if (joidedArray.length <= 0) {
    	return
    }
  	joidedArray.forEach(result => console.log('[', result, ']'))
}


main([[1,2], [3, 4], [5, 6]])