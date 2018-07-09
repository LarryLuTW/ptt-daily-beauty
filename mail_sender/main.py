import logging

from flask import Flask
from mail import send_mail
from datetime import datetime

app = Flask(__name__)

@app.route('/')
def hello():
    """Return a friendly HTTP greeting."""
    send_mail('pudding850806@gmail.com', '標題', '<html>你好阿 Hello World</html>' + str(datetime.now()))
    return 'Hello World! I am Larry.'


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
