# anteye
anteye is a small and simple monitor system. anteye should monitor cluster less then 50 instances. it can send notice msgs via **mail**、**sms**、**voice** or **callback(TODO)**.
we suggest you deploy more than one anteye instances in the production environment.

## install

You can install anteye from the latest [release](https://github.com/RosenLo/anteye/releases/download/v1.0.0/anteye-1.0.0.tar.gz),

```bash
# download release
wget -q https://github.com/RosenLo/anteye/releases/download/v1.0.0/anteye-1.0.0.tar.gz
tar -zxf anteye-$vsn.tar.gz

# config, change configs as you like
mv cfg.example.json cfg.json
vim cfg.json
...

# start
./control start

# stop
./control stop

```

Or you can install anteye from scratch

```bash
# download src
cd $GOPATH/src/github.com/RosenLo
git clone https://github.com/RosenLo/anteye.git
cd anteye
go get ./...

# build, get bin anteye
./control build

# config, change configs as you like
mv cfg.example.json cfg.json
vim cfg.json
...

# start
./control start

# stop
./control stop

```

## config
```python

debug: true/false, open debug log or not

http
    - enable: true/false, enable http-server or not
    - listen: listening port of http-server

mail
    - enable: true/false, enable sending alarm mails or not
    - url: http-url used to post mail content
    - receivers: mail accounts. if you have multiple accounts, then separate them by commas. eg. "a@gmail.com,b@yahoo.com"

sms
    - enable: true/false, enable sending alarm sms or not
    - url: http-url used to post sms content
    - receivers: mobile numbers. if you have multiple numbers, then separate them by commas. eg. "18001163876,13811685233"

voice
    - enable: true/false, enable sending alarm voice or not
    - url: http-url used to post voice content
    - receivers: mobile numbers.

callback
    - enable: true/false, enable alarm callback or not
    - url: http-url used to post alarm content

monitor
    - cluster: host instances to be monitored, one item goes like "module,hostname:port/health/url"

whiteCode
	- list of http status codes

maxStep
	- maximum number of alarms

cron
	- just like crontab in linux
```

## interface
anteye sends msgs via http interfaces. these interfaces defined as followings:

```bash
# sms interface
method: http.post
params:
  - tos: mobile numbers separated by commas
  - content: content of sms
  - from: optional, indicates who sends this sms 

# voice interface
method: http.post
params:
  - to: mobile numbers
  - content: content of voice
  - from: optional, indicates who sends this voice

# mail interface
method: http.post
params:
  - tos: mail accounts separated by commas
  - content: content of mail
  - subject: subject of mail
  - from: optional, indicates who sends this mail

# callback
method: http.post
params: body. anteye will post you a string object like '[instance] connection timeout', eg. [task,127.0.0.1:16269/health] connection timeout

```

## debug
```bash
# log
./test/debug tail

# get internal status
./test/debug counter

```

## reference
