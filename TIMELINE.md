# Historical Timeline for the "The Cistern Attack: Exploiting Absolute Trust" Project

## Initial finding
<2010 - While working as a unix administrator, I identified that there was an authentication from a vulnerability scanning tool using SSH to one of my servers. Edited OpenSSH and recompiled it so that the daemon would echo the U/P presented out to a log file. It worked and I verified that I now had another root level account for other servers that I did not manage.
### Thoughts at this time. I can steal ssh credential from a common vulnerability scanner on port 22. Interesting.

## Further research
2013 - Introduced Troy to this finding. Tested again with two vulnerability scanners and played with telnet service and found that this affected telnet
### Thoughts at this time. I can steal SSH and Telnet credentials from two security scanners. This is getting interesting, but I have other projects to focus on. Let me come back to this later but use it for "research gigs".

## Further research
2016 - Presented this issue to Tenable and Rapid7. Rapid7 said this is how it is designed to work and is not an issue. Tenable did not respond.
2017 - 2018 - Talked about the idea to Ben after going through some "hoops" at my employer. Also gained much more in depth knowledge on how vulnerability scanners worked using this process, figured out this impacted MUCH more than just a couple of vulnerability scanners. Identified that we could grab SMB also. Also identified that we could listen on any common port that scanners "check" and get the credentals for SSH and SNMP.
### Thoughts at this time. Oh damn. This is a bad day for asset discovery and vulnerability scanning tools. We better tell more people about this issue!

## Talk Submissions
2017 - Submit the talk idea to DerbyCon 7. Get Rejected. (Old title and presentation format - Social Engineering Vuln Scanners)
2017 - Submit the talk idea to ShmooCon 2018. Get Rejected. (Used the newer title of Cistern Attack)
2018 - Submit the talk idea to DerbyCon 8. Get rejected. (see the submission and rejection email - sad face)
### Thoughts at this time. WTF man! Why isn't this catching? Let me revisit it when I find someone that wants to hear about it.

## Get forced to move files and information from private repo to public one
I checked in at DerbyCon 9 (September 6th, 2019) and reviewed the program of talks for September 7th and immediately saw another researcher's talk on this exact subject matter that I had previously submitted for review. I immediately reached out to the other project members and we put up some of our smaller tools (the original goal of this was to have it remain small and usable with other existing tools) and the related project, finding, and talk submission information on Friday, September 6th, 2019. I attended their talk on Saturday morning and attempted to engage with them on this topic, but they were focused on the specific tool that they built for red teams to use this vulnerability in testing. You can find Jacob Griffith and Tim Right's tool and information at [SWARM FRAMEWORK](https://github.com/swarmframework/swarm). While I was excited that this information was getting out, I was a little sad that our almost identical talk submission wasn't accepted the year before.


