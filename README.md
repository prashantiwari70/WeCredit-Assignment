# Golang Logging Service  

A lightweight logging service built in Golang that logs messages to both the console and a file. The service supports multiple log levels (INFO, WARN, DEBUG, ERROR) and automatically creates the necessary log directories.  

## Features  

- Logs messages with timestamps and severity levels  
- Supports logging to both terminal and file  
- Automatically creates a log directory if not present  
- Runs as a continuous logging service  

## Project Structure  

go-log-service/ │-- Dockerfile │-- main.go │-- logs/ (Auto-created log folder)


## Installation & Setup  

### Clone the Repository  
`git clone https://github.com/Sumitpal21/wecredit.git && cd go-log-service`  

### Install Golang (if not installed)  
Download from https://go.dev/dl/ and verify with `go version`  

### Run the Service  

For Windows (Direct Execution):  
`go run main.go`  

For Windows (Build & Run Executable):  
`go build -o logservice.exe main.go && ./logservice.exe`  

## Running with Docker  

### Build Docker Image  
`docker build -t go-log-service .`  

### Run Docker Container  
`docker run -d --name log-service go-log-service`  

## Log File Location  
Logs are stored in `C:\Users\palsu\OneDrive\Desktop\go-log-service\logs\app.log`  

## Environment Variables (Optional)  
`LOG_FILE` -> Default: logs/app.log (Path where logs will be stored)  
`LOG_LEVEL` -> Default: INFO (Default logging level)  

## Stopping the Service  

For normal execution:  
`CTRL + C`  

For Docker container:  
`docker stop log-service && docker rm log-service`  



  


