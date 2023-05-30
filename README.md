# Process Monitor

Process Monitor is a Go application that monitors processes and alerts on high CPU usage. It collects process metrics, processes the data, and sends email alerts when CPU usage exceeds a specified threshold.

## Features

- Executes the `top` command to collect process data
- Parses the output to extract process metrics (PID, CPU usage, command)
- Alerts via email when CPU usage exceeds a threshold
- Configurable email settings (sender, recipient, SMTP server)

## Prerequisites

Before running the application, ensure you have the following prerequisites:

- Go 1.16 or later installed
- Access to the `top` command (usually available on Unix-like systems)
- Environment variables:

  📓 `OUTPUT_FILE`: Path to the output file for the `top` command
  
  📓 `EMAIL`: Sender email address for alerts
  
  📓 `PASSWORD`: Password for the sender email address
  
  📓 `TO_EMAIL`: Recipient email address for alerts
  
  📓 `SMTP_HOST`: SMTP server hostname
  
  📓 `SMTP_PORT`: SMTP server port

## Getting Started

1. Clone the repository:

   ```ruby
   git clone https://github.com/your-username/process-monitor.git
   ```
2. Build the application:

```ruby
  go build -o monitoring main.go process.go collect.go alert.go fileUtil.go
```
3. Run the application:
```ruby
   ./monitoring
```
The application will continuously monitor the processes and send email alerts when CPU usage exceeds the threshold.

## Configuration
The following environment variables are used for configuration:

📓 `OUTPUT_FILE`: Path to the output file for the top command. Default: output.txt.

📓 `EMAIL`: Sender email address for alerts.

📓 `PASSWORD`: Password for the sender email address.

📓 `TO_EMAIL`: Recipient email address for alerts.

📓 `SMTP_HOST`: SMTP server hostname.

📓 `SMTP_PORT`: SMTP server port.

Ensure you set these environment variables correctly before running the application.

## Acknowledgements
A heartfelt thank you to all contributors and developers who have made this project possible!

If you have any questions or suggestions, please feel free to reach out.

Thank you for using Process Monitor!