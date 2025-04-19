# Missile Go

Missile Go is a terminal-based application that allows you to control Satzuma missile devices via USB or over a network. The application supports multiple modes of operation, including a graphical display mode, a console mode for direct command input, and a network mode for remote control.

This is a rewrite of: https://github.com/momentofgeekiness/pymissile-ng 

## Features

- Control missile devices using keyboard commands.
- Supports both legacy and center missile devices.
- Operate in different modes:
  - **Display Mode**: A terminal-based user interface for controlling missiles.
  - **Console Mode**: Direct command input from the console.
  - **Network Mode**: Receive commands over a UDP network.

## Installation

To install the project, clone the repository and navigate to the project directory:

```bash
git clone git@github.com:wieringen/satzuma-missile-go.git
cd satzuma-missile-go
```

Then, run the following command to download the necessary dependencies:

```bash
go mod tidy
```

## Usage

You can run the application in different modes using command-line arguments:

- To run in display mode:

```bash
go run cmd/main.go --display
```

- To run in console mode:

```bash
go run cmd/main.go --console
```

- To run in network mode:

```bash
go run cmd/main.go --network
```

## License

This project is licensed under the MIT License. See the LICENSE file for more details.