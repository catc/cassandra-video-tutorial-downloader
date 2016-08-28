
# Cassandra video tutorial downloader


## How to use
Log into datastax, and obtain the cookies:

```
SESSdse=xxxxxxx%40xxxxx.com; 
SSESSxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
```

Create `config.json` with structure:
```
{
	"authToken": "SESSdse=xxxxxxx%40xxxxx.com=SSESSxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
}
```

Add any datastax tutorial urls to the `[]string` in `main.go` and run.