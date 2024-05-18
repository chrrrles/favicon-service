import redis
import requests
import favicon
import os


def extract_favicon(fqdn):
    try:
        favicons = favicon.get(f"https://{fqdn}")
    except Exception as e:
        print(f"Error: unable to fetch favicon from {fqdn}: {e}")
        return

    if not favicons:
        print(f"Error: no favicons found for {fqdn}")
        return

    # Find the smallest favicon image
    favicon_url = favicons[-1].url

    try:
        response = requests.get(favicon_url)
    except Exception as e:
        print(f"Error: unable to retrieve favicon from {fqdn}: {e}")
        return

    host, port = os.environ["FAVICON_REDIS_SERVER"].split(":")
    r = redis.Redis(host=host, port=port, db=0)
    r.set(fqdn, response.content)

    print(f"Successfully extracted and stored favicon for {fqdn}")
