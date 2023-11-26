# CSV-to-iCal

[![Actions Status](https://github.com/fantasticmao/csv-to-ical/workflows/ci/badge.svg)](https://github.com/fantasticmao/csv-to-ical/actions)
[![codecov](https://codecov.io/gh/fantasticmao/csv-to-ical/branch/main/graph/badge.svg)](https://codecov.io/gh/fantasticmao/csv-to-ical)
[![Docker Hub](https://img.shields.io/badge/docker_hub-released-blue.svg?logo=docker)](https://hub.docker.com/r/maomao233/csv-to-ical)
![Go Version](https://img.shields.io/github/go-mod/go-version/fantasticmao/csv-to-ical)
[![Go Report Card](https://goreportcard.com/badge/github.com/fantasticmao/csv-to-ical)](https://goreportcard.com/report/github.com/fantasticmao/csv-to-ical)
[![Release](https://img.shields.io/github/v/release/fantasticmao/csv-to-ical)](https://github.com/fantasticmao/csv-to-ical/releases)
[![License](https://img.shields.io/github/license/fantasticmao/csv-to-ical)](https://github.com/fantasticmao/csv-to-ical/blob/main/LICENSE)

## 特性

- [x] 支持本地 / 远程 csv 文件
- [x] 支持农历事件
- [x] [UID](https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.7) 生成策略: `姓名-日期-标签@主机地址`
- [x] 日期标签：公历日期、农历日期、公历生日（计算周岁）、农历生日（计算虚岁）
- [x] 支持 i8n

## 参数

- [x] 语言：英文（默认）、中文
- [x] 重复日程：5年（默认），0年 ～ 10年

## 参考资料

RFC: <https://datatracker.ietf.org/doc/html/rfc5545>
