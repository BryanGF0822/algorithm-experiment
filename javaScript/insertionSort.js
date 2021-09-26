let arr = [45,2,3,14,11,0,1,45,23,99,46,12,11,7,67,66,36,49,23,21,111,75,80,69,98];

function insertionSort( myArr ) {
    var size = myArr.length,
        slot,
        tmp;
   
    for ( var item = 0; item < size; item++ ) { // outer loop     
      tmp = myArr[item];
      for ( slot = item - 1; slot >= 0 && myArr[slot] > tmp; slot-- ){ // inner loop
        myArr[ slot + 1 ] = myArr[slot];
      }
      myArr[ slot + 1 ] = tmp;
    }
    return myArr;
  };

insertionSort(arr);
console.log(arr);