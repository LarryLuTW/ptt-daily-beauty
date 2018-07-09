from datetime import datetime
from mail import send_mail

content = """   <html>
                    <h1> Larry Lu 你好 </h1>
                    今天晚上想跟你約吃個飯，不知道你今天晚上有空嗎？
                </html>"""

send_mail('pudding850806@gmail.com', '你好，我是 subject', content)