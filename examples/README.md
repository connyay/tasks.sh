# Examples

## Hello World

[hello.star](./hello.star)

```console
$ go run ./cmd/eval --star examples/hello.star
2021/11/28 16:27:02 Hello World!
```

## Yahoo Finance

[yfinance.star](./yfinance.star)

```console
$ go run ./cmd/eval --star examples/yfinance.star
2021/11/28 16:45:04 Loading AAPL ticker data
[AAPL]: 2021-11-26 09:30:00 EST (open=159.56500244140625, close=159.92010498046875)
[AAPL]: 2021-11-26 09:31:00 EST (open=159.9980926513672, close=160.11000061035156)
[AAPL]: 2021-11-26 09:32:00 EST (open=160.08999633789062, close=159.9199981689453)
[AAPL]: 2021-11-26 09:33:00 EST (open=160, close=160.08999633789062)
[AAPL]: 2021-11-26 09:34:00 EST (open=160.08990478515625, close=160.25999450683594)
```

```console
$ go run ./cmd/eval --star examples/yfinance.star -p 'ticker=GME'
2021/11/28 16:45:47 Loading GME ticker data
[GME]: 2021-11-26 09:30:00 EST (open=208.0800018310547, close=209.5)
[GME]: 2021-11-26 09:31:00 EST (open=208.74000549316406, close=206)
[GME]: 2021-11-26 09:32:00 EST (open=206.6199951171875, close=205.32000732421875)
[GME]: 2021-11-26 09:33:00 EST (open=205.13999938964844, close=205.2052001953125)
[GME]: 2021-11-26 09:34:00 EST (open=204.13499450683594, close=204)
```

## Twitter

[twitter.star](./twitter.star)

This example requires a `TWITTER_BEARER_TOKEN` environment variable. A bearer
token can be retrieved from the [twitter developer
dashboard](https://developer.twitter.com/en/portal/dashboard).

```console
$ TWITTER_BEARER_TOKEN=example-token go run ./cmd/eval --star examples/twitter.star
[Funfacts]: "RT @dinozoiks: 👀🚨👉 Sneaky Sunday Dinosaw Alert! The new video is uploaded and ready to watch. Check it out for 7 curious things you may hav…" at 2021-11-28 19:54:57 +0000 UTC
[Funfacts]: "In terms of time, a tyrannosaurus is closer to us than to a stegosaurus." at 2021-11-27 14:53:16 +0000 UTC
[Funfacts]: "Google was founded in 2005 and attracted investment of $8.5m. It was bought a year later by Google for $1.65 billion.\n\nIn 2021, YouTube makes $1.65 billion every 3 weeks." at 2021-11-22 09:30:44 +0000 UTC
[Funfacts]: "A horse actually produces about 15 horsepower. An average human has just over 1 horsepower and can peak at 5 horsepower. 🐴💪" at 2021-11-06 23:36:18 +0000 UTC
[Funfacts]: "RT @dinozoiks: 🔥🔥🔥 The new #Dinosaw is out for week 44. In this week's edition, Meta, building with moving walls, cars that can charge each…" at 2021-11-01 14:31:15 +0000 UTC
```

## Twilio

[twilio.star](./twilio.star)

This example requires `TWILIO_ACCOUNT_SID`, `TWILIO_AUTH_TOKEN`, and `TWILIO_NUMBER_FROM` environment variables. These values can be retrieved from the [twilio developer
console](https://console.twilio.com/).

```console
$ TWILIO_ACCOUNT_SID=$accountSID TWILIO_AUTH_TOKEN=$authToken TWILIO_NUMBER_FROM=$numberFrom \
go run ./cmd/eval --star examples/twilio.star -p "number=$NUMBER_TO" -p 'body=hello there'
Message sent. ID=SM443f45a663c5472294aed98117868202
```

## Reddit

[reddit.star](./reddit.star)

This example requires `REDDIT_BOT_CLIENT_ID`, `REDDIT_BOT_CLIENT_SECRET`, `REDDIT_BOT_USERNAME`, and `REDDIT_BOT_PASSWORD` environment variables. These values can be retrieved from the [reddit apps
page](https://old.reddit.com/prefs/apps/).

```console
$ go run ./cmd/eval --star examples/reddit.star
[earthporn]: "Summer morning in Tuscany, Italy. [OC] [4031x3024]" (/u/Leo1762 at Sun, 28 Nov 2021 16:25:01 MST)
[earthporn]: "Olympic National Park -OC- 4032x3024" (/u/taoofjerry at Sun, 28 Nov 2021 16:14:41 MST)
[earthporn]: "Mount Rainier National Park - OC- 4032x3024" (/u/taoofjerry at Sun, 28 Nov 2021 16:13:22 MST)
[earthporn]: "Fallen tree in River. Atlanta, GA [3814x3054] [OC]" (/u/fenrirctj89 at Sun, 28 Nov 2021 16:22:01 MST)
[earthporn]: "Fallen tree on a trail in Atlanta, GA. [3677x2942] [OC]" (/u/fenrirctj89 at Sun, 28 Nov 2021 16:20:18 MST)
```

```console
$ go run ./cmd/eval --star examples/reddit.star -p 'subreddit=showerthoughts' -p 'count=3'
[showerthoughts]: "Our organs are our workers. Their salaries are nutrients that we give in order of doing their jobs. People who eat healthy food pay higher wages to them" (/u/WhiHd at Sun, 28 Nov 2021 16:14:18 MST)
[showerthoughts]: "Condoms nowadays are made by automated assembly lines meaning robots are helping to prevent human reproduction." (/u/o1pa1 at Sun, 28 Nov 2021 16:06:38 MST)
[showerthoughts]: "People laugh and giggle when animals have sex in public, but not when humans do it." (/u/Intense_Anal_Lawyer at Sun, 28 Nov 2021 16:45:17 MST)
```

## Reddit + Twilio

[text-new-reddit-posts.star](./text-new-reddit-posts.star)

This example is a combination of the above reddit and twilio examples. It pulls reddit posts and sends a text message with the titles.
