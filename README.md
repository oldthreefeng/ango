## ango

基于golang开发的一个用于部署项目至生产环境的部署工具

目前仅使用playbook部署/回滚相关业务并使用钉钉的`webhook`通知, 文档查看: https://github.com/oldthreefeng/ango

## Required

- `go version go1.13.4 linux/amd64`
- `export GO111MODULE="on"`
- `ansible2.6+`
- `.yml is ready to go`

## Usage

### download and compile

use git to download source code

```bash
$ git clone https://github.com/oldthreefeng/ango.git
$ cd ango && go mod download

# Linux
$ make linux
# darwin
$ make darwin

$ ./ango
ango is cli tools to running Ansible playbooks from Golang.
run "ango -h" get more help, more see https://github.com/oldthreefeng/ango
ango version :       1.0.0
Git Commit Hash:     a9a3c28
UTC Build Time :     2019-12-13 04:16:36 UTC
Go Version:          go version go1.13.4 linux/amd64
Author :             louis.hong
```

use go get 

```bash
$ go get github.com/oldthreefeng/ango
```

### run with palybook

first, to config your ansible, more to see [ansible](https://github.com/ansible/ansible)

```bash
$ vim /etc/ansible/hosts
[test]
192.168.0.62
192.168.0.63
$ ansible test -m ping
192.168.0.62 | SUCCESS => {
    "changed": false, 
    "ping": "pong"
}
192.168.0.63 | SUCCESS => {
    "changed": false, 
    "ping": "pong"
}
```

second, to export some env to hook to Dingding

```bash
$ export DingDingMobiles="158****6468"
$ export DingDingUrl="https://oapi.dingtalk.com/robot/send?access_token=*****"
```

third, to deploy your project

```bash
$ ango deploy -f test.yml -v v1.23  
## It's equal to  `ansible-playbook test.yml -e version=v1.23 -f 1`
## and to Post a test Message to Dingding

$ ango deploy -h 
use ango to deploy project with webhook to dingding

Usage:
  ango deploy [flags]

Examples:
  ango deploy -f api.yml -t v1.2.0

Flags:
  -m, --comments string   add comments when send message to dingding
  -h, --help              help for deploy

Global Flags:
      --author string     author name for copyright attribution (default "louis.hong")
  -f, --filename string   ansible-playbook for yml config(requried)
  -t, --tags string       tags for the project version(requried)
  -v, --verbose           verbose mode to see more detail infomation
```
fourth, to rollback your project 

```
$ ango rollback -f test.yml -t v1.2 --real

## 
rollback 回退版本, 需要指定回退版本的yml文件及要回退的version

Usage:
  ango rollback [flags]

Examples:
  ango rollback -f roll_api.yml -t v1.2  -r 

Flags:
  -h, --help   help for rollback
  -r, --real   really to rollback this version

Global Flags:
      --author string     author name for copyright attribution (default "louis.hong")
  -f, --filename string   ansible-playbook for yml config(requried)
  -t, --tags string       tags for the project version(requried)
  -v, --verbose           verbose mode to see more detail infomation
```

### logs

```bash 
# 查看发布日志
# -r is requried when rollback
$ ango rollback  -f test.yml -t v1.2.0  -r
$ ango deploy -f test.yml -t v1.4.0
$ tail -f fabu.log
[INFO] 2019-12-12 18:36:29 test-v1.2 回滚成功
[INFO] 2019-12-12 18:37:00 test-v1.4 部署成功
```

that's all

[thanks to jetbrains](https://www.jetbrains.com/?from=ginuse)

![](https://www.jetbrains.com/company/brand/img/jetbrains_logo.png)
