---
title: "Netstat"
date: 2020-04-14T15:59:31+08:00
draft: false

tags: ["linux"]

summary: "A command which often use to check network stats"

---

## Netstat

|command  |     meanings             |
|-------- |--------------------------|
|-a       | listen port              |
|-n       | don't do dns translation |
|-t       | tdp                      |
|-p       | show process             |


### To show the process which use port 443
```
sudo netstat -antp | grep 443
```

