from timeit import timeit


def benchmark(number=100000):
    def decorator(func):
        def wrappepr(*args, **kwargs):
            f = lambda: func(*args, **kwargs)
            t = timeit(f, number=number)
            print(t)
            return f()
        return wrappepr
    return decorator


@benchmark()
def bsort(A: list) ->  list:
    
    length = len(A)
    
    while True:
        flag=True
        for i in range(length-1):
            if A[i] > A[i+1]:
                A[i], A[i+1] = A[i+1], A[i]    
                #swap(A[i], A[i+1])
                flag=False
        if flag:
            break
    
    return A

if __name__ == "__main__":
    lst = [1, 6, 8, 9, 9, 3, -5]
    print(bsort(lst))
    
