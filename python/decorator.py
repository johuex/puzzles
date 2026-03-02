import time, functools


def timer(func):
    """
    Simple decorator w/o params
    """
    @functools.wraps(func) # use it for correct __name__, __module__ and other items (metadata)
    def wrapper(*args, **kwargs):
        start = time.perf_counter_ns()
        res = func(*args, **kwargs)
        end = time.perf_counter_ns() - start
        print(f"perf time [{func.__name__}]: {end} ns")
        return res

    return wrapper


def retry(max_attempts: int = 3, delay: float|int = 1):
    """
    Decorator with args, retrier
    """
    def decorator(func):
        @functools.wraps(func)
        def wrapper(*args, **kwargs):
            for attempt in range(1, max_attempts + 1):
                try:
                    print(f"[RETRY] Attempt {attempt}/{max_attempts} for {func.__name__}")
                    return func(*args, **kwargs)
                except Exception as e:
                    if attempt == max_attempts:
                        raise
                    print(f"  Failed: {e}, retrying in {delay}s...")
                    time.sleep(delay)
            return None
        return wrapper
    return decorator


def memoize(maxsize=128):
    """
    Cache decorator
    """
    def decorator(func):
        @functools.wraps(func)
        @functools.lru_cache(maxsize=maxsize) # cache func call with params
        def wrapper(*args):
            print(f"[CACHE] Miss for {func.__name__}{args}")
            return func(*args)
        return wrapper
    return decorator

class MyDecorator:
    """
    Counter on each function call
    """
    def __init__(self, function):
        self.function = function
        self.counter = 0

    def __call__(self, *args, **kwargs):
        self.function(*args, **kwargs)
        self.counter+=1
        print(f"Called {self.function.__name__} {self.counter} times")


def add_calc(target):
    """
    Add function to class
    """
    def calc(self):
        return 42

    target.calc = calc
    return target


def singleton(cls):
    instance = None
    def wrapper(*args, **kwargs):
        nonlocal instance
        if instance is None:
            instance = cls(*args, **kwargs)
        return instance
    return wrapper

# --------------------------



@singleton
class Database:
    pass


@MyDecorator
def some_function():
    return 42


@timer
def calc():
    print(100*200)


@retry(5,0.5)
def exp_exception():
    raise RuntimeError


@timer # in this case decorator wrap only __init__ func
class Test:
    def huge_calculation(self):
        time.sleep(1)
        print(42)


@add_calc
class Empty:
    def __init__(self):
        print("Empty __init__")


@memoize(maxsize=10)
def fibonacci(n):
    """Calculate nth Fibonacci number."""
    if n < 2:
        return n
    return fibonacci(n - 1) + fibonacci(n - 2)



if __name__ == "__main__":
    calc()
    print()
    #exp_exception()
    print()
    obj = Test()
    obj.huge_calculation()
    print()
    some_function()
    some_function()
    some_function()
    print()
    obj = Empty()
    print(f"{obj.calc.__name__} result is {obj.calc()}")
    print()
    print(fibonacci(100))
    print(fibonacci(100))
