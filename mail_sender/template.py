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

from datetime import datetime, timezone, timedelta

css = """
.content {
    padding: 0 8%;
}
.title {
    font-size: 38px;
    font-family: Georgia,Times,'Times New Roman',serif;
    font-weight: bold;
    color: #4e4e4e;
    text-align: center;
}
.list-title {
    font-family: Tahoma,Verdana,Segoe,sans-serif;
    font-size: 26px;
    font-weight: bold;
    color: #4e4e4e;
    border-bottom: 1px solid #ebebeb;
    padding-bottom: 10px;
    word-break: break-word;
}
.item {
    text-decoration: none;
    color: #2196f3;
}
"""

def genTitle():
    title = '''
    <h1 class="title">Daily Beauty 表特日報</h1>
    <div class="list-title"> 本日精選 </div>
    '''
    return title

def genItem(beauty):
    return  ''' 
    <h2>
        <a class="item" href="{0}"> {1} </a>
    </h2>
    '''.format(beauty['href'], beauty['title'])

def generateHTML(beauties):
    title = genTitle()
    content = ''.join(map(genItem, beauties))
    html = '''
    <html>
        <head>
            <style>{0}</style>
        </head>
        <body class="content">
            {1}
            {2}
        </body>
    </html>
    '''.format(css, title, content)
    print(html)
    return html

def generateSubject():
    tz = timezone(offset=timedelta(hours=8))
    date = datetime.today().astimezone(tz).strftime('%Y-%m-%d %H:%M:%S')
    subject = '[表特日報-{0}]'.format(date)
    return subject