# TSDNS-Server
[Archived] This TSDNS server handle and asnwear connection request from TSDNS clients.

---
## Build
GIT: `git clone https://github.com/Freaxbot/TSDNS-Server`

Install libs: `go get gopkg.in/ini.v1`

Build: `go build main.go server.go`

Output: main.exe (on WIN)

## Config
File: `config.ini`
```ini
[server]
# Set the default port for ths TDNS server.. Default :: 41144
port = 41144
[routing]
public.teamspeak.com = 12.13.14.15:10000
test.teamspeak.com = 12.13.14.15:12000
private.teamspeak.com = 12.13.14.15:14000
voice.teamspeak.com = NORESPONSE
*.teamspeak-systems.de = 1.2.3.4:15000
```

You can set a new rout with the domain and the ip:port.

|Lisening domain        | ip+port (or NORESPONSE )|
|-----------------------|-------------------------|
|`public.teamspeak.com` |`12.13.14.15:14000`      |
|`voice.teamspeak.com`  |`NORESPONSE`      |

---
[![WTFPL](https://upload.wikimedia.org/wikipedia/commons/thumb/0/0a/WTFPL_badge.svg/320px-WTFPL_badge.svg.png)](LICENSE.md)

This work is licensed under the "Do **W**hat **T**he **F**uck You Want To **P**ublic **L**icense". To see the full License Text go to the [LICENSE File](LICENSE.md)
