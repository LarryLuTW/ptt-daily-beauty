import logging
import requests

from flask import Flask
from mail import send_mail
from datetime import datetime

app = Flask(__name__)

def genContent(beauty):
    return '<h3><a href="{0}">{2}:{1}</a><h3>\n'.format(beauty['href'], beauty['title'], beauty['nVote'])

@app.route('/')
def hello():
    """Return a friendly HTTP greeting."""
    # send_mail('pudding850806@gmail.com', '標題', '<html>你好阿 Hello World</html>' + str(datetime.now()))

    r = requests.get('https://us-central1-daily-beauty-209105.cloudfunctions.net/getDailyBeauties')
    beauties = r.json()

    header = '<h1> 這是今天的日報 </h1>'
    content = ''.join(map(genContent, beauties))
    print(content)
    send_mail('pudding850806@gmail.com', '[日報第999期]這是 M/DD 的日報', header + content)

    return '發送成功'


@app.errorhandler(500)
def server_error(e):
    logging.exception('An error occurred during a request.')
    return """
    An internal error occurred: <pre>{}</pre>
    See logs for full stacktrace.
    """.format(e), 500


if __name__ == '__main__':
    # This is used when running locally. Gunicorn is used to run the
    # application on Google App Engine. See entrypoint in app.yaml.
    app.run(host='0.0.0.0', port=8080, debug=True)
