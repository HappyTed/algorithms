from math import *
from timeit import timeit

SRC = [3, 14, 15, 19, 26, 53, 58]
TARGET = 53

def benchmark(number=100000):
    def decorator(func):
        def wrapper(*args, **kwargs):
            f = lambda: func(*args, **kwargs)
            t = timeit(f, number=number)
            print(f"{func.__name__}: {t:.6f} сек")
            return func(*args, **kwargs)
        return wrapper
    return decorator
    
@benchmark()
def binary_array_search(A, target):
    lo=0
    hi=len(A)-1
    
    while lo <= hi:
        mid = (lo+hi) // 2

        if target < A[mid]:
            hi = mid - 1
        elif target > A[mid]:
            lo = mid + 1
        else:
            return True
    
    return False

@benchmark()
def binary_array_search_idx(A, target):
    lo=0
    hi=len(A)-1
    
    while lo <= hi:
        mid = (lo+hi) // 2

        if target < A[mid]:
            hi = mid - 1
        elif target > A[mid]:
            lo = mid + 1
        else:
            return mid
    
    return -1

@benchmark()
def binary_array_search_optimisation(A, target):
    """
    только 1 действие над элементом A
    
    обращение по индексу затратнее по ресурсам, чем один раз произвести сравнение обычных целых чисел
    """
    lo=0
    hi=len(A)-1
    
    while lo <= hi:
        mid = (lo+hi) // 2
        
        diff = target - A[mid]

        if target < diff:
            hi = mid - 1
        elif target > diff:
            lo = mid + 1
        else:
            return True
    
    return False

# верхняя граница для наихудшего случая для конкретной реализации алгоритма
def k(A: list):
    N = len(A)
    return floor(log2(N)) + 1 

# класс сложности алгоритма O(log N) - логарифический, потому что определяющая его функ­ция f(N) = log N.

if __name__ == "__main__":
    print(binary_array_search(SRC, TARGET))
    print(binary_array_search_idx(SRC, TARGET))
    print(binary_array_search_optimisation(SRC, TARGET))
    
    

