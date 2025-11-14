
def shaker_sort(A: list) -> list:
    
    left=0
    right=len(A)-1
    
    while left<=right:
        for i in range(left, right):
            if A[i] > A[i+1]:
                A[i], A[i+1] = A[i+1], A[i]
        left+=1

        for i in range(right, left, -1):
            if A[i-1] > A[i]:
                A[i-1], A[i] = A[i], A[i-1]
        right-=1
        
    return A

if __name__ == "__main__":
    lst = [1, 6, 8, 9, 9, 3, -5]
    print(shaker_sort(lst))