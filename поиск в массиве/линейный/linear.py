

def linear_search(arr, target) -> int:
    
    for i, el in enumerate(arr):
        if el == target:
            return i
    return -1

