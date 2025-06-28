import multiprocessing
import time

def heavy_computation(n):
    result = 0
    for i in range(n):
        # Simulate CPU work (compute Fibonacci(20))
        a, b = 0, 1
        for _ in range(20):
            a, b = b, a + b
        result += a
    return result

if __name__ == "__main__":
    num_workers = 4
    tasks_per_worker = 100000

    start = time.time()

    with multiprocessing.Pool(num_workers) as pool:
        results = pool.map(heavy_computation, [tasks_per_worker] * num_workers)

    total = sum(results)
    elapsed = time.time() - start
    print(f"Total result: {total}")
    print(f"Elapsed: {elapsed:.2f}s (on {num_workers} processes)")
