import logging
from flask import Flask

from mail import send_mail
from template import generateHTML, generateSubject
from api import getBeauties

app = Flask(__name__)


@app.route('/test')
def test():
    beauties = getBeauties()
    html = generateHTML(beauties)
    subject = generateSubject()
    toMails = ['pudding850806@gmail.com']
    send_mail(toMails, subject, html)
    return 'test success'

@app.route('/publish')
def publish():
    beauties = getBeauties()
    html = generateHTML(beauties)
    subject = generateSubject()
    toMails = ['pudding850806@gmail.com', 'w5151381guy@gmail.com']
    send_mail(toMails, subject, html)
    return 'publish success'


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
