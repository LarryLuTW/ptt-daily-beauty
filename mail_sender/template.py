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

def genHeader():
    header = '<h1> 這是今天的日報 </h1>'
    return header

def genItem(beauty):
    return  ''' 
    <h3>
        <a href="{0}"> {2}:{1} </a>
    <h3>
    '''.format(beauty['href'], beauty['title'], beauty['nVote'])

def generateHTML(beauties):
    header = genHeader()
    content = ''.join(map(genItem, beauties))
    html = header + content
    print(html)
    return html

def generateSubject():
    subject = '[日報第001期]這是 M/DD 的日報'
    return subject