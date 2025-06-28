import threading
import time

def fetch_data(source, start_event):
    start_event.wait()  # Wait until main thread signals "go"
    print(f"Fetching data from source {source}")
    time.sleep(10)
    print(f"Source {source} done")

if __name__ == "__main__":
    start_event = threading.Event()
    threads = []

    for i in range(100):
        t = threading.Thread(target=fetch_data, args=(i, start_event))
        threads.append(t)
        t.start()

    # Small sleep to ensure all threads have started and are waiting on the event
    time.sleep(0.1)
    start_event.set()  # Signal all threads to start their task

    # for t in threads:
    #     t.join()

    print("All fetches complete")