# What's this
When we upload a photo to Instagram, we have to think about aspect ratio...  
From now on, you don't need to worry about your photo's aspect ratio!  
You can just upload your photo, in the original ratio.  
And of course, margins are beautiful, good photo needs proper margins.

# How to use
easy peasy.
```
insta-margin $ go get golang.org/x/image/draw
insta-margin $ go build insta.go 
insta-margin $ ./insta landscape.jpg 
```

Output file has random base64 suffix, don't worry about overwrite.

# Configuration
You can configure `baseSize` and `marginPercent`.  
For now, as Instagram's maximum upload size is 1080 x 1080, the default `baseSize` value follows this limit.

# Sample
Landscape ratio.  

<img src="https://github.com/ykhr53/insta-margin/blob/images/land.png" width="650">
  
  

Portrait ratio.  

<img src="https://github.com/ykhr53/insta-margin/blob/images/port.png" width="650">
  

Thanks, all photos are mine 😊  
https://www.flickr.com/people/yukihira_fl/