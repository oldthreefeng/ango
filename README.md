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
$ go mod download

# Linux
$ make linux
# darwin
$ make darwin

$ ./ango
ango is cli tools to running Ansible playbooks from Golang.
run "ango -h" get more help, more see https://github.com/oldthreefeng/ango
ango version, :      1.0.0
Git Commit Hash:     51c3c6e
UTC Build Time :     2019-12-12 12:12:41 UTC
Go Version:          go version go1.13.4 linux/amd64
Author :             louis.hong
```

use go get 

```bash
$ go get  github.com/oldthreefeng/ango
```

### to run palybook

```bash
$ export DingDingMobiles="158****6468"
$ export DingDingUrl="https://oapi.dingtalk.com/robot/send?access_token=*****"
$ ango deploy -h 
use ango to deploy project with webhook to dingding

Usage:
  ango deploy [flags]

Examples:
  ango deploy -f test.yml -t v1.2.0

Flags:
  -h, --help   help for deploy

Global Flags:
  -a, --author string   author name for copyright attribution (default "louis.hong")
  -f, --filename string   ansible-playbook for yml config
  -t, --tag string      tags for the project version

$ ango rollback -h
rollback 回退版本, 需要指定回退版本的yml文件及要回退的version

Usage:
  ango rollback [flags]

Examples:
  ango rollback -f test.yml -t v1.2

Flags:
  -h, --help   help for rollback
  -r, --real   really to rollback this version

Global Flags:
  -a, --author string   author name for copyright attribution (default "louis.hong")
  -f, --filename string   ansible-playbook for yml config
  -t, --tags string     tags for the project version
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
