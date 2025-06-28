import asyncio

async def fetch_data(source):
    print(f"Fetching data from source {source}")
    await asyncio.sleep(1)
    print(f"Source {source} done")

async def main():
    tasks = [fetch_data(i) for i in range(100)]
    await asyncio.gather(*tasks)

asyncio.run(main())
print("All fetches complete")