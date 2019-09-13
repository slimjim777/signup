# Sign-up Form

## Run
```
export DATABASE_URL="postgres://..."
export PORT=8000
export DAY=3
export TITLE="Wednesday Night Sign-up"
export BANNER="iVBORw0K..."
export LIMIT=20

go run main.go
```

Banner: base64 encoded image (height=120px)
Limit: maximum number of attendees
