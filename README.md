# Golang Logging Service  

A lightweight logging service built in Golang that logs messages to both the console and a file. The service supports multiple log levels (INFO, WARN, DEBUG, ERROR) and automatically creates the necessary log directories.  

## Features  

- Logs messages with timestamps and severity levels  
- Supports logging to both terminal and file  
- Automatically creates a log directory if not present  
- Runs as a continuous logging service  

## Project Structure  

go-log-service/ │-- Dockerfile │-- main.go │-- logs/ (Auto-created log folder)
