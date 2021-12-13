# Summary

Yesterdy which is Decemeber 12, 2021. One of my friend send me a message on twitter that he want me to write a script that brute force list of headers with payloads.
Today, I've written it.

# Install

```
▶ go get github.com/channyein1337/log4j-detect
```

# Usage

```
▶ log4j-detect -h
Usage of log4j-detect:
  -hf string
        Set the headers file
  -p string
        Set the payload file
 ```
 
 Scan list of urls with headers and payloads
 ```
 ▶ echo https://www.google.com | log4j-detect -hf header.txt -p payload.txt
 ```
 
 ![](https://github.com/channyein1337/log4j-detect/blob/main/log4j-detect.png?raw=true)
 
# Note
If there is any issue in running this tool, just contact me on twiiter or create issue.

Thanks!
