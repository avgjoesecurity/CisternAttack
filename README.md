# The Cistern Attack: Exploiting Absolute Trust

We’ve all seen the devastating impact of the Watering Hole attacks on enterprise networks. The attackers watch the victim(s), learn their behavior and trick them using their own behavior to gain access or escalation. We have mimicked those techniques and found a method to do something similar against common infrastructure management and security assessment tools as an attacker which has compromised an asset in a network. Using the behavior of the existing enterprise infrastructure against itself, we easily identify and collect administrative level credentials for several classes of assets (windows, unix, network, FWs, etc) just by deploying a small set of tools. This presentation will go into detail on the attack methodology, why it works, what is/could be impacted, and finally the controls and mitigations to reduce the attack surface area. The source code for the tools we built which perform this attack will also be released for the presentation.

## Problem Statement

Security and infrastructure scanning tools are an integral part of a mature Infosec and Information Technology operational framework. In order to identify, verify and assess the assets, credentials with domain administrator level access are required (by documentation) and utilized. Many companies blindly provide the tools with credentials which work across all assets (or large groups of assets). The security and infrastructure tools then use common methods for identifying when a credential is supplied and to what assets. This is where the Cistern Attack comes into play.

Tyes of Common Discovery/Scanning Credentials

Windows Local Admin, Windows Domain Admin, SSH, SNMP


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
