import math
import time
from multiprocessing import Pool, cpu_count

def compute_chunk(start_end):
    start, end = start_end
    return sum(math.sqrt(i) for i in range(start, end))

if __name__ == "__main__":
    total = 100_000_000
    cores = cpu_count()
    chunk = total // cores
    ranges = [(i * chunk, (i + 1) * chunk) for i in range(cores)]

    start_time = time.time()

    with Pool(cores) as pool:
        results = pool.map(compute_chunk, ranges)

    total_sum = sum(results)
    print(f"Result: {total_sum:.2f}, Time: {time.time() - start_time:.2f} seconds")
