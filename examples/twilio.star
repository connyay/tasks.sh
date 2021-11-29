load("twilio", "client")
load("environment", "TWILIO_ACCOUNT_SID", "TWILIO_AUTH_TOKEN", "TWILIO_NUMBER_FROM")

def run():
    number = parameters.get("number", "")
    body = parameters.get("body", "hello world - tasks.sh")
    id = client(TWILIO_ACCOUNT_SID, TWILIO_AUTH_TOKEN, TWILIO_NUMBER_FROM).SendSMS(number, body)
    printf("Message sent. ID=%s\n", id)

run()