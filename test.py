def mergeArrays(arr1, arr2):
    merged_map = {}
    for i in arr1:
        merged_map[i] = True
    for i in arr2:
        merged_map[i] = True
    merged_arr = list(merged_map.keys())

    sameString = {}
    for i in arr1:
        sameString[i] = True
    for i in arr2:
        if i in sameString:
            del sameString[i]
        else:
            sameString[i] = True
    arr = list(sameString.keys())

    return merged_arr, arr

arr1 = ["a", "b", "c"]
arr2 = ["b", "c", "d"]
merged_arr, common_arr = mergeArrays(arr1, arr2)
print(merged_arr)
print(common_arr) 
