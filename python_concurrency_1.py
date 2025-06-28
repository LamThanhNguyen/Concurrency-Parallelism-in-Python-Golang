import threading
import time

def fetch_data(source):
    print(f"Fetching data from source {source}")
    time.sleep(10)
    print(f"Source {source} done")

if __name__ == "__main__":
    threads = []
    for i in range(100):
        t = threading.Thread(target=fetch_data, args=(i,))
        threads.append(t)

    # All threads are created but NOT started yet.

    # Now, start all threads
    for t in threads:
        t.start()

    # time.sleep(0.0000000001)

    # for t in threads:
    #     t.join()

    print("Done.")