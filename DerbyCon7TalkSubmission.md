(Submitted May 2017)

**== Title of Presentation ==**

Social Engineering the Vuln Scanner

**== Presenters Name ==**

Joe Tegg
Ben Glass 

**== Abstract ==**

Vulnerability scanning software is a primary part of any mature vulnerability and risk management program. These tools are provided great levels of access across networks and devices in order to present us with useful risk exposure metrics. Security teams provide this access to these tools because we "trust" them. There is just one problem, the vuln scanners handle authentication credentials like humans. They can be socially engineered to provide all of them, in several easily implemented techniques. An attacker can simply "sweet talk" their way into admin level access across networks by just sitting down in the bar and offering the scanner a drink.

**== Bio(s) ==**

Joe has been involved in information technology and security for over 20 years. Starting with a simple challenge to root a shelled server, morphing into a unix sysadmin, and diving into enterprise vulnerability management. Now, as a passionate Rapid7 senior security consultant, he has experience in all aspects of risk assessment, mitigation strategy and processes. He provides strategic insight, asks the difficult questions, and drives the conversations that Vulnerability Management teams should be having with their CISOs and CIOs. As a Florida native, urban redneck, backcountry explorer, experienced technical cave diver, and father, his demeanor and delivery can be thought provoking, educating and hilariously awkward at the same time.

**== Detailed Outline ==**

Intro / Problem Statement (7 mins) - Vulnerability scanning tools are a required part of a mature vulnerability management program. In order to identify and verify the vulnerabilities present on an asset, credentials with local administrator level access are required. Many companies provide the tools with credentials which work across all assets (or large groups of assets). The vulnerability scanners then use a common method for identifying when a credential is supplied and to what assets. This is where they walk into the bar and take a seat.

Common vuln scanning process  (7 mins) - How do the majority of vulnerability scanners scan? We explain.

Types of credentials used in vuln scanning (5 mins) - Review the common credentials and levels of access required by the most popular vuln scanners. 

Offer the vuln scanner a drink (15 mins) - Not sure what the vuln scanner likes to drink? How about offering ALL THE DRINKS? 

Oh look, Metasploit too. - (5 mins) We even add some capability in Metasploit to exploit this capability right now.

Q/A (5 mins) - Self explanatory 

**== List of other conferences ==**

None

**== Why is this a good fit for Con ==**

This talk provides an high level analysis of one of the critical functions within the industry vulnerability scanning tools which have had a gaping security risk for quite some time. We're hoping that by inspecting our own (Infosec) tools and talking openly about their issues we can encourage the software development teams to provide easily implemented and user configurable controls around credential management. DerbyCon is a highly respected Con in the Infosec space, and the addition of publicly posting the conference talks encourages the knowledge to spread across the industry. 

**== Previous Experience ==**

Stable Talk at DerbyCon 2016
