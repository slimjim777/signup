# Sign-up Form

Simple sign-up form that tracks a list of those attending and not attending a weekly event.
No logins. No authentication. No authorisation. No barriers.

## Test it
```
export DATABASE_URL="postgres://..."
export PORT=8000
export DAY=3
export TITLE="Wednesday Night Sign-up"
export BANNER="iVBORw0K..."
export LIMIT=20
export LABELPLUS="Attending"
export LABELMINUS="Not attending"

go run main.go
```

- BANNER: base64 encoded image (height=120px)
- LIMIT: maximum number of attendees
- DAY: day of the week for the rota (Sunday=0, Monday=1,...)

## Run it

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

Note: a simple way to generate the base64-encoded banner image is to upload it to a service like [https://base64.guru/converter/encode/file](https://base64.guru/converter/encode/file). Then copy the text into the BANNER field.
