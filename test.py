import requests

urlGet = "http://127.0.0.1:5000/send-json"
urlPost = "http://127.0.0.1:5000/users"

body = {
    "name" : "Raju Rastogi",
    "email" : "Raju Rastogi@gmail.com"
}

# response: requests.Response = requests.post(url=urlGet, json=body)


response: requests.Response = requests.post(url=urlPost, json=body)


print(response.status_code)
print(response.text)

# # Only try to parse JSON if there is content and it's JSON
# if response.text and 'application/json' in response.headers.get('Content-Type', ''):
#     try:
#         data = response.json()
#         print("JSON:", data)
#     except requests.exceptions.JSONDecodeError as e:
#         print("Failed to decode JSON:", e)
# else:
#     print("Response is not JSON")