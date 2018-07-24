import requests

def getBeauties():
    url = 'https://us-central1-daily-beauty-209105.cloudfunctions.net/getDailyBeauties'
    r = requests.get(url)
    beauties = r.json()
    # Sort by popularity ascendantly
    return list(reversed(beauties))