## What is this?

Converts [Starling's](https://www.starlingbank.com/) CSV bank statements to [FreeAgent's](https://www.freeagent.com/) CSV format.

FreeAgent's CSV format can be found here: https://support.freeagent.com/hc/en-gb/articles/115001222564

## Why?

I was tired of waiting for my accounting software to support my bank's international account statements, and for my bank to support CSV exports for non-UK accounts. So I made this. 

## Installation

### Binary Releases

For Windows (?), Mac OS(10.12+) or Linux, you can download a binary release [here](../../releases).

### Go

```sh
go install github.com/levriero/starling-agent@latest
```

## Usage

```sh
starling-agent ~/MyIncredibleStarlingStatement.csv
```
