# iwantmyname-ddns

A command-line tool written in Go for updating DNS records using the iwantmyname DDNS (Dynamic DNS) interface.

## Features

- Update A or AAAA records.
- Add or delete other record types such as TXT, MX, CNAME, NAPTR, and SRV.

## Installation

1. Clone the repository:

`git clone https://github.com/yourusername/iwantmyname-ddns.git`

2. Change to the project directory:

`cd iwantmyname-ddns`

3. Build the project:

`go build -o iwantmyname`

This command will compile the Go program into an executable named `iwantmyname`.

## Usage

Set the environment variables `IWANTMYNAME_USERNAME` and `IWANTMYNAME_PASSWORD` with your iwantmyname login credentials.

For Bash or Zsh:

`export IWANTMYNAME_USERNAME="your_username"`
`export IWANTMYNAME_PASSWORD="your_password"`

For Fish shell:

`set -x IWANTMYNAME_USERNAME your_username`
`set -x IWANTMYNAME_PASSWORD your_password`

Run the `iwantmyname` executable with the required arguments:

`./iwantmyname hostname recordType value [ttl]`

Replace `hostname`, `recordType`, `value`, and `ttl` (optional) with the appropriate values for your use case.

## Examples

- Update an A record:

`./iwantmyname home.example.com A 192.0.2.1`

- Add a TXT record:

`./iwantmyname verify.example.com TXT "Xh90WmC95K1GAlAcEF0psz3c2GOandqdqMInclDxd"`

- Delete a TXT record:

`./iwantmyname verify.example.com TXT delete`
