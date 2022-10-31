### ShmooCon 2018 Submission

== Title of Presentation ==
Portable Cistern Attack: Deceive the Infrastructure, Get the Creds

== Presenters Name ==
Joe Tegg - @AvgJoeSecurity
Ben Glass 

== Abstract ==
Keywords: Attack, Credential Harvesting, Vuln Scanners, Infrastructure tools

We’ve all seen the devastating impact of the Watering Hole attacks on victims. The attackers watch the victim, learn their behavior and trick them using their own behavior to capture creds (or install malware). Ben and I have mimicked those techniques and found a method to do something similar against standard infrastructure security and discovery tools once an attacker has completed an initial compromise (non-privileged, or privileged) in an Enterprise network. Using the behavior of the existing enterprise infrastructure against itself, we easily identify and collect administrative level credentials for several classes of assets (windows, unix, network, FWs, etc) just by deploying a small package of tools.

The presentation will go into detail on the attack methodology, why it works, what is/could be impacted, and finally the controls and mitigations to reduce the attack surface area. The source code for the tools we built which perform this attack will also be released for the presentation.

== Bio ==
Joe has been involved in information technology and security for over 20 years. As a Florida native, security consultant, experienced technical cave diver, and father, his demeanor and delivery can be thought provoking, educating and hilariously awkward at the same time.

Ben has held a variety of roles over the years, IT and contract audit, programming, DBA, infosec administration and consultant, and vulnerability management. He may not be a SME at any one topic, but he goes about his life like anyone else would… learning a new programming language simply because he’s tired of whatever he is currently using.

== Detailed Outline ==
Intro / Problem Statement (3 mins) - Security and infrastructure scanning tools are an integral part of a mature Infosec and Information Technology operational framework. In order to identify, verify and assess the assets, credentials with local administrator (and many times domain admin) level access are required and utilized. Many companies blindly provide the tools with credentials which work across all assets (or large groups of assets). The security and infrastructure tools then use common methods for identifying when a credential is supplied and to what assets. This is where the Cistern Attack comes into play.

Common security and infrastructure scanning tools/processes  (3 mins) - We list out the tools that we have tested against and some that we can extrapolate may be impacted (The top 3 vuln scanners, several infrastructure discovery tools). We go into a quick high level of how many of them work.

Types of credentials used in the tools and why (3 mins) - Review the common credentials and levels of access used and even required by the most popular security and infrastructure tools. Windows Local Admin, Windows Domain Admin, SSH, SNMP

Deceive the tools and Get the Creds (7 mins) - We demo the attack method, detail the cred collection. Then talk about the controls and configurations that can defeat this attack.

Here’s the tools, too. - (2 mins) We talk about the simple tools, and tell the audience where to get the source.

Buffer / Quick Q/A (2 mins) - Self explanatory 

== Track Preference ==
BRING IT ON, Either timeframe. This is condensed for the 20 min talk, if we get 50 min, we will have more time to dive deeper into the details and add a some time for real Q/A.

== Why is this a good fit for ShmooCon ==
We feel that topics around the use of security and infrastructure management tools should be handled with care, just like an external facing web application. So many companies fall into the trap of blind trust on these types of tools forgetting that they usually hold the keys to their kingdoms. This talk is to address this complacency issue and remind teams to be cautious and use proper controls with ALL tools where elevated credentials are utilized.

== Other Conferences submitted/presented ==
A reduced version of this attack was submitted and not presented to DerbyCon 2017 under a completely different presentation method, it wasn’t even considered an attack method at that point.

== Previous Experience ==
Both of us have never spoken at ShmooCon
Joe did a Stable talk at DerbyCon 2016

== Proceedings ==
Yes. 

== List of facilities requested ==
We just need access to project from one laptop, we will do slides and demo from that device. If this is an issue, let us know, we’re flexible.


### ShmooCon 2018 Confirmation

From: SC2018 <XXXXXXXXXX@shmoocon.org>
Subject: [SC2018] Submission ID 48
Date: October 14, 2017 at 11:26:02 AM EDT
To: joe@tegg-inc.com

Thank you for your submission to SC2018.  Below is a copy of the information submitted for your records.

Submission ID: 48

Title: Portable Cistern Attack: Deceive the Infrastructure, Get the Creds

Author 1:
   First Name: Joe
   Last Name: Tegg
   Organization:
   Country: United States
   Email: joe@tegg-inc.com
   Twitter: https://twitter.com/AvgJoeSecurity


Author 2:
   First Name: Ben
   Last Name: Glass
   Organization:
   Country: United States
   Email: glassjben@gmail.com
   Twitter:


Contact Author: Author 1

Alternate Contact:

Topic(s):
   - Bring It On
   - Bring It On - 20


Keywords: Attack, Credential Harvesting, Vuln Scanners, Infrastructure tools

Abstract: We’ve all seen the devastating impact of the Watering Hole attacks on victims. The attackers watch the victim, learn their behavior and trick them using their own behavior to capture creds (or install malware). Ben and I have mimicked those techniques and found a method to do something similar against standard infrastructure security and discovery tools once an attacker has completed an initial compromise (non-privileged, or privileged) in an Enterprise network. Using the behavior of the existing enterprise infrastructure against itself, we easily identify and collect administrative level credentials for several classes of assets (windows, unix, network, FWs, etc) just by deploying a small package of tools.

The presentation will go into detail on the attack methodology, why it works, what is/could be impacted, and finally the controls and mitigations to reduce the attack surface area. The source code for the tools we built which perform this attack will also be released for the presentation.

Comments:



File: uploaded
