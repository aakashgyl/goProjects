# UrlShortener

HTTP Request:
```
curl -X POST http://localhost:8888/shortenurl -d \
'{
    "url":"mail.google.com/gmail/u/0"
}'
```

HTTP Response:
```
{
    "short_url":"localhost:8888/aaaaaaa"
}
```