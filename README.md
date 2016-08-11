
# Cassandra video tutorial downloader


## How to use
Log into datastax, and obtain the cookies:

```
SESSdse=xxxxxxx%40xxxxx.com; 
SSESSxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
```

use this on the datastax page url to get page source, eg:
- https://academy.datastax.com/courses/ds220-data-modeling/introduction-introduction-killrvideo

Then grab the vimeo video id and use in the url below 
- https://player.vimeo.com/video/VIDEO_ID
- example: https://player.vimeo.com/video/133680477

From the returned data, get the url from the JSON of a 720p video and download from uel



## Config
Structure:

```
{
	"session": "SESSdse=some_email@domain.com",
	"authToken": "some_long_token_1324io4rhoerhg09fgh024n0h0g9h9gh"
}
```

For the `session` and `authToken`, log into datastax and copy the cookie info.