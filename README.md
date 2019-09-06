# The Cistern Attack: Exploiting Absolute Trust

We’ve all seen the devastating impact of the Watering Hole attacks on enterprise networks. The attackers watch the victim(s), learn their behavior and trick them using their own behavior to gain access or escalation. We have mimicked those techniques and found a method to do something similar against common infrastructure management and security assessment tools as an attacker which has compromised an asset in a network. Using the behavior of the existing enterprise infrastructure tool against itself, we easily identify and collect administrative level credentials for several classes of assets (windows, unix, network, FWs, etc) just by deploying a small set of tools.

## Problem Statement

Security and infrastructure scanning tools are an integral part of a mature Infosec and Information Technology operational framework. In order to identify, verify and assess the assets, credentials with domain administrator level access are required (by documentation) and utilized. Many companies blindly provide the tools with credentials which work across all assets (or large groups of assets). The security and infrastructure tools then use common methods for identifying when a credential is supplied and to what assets. This is where the Cistern Attack becomes a useful adversary tool.

### Tyes of Common Discovery/Scanning Credentials

Windows Local Admin, Windows Domain Admin, SSH, SNMP

### Vulnerable security and infrastructure tools:

Qualys Security Scanner

Rapid7 Vulnerability Scanner

Tenable Vulnerability Scanner

ServiceNow Asset Discovery

HP Asset Manager


## Attack Methhod Workflows

### Method 1: Phishing (least effective from our research)
```
Send user email link to click

Exploit vulnerable user application

Install / run payload using the code for fake services

Use monitoring wrapper to keep malware running

Listen for discovery / vulnerability scanning tool to hit the exposed serviceand provide the credential(s) to your fake service as part of their discovery/scan process.

Exfil the credential data in your malware wrapper.

Profit.
```

### Method 2: Remote Exploitation (moderately effective)
```
Exploit vulnerable service / application

Install / run payload using the code for fake services 

Use monitoring wrapper to keep malware running

Listen for discovery / vulnerability scanning tool to hit the exposed serviceand provide the credential(s) to your fake service as part of their discovery/scan process.

Exfil the credential data in your malware wrapper.

Profit.
```

### Method 3: BlackBox Drop (highest effectiveness)
```
Build a custom "BlackBox" to plug in and drop on-site on the network, or build a custom docker container to spin up on a compromised host.

Run fake services for SMB, SSH, and SNMP.

Listen for discovery / vulnerability scanning tool to hit the exposed serviceand provide the credential(s) to your fake service as part of their discovery/scan process.

Exfil the credential data.

Profit.
```


## Attack Tools and Useage

SMB 
```
Use SMB Responder
```
SSH Listener 
```
ssh_server.go
```
SNMP Listener
```
snmptrap.go
```


### SMB Credentials (Windows Local Admin or Domain Admin)
Port: 139

In our reseach, SMB was the hardest protocol to replicate with a Go script. There are handshakes that take place which caused a fair amount of work and we decided that at this time it was easier to use SMB Responder and write a wrapper for it. If anyone wants to build a cool tool for this mehod of credential collection to go with our other Go tools, please feel free to do so.

We also identified that Windows does not allow you to unbind SMB and take it over, thus if you're on a Windows compromised asset, you will have a harder time collecting SMB credentials. You can however get SSH and SNMP credentials.

If you compromise a linux box, or a cutom made "BlackBox" SMB credentials are trivial to collect using this method.


### SSH Credentials (User / Pass - usually root equivilent or SUDO capability)
Port: ANY common port

In our research, SSH username/pass was really simple to acquire with our ssh_server.go tool. You can spin up the listener on ANY common port which is used to discover assets and the scanning tools will connect and authenticate with the service so long as it "looks" like a true SSH service. This is becuase part of what these tools do is to make a "guess" of the service and then apply any availible credentials to it.



### SNMP (Private Community String)
Port: ANY common port

In our research, SNMPv1/SNMPv2 is still heavily used in Enterprise environments. These are weak versions of the protocol and allow us to easily gather the credentials to interact with devices configured to allow authentication with them. This tends to be network devices and even some FWs. 


## Enterprise Mitigation Options

### Can you prevent this attack methodology from affecting your network?

Kinda.

Many of the vulnerability and asset discovery tools do NOT use good controls around how they present credentials to "identified services".

Serveral allow SOME restrictions on the use of credentials, but their usefulness is limited. One vendor tool we tested allows you to limit credentail use in the configuration, but that control was found to be ineffective and did not work as intended. (broken?)

### So how the heck can we prevent people like you from doing this on my network?

1. Configure asset discovery tools to NOT authenticate, just find the assets and fingerprint them.

2. Encourage the tool vendors to address this long known weakness in their tool capability by offering the following controls:
   
   * Limit the port to the specific TCP/UDP ports that are expected, no other ports
   * Limit the exposure of certain credentials to a "known" asset group. (Don't expose Windows (SMB) credentials to fingerprinted Linux devices, don't use SNMP on Windows devices, etc)
   * Give administrators MORE control on where credentials are utilized, period.

3. Use an integration with a IAM tool (secrets management) for one time passwords on scanned assets. Ensure that this integration can scale and is effective.

4. Use security automation tools to review and verify that scanning credentials are only being used DURING SCANS. This is so easy but NO ONE does it, in the research that we have done. Hell, most places don't even monitor their credentials at all....

## Authors / Researchers / Contributors
Joe Tegg
Ben Glass
Troy Newcomb
Bren Briggs

## License

Just give credit where credit is due. This has been something in the works for a LONG time and was strictly used in testing activities but we thought it should be shared because of the impact to well known asset discovery and security tools that share high levels of access and have crappy recomendations for credential use.

## TIMELINE
Please see the TIMELINE.md file for the initial finding, research, and public talk submissions (trying to make this public).
