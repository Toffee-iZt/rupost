# rupost
Russian post tracking.

```golang
resp, postErr, err := rupost.Track("AB123456789CD")
...
item := resp.TrackingItem
```