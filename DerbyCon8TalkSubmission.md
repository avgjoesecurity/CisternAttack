### --- DerbyCon 2018 Submission ---

Name(s) of speaker(s) *
Provide your name(s), these will be printed in the handout and on the website unless notice is given
Twitter Handle(s) of speaker(s)
Example: Karl - @dorkultra , Bob Speaker - @bobspeaks ....(clearly showing name association with twitter handle, separating multiple speakers by commas). This info will be published on the handout and website unless notice is given

Joe Tegg - @AvgJoeSecurity
Fraq - @0xfraq (née @BrenBriggs)

Speaker(s) Bio *
Provide a brief bio for the Speaker(s)

Joe has been involved in information technology and security for over 20 years. As a Florida native, security consultant, experienced technical cave diver, and father, his demeanor and delivery can be thought provoking, educating and hilariously awkward at the same time.

Fraq


Talk Title *

The Cistern Attack: Exploiting absolute trust

Talk Description *

We’ve all seen the devastating impact of the Watering Hole attacks on enterprise networks. The attackers watch the victim(s), learn their behavior and trick them using their own behavior to gain access or escalation. We have mimicked those techniques and found a method to do something similar against common infrastructure management and security assessment tools as an attacker which has compromised an asset in a network. Using the behavior of the existing enterprise infrastructure against itself, we easily identify and collect administrative level credentials for several classes of assets (windows, unix, network, FWs, etc) just by deploying a small set of tools. This presentation will go into detail on the attack methodology, why it works, what is/could be impacted, and finally the controls and mitigations to reduce the attack surface area. The source code for the tools we built which perform this attack will also be released for the presentation.


Talk Outline *

Intro / Problem Statement (3 mins) - Security and infrastructure scanning tools are an integral part of a mature Infosec and Information Technology operational framework. In order to identify, verify and assess the assets, credentials with domain administrator level access are required (by documentation) and utilized. Many companies blindly provide the tools with credentials which work across all assets (or large groups of assets). The security and infrastructure tools then use common methods for identifying when a credential is supplied and to what assets. This is where the Cistern Attack comes into play.

Types of credentials used in the infrastructure tools and why (5 mins) - Review the common credentials and levels of access used and even required by the most popular security and infrastructure tools. Windows Local Admin, Windows Domain Admin, SSH, SNMP

Attack the Tools (15 mins) - We demo the attack method for 3 protocols, detailing the code used, the conversation taking place, and the creds collected.

Common security and infrastructure scanning tools/processes  (7 mins) - We list out the tools that we have tested against and some that we can extrapolate may be impacted (several popular infrastructure discovery tools, the top 3 vuln scanners). We go into a moderate level of detail on how many of the tools function when dealing with assets and why they are vulnerable to attack.

Reduce the attack surface (10 mins) - Talk about the controls and tool configurations that can limit or defeat this attack method.

Here’s the goods, too. - (5 mins) We talk about our simply written tools, and tell the audience where to get the source. Then Q/A

Provide a category for your talk *
Ex: password cracking, social engineering, phishing, blue team, etc
Red Team, Exploitation, Attack Method

Has this talk been given before? If so.. Where?
Let us know if and where this talk was given before
No

Talk Length *
How long is your talk? Stable talks are 30 minutes, normal talks are 45. Please note that we reserve the right to change talk times based on available time slots and variety of content.
45 Minutes (Standard Talk)


### --- DerbyCon 2018 Rejection ---

From: David Kennedy<davek@derbycon.com>
Date: On Wed, Jul 25, 2018 at 15:20
Subject: Fwd: DerbyCon Submission - Not this year.
To: DerbyCon 2018 <info@derbycon.com>
Cc:

Greetings,

 

This is always the toughest time of the year for us because we truly had some amazing submissions. Unfortunately, your talk did not get accepted this year. We wish we had enough space to accommodate everyone because out of all of the years - this year felt like the talk topics and the research and experiences you all had was above anything we've seen before.  This year we had a blind CFP review process similar to what we do every year however we engaged multiple folks with diverse backgrounds to help with the selection process.

 

It was tough for the CFP review board and for the DerbyCon founders as well. The review board had some excellent feedback on what they saw that may help you in the future:

 

 

https://twitter.com/0xAmit/status/1017078080324677633


http://www.leeholmes.com/blog/2018/07/16/creating-a-good-security-conference-cfp-submission/


https://tisiphone.net/2018/07/16/i-reviewed-600-call-for-paper-submissions-and-youll-probably-guess-what-happened-next/

 

We'll have an additional post from our side soon that can hopefully shed some light on what we look for, how we select, and some of the top ones that were selected that can push yours over the edge next time.

 

We hope you continue to submit to DerbyCon and do hope to see you there this year as well.

 

Sincerely,

 

The DerbyCon Team
