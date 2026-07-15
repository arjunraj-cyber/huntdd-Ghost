Huntdd-Ghost is a lightweight, local-only filesystem auditing tool built in Go. It’s designed for security enthusiasts who want a quick, "no-nonsense" way to scan directories for potential misconfigurations, sensitive files, and known vulnerabilities based on file hashing.
Unlike heavy antivirus suites that look for malware, huntdd-Ghost focuses on system hygiene and visibility. It helps you identify Sensitive Enumeration, Vulnerability Audit, Permission Checks. 

This tool is purely for auditing. It does not modify, delete, or clean files. It only observes and reports!!

To get this tool you need GO installed on your machine.

git clone https://github.com/arjunraj-cyber/huntdd-Ghost.git

cd huntdd-Ghost

go run huntdd.go

To scan a specific path : 
go run huntdd.go -scan /path/to/your/files


⚠️ Disclaimer
huntdd-Ghost is provided "as is" without warranty of any kind. This tool is intended solely for educational purposes, personal security auditing, and authorized system administration. The creator, Arjun Raj assumes no responsibility for any misuse, data loss, or system instability caused by the use of this tool. Always ensure you have appropriate authorization before running security audit tools on any system.
