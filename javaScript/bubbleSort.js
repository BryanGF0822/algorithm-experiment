let arr = [45,2,3,14,11,0,1,45,23,99,46,12,11,7,67,66,36,49,23,21,111,75,80,69,98];

function bubbleSort(items) {
    var length = items.length;  
    for (var i = 0; i < length; i++) { 
          for (var j = 0; j < (length - i - 1); j++) { 
                   if(items[j] > items[j+1]) {
                           var tmp = items[j]; 
                items[j] = items[j+1]; 
                items[j+1] = tmp; 
            }
        }        
    }
}

bubbleSort(arr);
console.log(arr);