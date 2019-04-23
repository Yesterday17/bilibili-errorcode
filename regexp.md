## 用到的正则替换

### common_ecode.go

From:
```
^\t(\w+) * = add\((-\d+)\) +// ?([^\n]+)$
```

To:
```
case $2:\nresult.Message = "$1"\nresult.Description = "$3"\nbreak
```

### main_ecode.go

From:
```
^\t(\w+) * = New\((\d+)\) // ?([^\n]+)$
```

To:
```
case $2:\nresult.Message = "$1"\nresult.Description = "$3"\nbreak
```