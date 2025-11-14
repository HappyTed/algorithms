from timeit import timeit

def benchmark(number=100000):
    def decorator(func):
        def wrapper(*args, **kwargs):
            f = lambda: func(*args, **kwargs)
            t = timeit(f, number=number)
            print(f"{func.__name__}: {t:.6f} сек")
            return func(*args, **kwargs)
        return wrapper
    return decorator