# Go Monitor Slack

Go Monitor Slack is a Go application that monitors a specified port and sends notifications to Slack if the port is not reachable.

## Features

- Monitors a configurable TCP port.
- Sends Slack notifications when the port is down or comes back up.
- Configurable check interval.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go (version 1.16 or higher)
- Git

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/bas/go-moniter-slack.git
   cd go-moniter-slack
   ```

2. Build the application:

   ```bash
   go build -o go-moniter-slack
   ```

### Configuration

The application can be configured using environment variables or a `.env` file.

| Environment Variable | Description                                   | Default Value |
| :------------------- | :-------------------------------------------- | :------------ |
| `PORT_TO_MONITOR`    | The TCP port to monitor.                      | `8080`        |
| `CHECK_INTERVAL`     | The interval (in seconds) between checks.     | `10`          |
| `SLACK_WEBHOOK_URL`  | The Slack webhook URL for notifications.      | `(required)`  |
| `SERVICE_NAME`       | The name of the service being monitored.      | `My Service`  |

Example `.env` file:

```
PORT_TO_MONITOR=80
CHECK_INTERVAL=5
SLACK_WEBHOOK_URL=https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX
SERVICE_NAME=Web Server
```

### Running the Application

1. **Using environment variables:**

   ```bash
   PORT_TO_MONITOR=80 SLACK_WEBHOOK_URL=YOUR_SLACK_WEBHOOK_URL ./go-moniter-slack
   ```

2. **Using a `.env` file:**

   Create a `.env` file in the same directory as the executable and then run:

   ```bash
   ./go-moniter-slack
   