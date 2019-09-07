codec
---
- encoder
```go
xml.NewEncoder(bean).Encode(bs) // obj->bs
```

- decoder
```go
err = xml.NewDecoder(resp.Body).Decode(&document) // bs -> obj
```

regex
---
```go
matched, err := regexp.MatchString(searchTerm, channelItem.Title)
```
正则匹配，可以看看[demo](./../note/learn_regex.go)