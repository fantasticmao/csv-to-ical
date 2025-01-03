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

CSV-to-iCal 是一个基于 Go 语言的可快速部署的 Web 服务，用于将 [CSV](https://datatracker.ietf.org/doc/html/rfc4180) 格式的内容转化成 [iCal](https://datatracker.ietf.org/doc/html/rfc5545) 格式的在线订阅链接。例如，你可以通过订阅 URL [https://csv-to-ical.fantasticmao.cn/remote?url=....../calendar_test.csv](https://csv-to-ical.fantasticmao.cn/remote?url=https://raw.githubusercontent.com/fantasticmao/csv-to-ical/main/csv/testdata/calendar_test.csv) 来将 [calendar_test.csv](csv/testdata/calendar_test.csv) 中的日程事件导入到你的 iOS 日历 / Google 日历中。

[csv-to-ical.fantasticmao.cn](https://csv-to-ical.fantasticmao.cn) 是一个为普通用户提供的开箱即用的 CSV-to-iCal 服务实例，实际运行于我家的 HomeLab，并通过 Cloudflare Tunnel 暴露到公网。

CSV-to-iCal 当前支持以下特性：

- [x] 充分尊重用户隐私，不留存任何用户内容，代码 100% 开源和透明
- [x] 支持读取本地 CSV 文件，支持通过 HTTP 协议读取远程 CSV 文件
- [x] 支持公历与农历事件
- [x] 支持计算公历生日事件的周岁年龄，支持计算农历生日事件的虚岁年龄
- [x] 支持英文（默认）、中文两种语言
- [x] 支持最大 5 年的重复日程
- [x] 支持最大 3 年的回溯日程

## 快速开始

### 下载安装

### 使用示例

## 常见的问题和回答

## 许可声明

CSV-to-iCal [License](https://github.com/fantasticmao/csv-to-ical/blob/main/LICENSE)

Copyright (c) 2023 fantasticmao
