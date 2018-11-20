from smtplib import SMTP

from email.mime.multipart import MIMEMultipart
from email.mime.text import MIMEText

msg = MIMEMultipart('alternative')
msg['Subject'] = "Link"
msg['From'] = "Daily Beauty <service@daily-beauty.xyz>"
msg['To'] = "pudding850806@gmail.com"

html = """\
<html>
  <head></head>
  <body>
    <p>Hi!<br>
       你好阿 <br>
       Here is the <a href="http://www.python.org">link</a> you wanted.
    </p>
  </body>
</html>
"""


msg.attach(MIMEText(html, 'html'))

# msg = """
# Hello, this is doge.
# """

smtp = SMTP('email-smtp.us-west-2.amazonaws.com', 587)
smtp.starttls()
smtp.login('AKIAIKPNOS3WJXVHATCQ', 'An1t90naxXpgSaoZeQBiHEqlLDzPv1C1ZbVy3XEIlyhs')

smtp.sendmail("service@daily-beauty.xyz", "pudding850806@gmail.com", msg.as_string())

smtp.close()