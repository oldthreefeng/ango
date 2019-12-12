/*
 * Copyright (c) 2019. The ango Authors. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */

package play

import (
	"fmt"
	"go.uber.org/config"
	"strings"
)

const API = `
---
# MAINTAINER louis.hong@junhsue.com
# 雅集后台
# example: ansible-playbook api.yml  -e version=$current_version -f 1
- hosts: yj-back
  serial: 2
  remote_user: root
  tasks:
  - name: Create a directory if it does not exist
    shell: "test -d /usr/local/e-mall/{{item}} || mkdir /usr/local/e-mall/{{item}}/config -p"
    with_items:
      - api
      - api1
      - api2
  - name: copy and replace version the startapi.j2 to nodes
    template:
        src: /opt/playbook/devops/script/startapi.t2
        dest: /usr/local/e-mall/{{item.name}}/startapi.sh
        mode: "u=rwx,g=r,o=r"
    with_items:
      - { name: 'api', port: '8888' }
      - { name: 'api1', port: '22332' }
      - { name: 'api2', port: '22333' }

  - name: copy api jar  to nodes
    copy:
        src: /opt/playbook/devops/weimall/weimall-api-single-{{version}}.jar
        dest: /usr/local/e-mall/{{item.name}}/weimall-api-single-{{version}}-{{item.port}}.jar
    with_items:
      - { name: 'api', port: '8888' }
      - { name: 'api1', port: '22332' }
      - { name: 'api2', port: '22333' }
  - name: extra config file to config directory
    template:
        src: /opt/playbook/devops/script/application-prod.yml
        dest: /usr/local/e-mall/{{item.name}}/config/application-prod.yml
    with_items:
      - { name: 'api', port: '8888' }
      - { name: 'api1', port: '22332' }
      - { name: 'api2', port: '22333' }
  - name: get pid of weimall-api last time
    shell: "ps -ef | grep -v grep | grep -E [w]eimall-api-single | awk '{print $2}'"
    register: running_processes
  - name: Kill running processes
    shell: "kill {{ item }}"
    with_items: "{{ running_processes.stdout_lines }}"
  - wait_for:
      path: "/proc/{{ item }}/status"
      state: absent
      timeout: 60
    with_items: "{{ running_processes.stdout_lines }}"
    ignore_errors: yes
    register: killed_processes
  - name: Force kill stuck processes
    shell: "kill -9 {{ item }}"
    with_items: "{{ killed_processes.results | select('failed') | map(attribute='item') | list }}"
  - name: start weimall-api
    shell: "cd /usr/local/e-mall/{{item}} && ./startapi.sh"
    with_items:
      - api
      - api1
      - api2
`

// TODO
// add a template to deploy

func Generate() {
	reader := strings.NewReader(API)
	var conf  = config.Source(reader)
	yml, err  := config.NewYAML(conf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(yml.Name())
}