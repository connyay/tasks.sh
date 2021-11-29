load("twilio", "client")
load("environment", "TWILIO_ACCOUNT_SID", "TWILIO_AUTH_TOKEN", "TWILIO_NUMBER_FROM")

def main(args):
    number = args.get("number", "")
    body = args.get("body", "hello world - tasks.sh")
    id = client(TWILIO_ACCOUNT_SID, TWILIO_AUTH_TOKEN, TWILIO_NUMBER_FROM).SendSMS(number, body)
    printf("Message sent. ID=%s\n", id)
