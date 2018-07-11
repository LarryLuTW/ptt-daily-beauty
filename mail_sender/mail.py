from __future__ import print_function
from apiclient.discovery import build
from httplib2 import Http
from oauth2client import file, client, tools
from email.mime.text import MIMEText
import base64
from googleapiclient.errors import HttpError

def create_message(sender, to, subject, message_text):
  message = MIMEText(message_text, 'html')
  message['to'] = to
  message['from'] = sender
  message['subject'] = subject
  return {'raw': base64.urlsafe_b64encode(message.as_bytes()).decode('utf-8')}

def send_message(service, user_id, message):
  """Send an email message.

  Args:
    service: Authorized Gmail API service instance.
    user_id: User's email address. The special value "me"
    can be used to indicate the authenticated user.
    message: Message to be sent.

  Returns:
    Sent Message.
  """
  try:
    message = (service.users().messages().send(userId=user_id, body=message)
               .execute())
    print('Message Id: %s' % message['id'])
    return message
  except HttpError as error:
    print('An error occurred: %s' % error)

def send_mail(toMails, subject, content):
      # Setup the Gmail API
  # scope ref: https://developers.google.com/gmail/api/auth/scopes
  SCOPES = 'https://www.googleapis.com/auth/gmail.compose'
  store = file.Storage('credentials.json')
  creds = store.get()
  if not creds or creds.invalid:
      flow = client.flow_from_clientsecrets('client_secret.json', SCOPES)
      creds = tools.run_flow(flow, store)
  service = build('gmail', 'v1', http=creds.authorize(Http()))

  for to in toMails:
    message = create_message("me", to, subject, content)
    send_message(service, "me", message)
    print('send to {} success'.format(to))