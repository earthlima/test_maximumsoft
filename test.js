function mergeArrays(arr1, arr2) {
    let mergedMap = {};
    for (let i of arr1) {
      mergedMap[i] = true;
    }
    for (let i of arr2) {
      mergedMap[i] = true;
    }
    let mergedArr = Object.keys(mergedMap);
    
    let sameStr = {};
    for (let i of arr1) {
      sameStr[i] = true;
    }
    for (let i of arr2) {
      if (sameStr[i]) {
        delete sameStr[i];
      } else {
        sameStr[i] = true;
      }
    }
    let commonArr = Object.keys(sameStr);
    
    return [mergedArr, commonArr];
  }
  
  let arr1 = ["a", "b", "c"];
  let arr2 = ["b", "c", "d"];
  let [mergedArr, commonArr] = mergeArrays(arr1, arr2);
  console.log(mergedArr);
  console.log(commonArr); 
  