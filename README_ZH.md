# CSV-to-iCal

[![Actions Status](https://github.com/fantasticmao/csv-to-ical/workflows/ci/badge.svg)](https://github.com/fantasticmao/csv-to-ical/actions)
[![codecov](https://codecov.io/gh/fantasticmao/csv-to-ical/branch/main/graph/badge.svg)](https://codecov.io/gh/fantasticmao/csv-to-ical)
[![Docker Hub](https://img.shields.io/badge/docker_hub-released-blue.svg?logo=docker)](https://hub.docker.com/r/maomao233/csv-to-ical)
![Go Version](https://img.shields.io/github/go-mod/go-version/fantasticmao/csv-to-ical)
[![Go Report Card](https://goreportcard.com/badge/github.com/fantasticmao/csv-to-ical)](https://goreportcard.com/report/github.com/fantasticmao/csv-to-ical)
[![Release](https://img.shields.io/github/v/release/fantasticmao/csv-to-ical)](https://github.com/fantasticmao/csv-to-ical/releases)
[![License](https://img.shields.io/github/license/fantasticmao/csv-to-ical)](https://github.com/fantasticmao/csv-to-ical/blob/main/LICENSE)

README [English](README.md) | [中文](README_ZH.md)

## 这是什么

CSV-to-iCal 是一个基于 Go 语言的可快速部署的 Web 服务，用于将 [CSV](https://datatracker.ietf.org/doc/html/rfc4180) 格式的内容转化成 [iCal](https://datatracker.ietf.org/doc/html/rfc5545) 格式的在线订阅链接。

CSV-to-iCal 当前支持以下特性：

- 支持本地 csv 文件 / 远程 csv 文件（通过 HTTP 协议访问）
- 支持中国农历事件
- 日期标签：公历日期、农历日期、公历生日（计算周岁）、农历生日（计算虚岁）
- 支持 i8n：英文、中文

[csv-to-ical.fantasticmao.cn](https://csv-to-ical.fantasticmao.cn) 是一个实际运行于我家 HomeLab 中的 CSV-to-iCal 服务（通过 Cloudflare Tunnel 暴露到公网），我会尽量保障它的可用性，你可以直接使用它，例如通过 <https://csv-to-ical.fantasticmao.cn/remote?url=https://raw.githubusercontent.com/fantasticmao/csv-to-ical/main/csv/testdata/calendar_test.csv> 可以转化和订阅 [calendar_test.csv](csv/testdata/calendar_test.csv) 中的日历事件。

## 快速开始

### 下载安装

### 使用示例

参数：

- [x] 语言：英文（默认）、中文
- [x] 重复日程：5年（默认），0年 ～ 10年

## 常见的问题和回答

## 许可声明

CSV-to-iCal [License](https://github.com/fantasticmao/csv-to-ical/blob/main/LICENSE)

Copyright (c) 2023 fantasticmao
