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

## POST request body (expects JSON)
![Screen Shot 2022-01-26 at 6 55 35 AM](https://user-images.githubusercontent.com/24352472/151189518-55560440-6293-484b-b480-bb85b1af9e4d.png)

## Command to run

- $ go run main.go

## Heroku URI

## API endpoint

- POST "/email"
