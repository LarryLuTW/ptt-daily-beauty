import requests
from datetime import datetime
from mail import send_mail
from pprint import pprint

r = requests.get('https://us-central1-daily-beauty-209105.cloudfunctions.net/getDailyBeauties')
beauties = r.json()
# beauties = [{'date': '7/08',
#              'href': 'https://www.ptt.cc/bbs/Beauty/M.1530990046.A.580.html',
#              'mark': '',
#              'nVote': 54,
#              'title': '[正妹] 俄羅斯小龍女'},
#             {'date': '7/08',
#              'href': 'https://www.ptt.cc/bbs/Beauty/M.1531041085.A.1AB.html',
#              'mark': '',
#              'nVote': 24,
#              'title': '[正妹] 17歲'},
#             {'date': '7/08',
#              'href': 'https://www.ptt.cc/bbs/Beauty/M.1531040102.A.68C.html',
#              'mark': '',
#              'nVote': 18,
#              'title': '[正妹] 辰巳唯 (辰巳ゆい)'}]

def genContent(beauty):
    return '<h3><a href="{0}">{2}:{1}</a><h3>\n'.format(beauty['href'], beauty['title'], beauty['nVote'])

header = '<h1> 這是今天的日報 </h1>'
content = ''.join(map(genContent, beauties))
print(content)
# <h3><a href="https://www.ptt.cc/bbs/Beauty/M.1530990046.A.580.html">54:[正妹] 俄羅斯小龍女</a><h3>
# <h3><a href="https://www.ptt.cc/bbs/Beauty/M.1531041085.A.1AB.html">24:[正妹] 17歲</a><h3>
# <h3><a href="https://www.ptt.cc/bbs/Beauty/M.1531040102.A.68C.html">18:[正妹] 辰巳唯 (辰巳ゆい)</a><h3>

send_mail('pudding850806@gmail.com', '[日報第X期]這是 M/DD 的日報', header + content)