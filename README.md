# Email Microservice

- Microservice for CS361 - Software Engineering I
- Email Microservice that handles sending emails for you

## Authentication

- Must provide microservice with sender email & pw and recipient email

### Give microservice access to gmail account

- You will need to authorize less secure apps on your Google account
- https://kb.synology.com/en-global/SRM/tutorial/How_to_use_Gmail_SMTP_server_to_send_emails_for_SRM

## Currently only setup for Gmail

- Should be similar to other email providers

## Troubleshooting

- If the POST route works, but emails are not being sent, you may need to change your email password if it hasn't been changed in a while

## Command to run

- $ go run main.go

## Heroku Deployment

- https://cryptic-basin-36672.herokuapp.com

## POST /email

senderEmail: <string>
password: <string>
senderName: <string>
subject: <string>
htmlTemplate: <string> (Currently under construction)
recipients: []
{name: <string>, email: <string>}
]
