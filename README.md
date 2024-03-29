# What's this
You can add a white margin to your photo.

# How to use
easy peasy.
```
insta-margin $ go get golang.org/x/image/draw
insta-margin $ go build insta.go 
insta-margin $ ./insta landscape.jpg 
```

Output file has random base62 suffix, don't worry about overwrite.

## Options
### -p
You can specify a white margin percentage by `-p` option.  
Default value is defined as `defaultMarginPercent`, and it is `80`.

### -d
Also, you can specify a directory with `-d` option.  
All photos inside the specified directory will be converted.

```
insta-margin $ ./insta -d photos/
```

# Configuration
You can configure `baseSize` and `defaultMarginPercent`.  
For now, as Instagram's maximum upload size is 1080 x 1080, the default `baseSize` value follows this limit.

# Sample
Landscape ratio.  

<img src="https://github.com/ykhr53/insta-margin/blob/images/land.png" width="650">
  
  

Portrait ratio.  

<img src="https://github.com/ykhr53/insta-margin/blob/images/port.png" width="650">
  

Thanks, all sample photos are mine 😊  
https://www.flickr.com/people/yukihira_fl/