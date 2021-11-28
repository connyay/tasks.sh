load("twilio", "client")

def run():
    accountSID, authToken, numberFrom = env("TWILIO_ACCOUNT_SID"), env("TWILIO_AUTH_TOKEN"), env("TWILIO_NUMBER_FROM")
    number = parameters.get("number", "")
    body = parameters.get("body", "hello world - tasks.sh")
    id = client(accountSID, authToken, numberFrom).SendSMS(number, body)
    printf("Message sent - ID: %s", id)

run()