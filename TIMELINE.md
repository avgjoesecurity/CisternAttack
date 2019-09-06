# Historical Timeline for the "The Cistern Attack: Exploiting Absolute Trust" Project

## Initial finding
2010 - While working as a unix administrator, I identified that there was an authentication from a vulnerability scanning tool using SSH to one of my servers. Edited OpenSSH and recompiled it so that the daemon would echo the U/P presented out to a log file. It worked and I verified that I now had another root level account for other servers that I did not manage.
### Thoughts at this time. I can steal ssh credential from a common vulnerability scanner on port 22. Interesting.

## Further research
2013 - Introduced Troy to this finding. Tested again with two vulnerability scanners and played with telnet service and found that this affected telnet
### Thoughts at this time. I can steal SSH and Telnet credentials from two security scanners. This is getting interesting, but I have other projects to focus on. Let me come back to this later but use it for "research gigs".

## Further research
2017 - Talked about the idea to Ben after going through some "hoops" at my employer. Also gained much more in depth knowledge on how vulnerability scanners "worked" in this process, figured out this impact MUCH more than just a couple of vulnerability scanners.
### Thoughts at this time. Oh shit. This is a bad day for asset discovery and vulnerability scanning tools. I better tell more people about this issue!

## Talk Submissions
2018 - Submit the talk idea to DerbyCon 8. Get rejected. (see the submission and rejection email - sad face)
2018 - Submit the talk idea to ShmooCon. Get rejected. (see the submission, lost the rejection email - sad face)
### Thoughts at this time. WTF man! Why isn't this catching? Let me revisit it when I find someone that wants to hear about it.
