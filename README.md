# Red Team Examples
Red team examples and sample code.  
  
### 1. AWS Account Exploit  
  
This script performs AWS CLI commands to simulate offensive security actions:   
  - lists EC2 instances, security groups, secrets for each region. 
  - creates a new AWS IAM user eg. 'baduser' with power user privileges.    
  - using 'baduser', creates an S3 bucket, and puts a file into it.  
  - cleans up: deletes the bucket, deletes 'baduser'.
  - accepts an optional argument for the username that gets created, otherwise defaults to `baduser`. 
  
##### Prerequisites
  - AWS CLI with profile set. 

##### Usage
Arguments: newusername awsprofile. 
  
```bash 
cd aws_account_exploit
sh ./aws-lateral-movement.sh baduser default
```
  
### 2. Exfiltration Simulation
  
This example pulls data down some random endpoints and then attempts to post it to pastebin. The script does not do anything inherently malicious but is the type of activity indicating data exfiltration.  
  
##### Prerequisites
  - AWS EC2 Linux with Lacework Agent
  - Golang (ag. apt install golang)
  
##### Usage
```bash
cd GoExfil
go build
go run BasicExfil.go
```
  

