


def merge(left: list, right: list) -> list:
    result = []
    i = j = 0
    
    while i < len(left) and j < len(right):
        if left[i] <= right[j]:
            result.append(left[i])
            i += 1
        else:
            result.append(right[j])
            j += 1
    result.extend(left[i:])
    result.extend(right[j:])
    return result


def msort(A: list) -> list:
    l = len(A)
    if l <= 1:
        return A
    
    m = l // 2
    left = msort(A[:m])
    right = msort(A[m:])
    
    return merge(left, right)


if __name__ == "__main__":
    lst = [1, 6, 8, 9, 9, 3, -5]
    print(msort(lst))