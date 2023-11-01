# CSV-to-iCal

## 参考资料

RFC: <https://datatracker.ietf.org/doc/html/rfc5545>

## 实现细节

1. 支持远程 / 本地 csv 文件
2. 支持农历事件
3. [UID](https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.7) 生成策略
4. 日期标签：公历生日（计算周岁）、农历生日（计算虚岁）、公历日期、农历日期
5. 支持多语言 i8n

## 参数

1. 过去日程：2周、1个月、3个月（默认）、6个月
2. 未来日程：3年、5年（默认）、10年
3. 语言：中文（默认）、英文

## UID 生成策略

`姓名-日期-标签@主机地址`
