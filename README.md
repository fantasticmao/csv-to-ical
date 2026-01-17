# CSV-to-iCal

[![Actions Status](https://github.com/fantasticmao/csv-to-ical/workflows/ci/badge.svg)](https://github.com/fantasticmao/csv-to-ical/actions)
[![codecov](https://codecov.io/gh/fantasticmao/csv-to-ical/branch/main/graph/badge.svg)](https://codecov.io/gh/fantasticmao/csv-to-ical)
[![Docker Hub](https://img.shields.io/badge/docker_hub-released-blue.svg?logo=docker)](https://hub.docker.com/r/maomao233/csv-to-ical)
![Go Version](https://img.shields.io/github/go-mod/go-version/fantasticmao/csv-to-ical)
[![Go Report Card](https://goreportcard.com/badge/github.com/fantasticmao/csv-to-ical)](https://goreportcard.com/report/github.com/fantasticmao/csv-to-ical)
[![Release](https://img.shields.io/github/v/release/fantasticmao/csv-to-ical)](https://github.com/fantasticmao/csv-to-ical/releases)
[![License](https://img.shields.io/github/license/fantasticmao/csv-to-ical)](https://github.com/fantasticmao/csv-to-ical/blob/main/LICENSE)

README [English](README.md) | [中文](README_ZH.md)

## What is this

CSV-to-iCal is a rapidly deployable Web application based on the Go language, used to convert [CSV](https://datatracker.ietf.org/doc/html/rfc4180) format content into [iCal](https://datatracker.ietf.org/doc/html/rfc5545) format online subscription links. For example, you can import calendar events from [calendar_test.csv](csv/testdata/calendar_test.csv) into your iOS Calendar / Google Calendar by subscribing to the URL <https://csv-to-ical.fantasticmao.cn/remote?url=https://raw.githubusercontent.com/fantasticmao/csv-to-ical/main/csv/testdata/calendar_test.csv>.

![usage](usage.png)

[csv-to-ical.fantasticmao.cn](https://csv-to-ical.fantasticmao.cn) is an out-of-the-box CSV-to-iCal service instance provided for general users, actually running on my HomeLab and exposed to the public network via Cloudflare Tunnel.

CSV-to-iCal currently supports the following features:

- [x] Fully respects user privacy, does not retain any user content, code is 100% open source and transparent
- [x] Supports reading local CSV files, supports reading remote CSV files via HTTP protocol
- [x] Supports Gregorian and Lunar calendar events
- [x] Supports calculating the age for Gregorian birthday events, supports calculating the nominal age for Lunar birthday events
- [x] Supports English (default) and Chinese languages
- [x] Supports recurring events for up to 5 years
- [x] Supports retrospective events for up to 3 years

## Download and Install

<details>
<summary>Compile from source code</summary>

Clone the repository:

```bash
git clone https://github.com/fantasticmao/csv-to-ical.git
cd csv-to-ical
```

Start compiling:

```bash
go build -o csv-to-ical .
```

This will generate an executable file named `csv-to-ical` in the current directory.

</details>

<details>
<summary>Use go install</summary>

If you have a Go environment, you can install directly using `go install`:

```bash
go install github.com/fantasticmao/csv-to-ical@latest
```

This will install the `csv-to-ical` command to your `GOPATH/bin` directory.

</details>

<details open>
<summary>[Recommended] Use Docker</summary>

The project provides a Docker image, which can be deployed via Docker:

```bash
docker pull maomao233/csv-to-ical:latest
```

Then you can run the container:

```bash
docker run -d -p 7788:7788 -v /path/to/your/config/:/opt/csv-to-ical/ maomao233/csv-to-ical:latest
```

Please make sure to replace `/path/to/your/config` with the actual directory where your `config.yaml` is located. The `config.yaml` storage directory inside the container is `/opt/csv-to-ical`.

</details>

## Quick Start

<details open>
<summary>Configure config.yaml</summary>

Create a configuration file `config.yaml`. The program will by default look for `config.yaml` in `.config/csv-to-ical/` under the user's home directory. You can also specify the configuration directory using the `-d` command line parameter.

Below is an example `config.yaml`:

```yaml
bind-address: 0.0.0.0:7788 # Program listening address and port, default is 0.0.0.0:7788

http-client:
  timeout: 3000 # HTTP client timeout (milliseconds), default is 3000ms
  proxy: ""     # Optional HTTP proxy address, e.g., http://127.0.0.1:7890

csv-providers:
  # 'remote' provider example, used to process remote CSV files
  my-remote-calendar:
    url: "https://raw.githubusercontent.com/fantasticmao/csv-to-ical/main/csv/testdata/calendar_test.csv"
    language: "zh-cn" # Event language, optional "en" (default) or "zh-cn"
    recurCnt: 5       # Maximum number of years for recurring events, default is 3, max 5
    backCnt: 2        # Maximum number of years for retrospective events, default is 1, max 3

  # 'local' provider example, used to process local CSV files
  my-local-calendar:
    file: "/path/to/your/calendar.csv" # Local CSV file path
    language: "en"
    recurCnt: 3
    backCnt: 1
```

</details>

<details open>
<summary>Start CSV-to-iCal</summary>

Run in the project root directory:

```bash
go run main.go
```

Or, if you have installed the binary:

```bash
csv-to-ical -d /path/to/your/config
```

</details>

<details open>
<summary>Access CSV-to-iCal</summary>

After the program starts successfully, it will output the following log. You can then access CSV-to-iCal via the address configured in `bind-address`.

```
start HTTP server success, bind address: 0.0.0.0:7788
```

</details>

<details open>
<summary>Get iCal subscription link</summary>

By visiting <http://0.0.0.0:7788/remote?url=...> or <http://0.0.0.0:7788/local/my-local-calendar>, you can get the iCal subscription link.

</details>

## Usage Instructions

**CSV File Format**

The CSV file format has specific requirements and should contain the following columns:

| Column Name (case-insensitive) | Description   | Format                                                                                                            | Example                            |
|--------------------------------|---------------|-------------------------------------------------------------------------------------------------------------------|------------------------------------|
| Name                           | Event Name    | String                                                                                                            | Xiao Ming's Graduation Anniversary |
| Month                          | Month         | Integer (1-12)                                                                                                    | 6                                  |
| Day                            | Day           | Integer (1-31)                                                                                                    | 1                                  |
| Year                           | Year          | Integer                                                                                                           | 2022                               |
| Calendar_Type or CalendarType  | Calendar Type | Type enumeration (solar Gregorian, lunar Lunar, birthday_solar Gregorian Birthday, birthday_lunar Lunar Birthday) | solar                              |

**CSV File Example**

For details, please see [calendar_test.csv](csv/testdata/calendar_test.csv).

**Using Remote CSV Files**

You can subscribe to a publicly accessible CSV file via the `remote` interface. For example, if you have a public CSV file link <https://example.com/your-calendar.csv>, you can use the following URL to access it:

```
https://<your-csv-to-ical-host>/remote?url=https://example.com/your-calendar.csv
```

- `url`: The full URL of the remote CSV file.
- `lang` (optional): Event language, `en` (default) or `zh-cn`.
- `recurCnt` (optional): Maximum number of years for recurring events, default is 3, max 5.
- `backCnt` (optional): Maximum number of years for retrospective events, default is 1, max 3.

For example, to subscribe to a Chinese recurring event, repeating for a maximum of five years, and retrospecting for a maximum of two years:

```
https://<your-csv-to-ical-host>/remote?lang=zh-cn&recurCnt=5&backCnt=2&url=https://example.com/your-calendar.csv
```

**Using Local CSV Files**

If you want to use local CSV files, you need to configure them via `config.yaml`. Add an entry in the `csv-providers` section, specifying the `file` path.

For example, configure in `config.yaml`:

```yaml
csv-providers:
  my-local-calendar:
    file: "/path/to/your/calendar_example.csv"
    language: "zh-cn"
    recurCnt: 5
    backCnt: 2
```

Then, you can access it via the following URL:

```
https://<your-csv-to-ical-host>/local/my-local-calendar
```

Here, `my-local-calendar` corresponds to the key name under `csv-providers` in `config.yaml`.

## Frequently Asked Questions

Q: Why aren't my calendar events displayed?

A: Please check if your CSV file format is correct, especially if the columns `Name`, `Month`, `Day`, `Year`, `CalendarType` exist and the data format is correct. Also, check if the `file` or `url` path in `config.yaml` is correct.

---

Q: How to enable HTTPS?

A: CSV-to-iCal itself does not directly provide HTTPS functionality. It is recommended to configure a reverse proxy (such as Nginx, Caddy, Cloudflare Tunnel, etc.) in front to handle HTTPS.

---

Q: What do `recurCnt` and `backCnt` mean?

A: `recurCnt` (recurrence count) specifies the maximum number of years for which recurring events will be generated (starting from the current year). `backCnt` (back count) specifies the maximum number of years for retrospective events (looking back from the current year).

## License

CSV-to-iCal [License](https://github.com/fantasticmao/csv-to-ical/blob/main/LICENSE)

Copyright (c) 2023 fantasticmao
