# Unified IMAP Proxy
I started writing this tool to help provide a single sign on experience for members of my family to view their email via NextCloud.
We all use different providers and configuring each of these accounts manually via NextCloud could be cumbersome for some of them.
I personally am using ProtonMail which doesn't have IMAP capabilities unless you use the bridge which I cannot seem to get working
in any reasonably sane way.

Ultimately the end-use of this tool will look something like this:

* User signs into NextCloud
* NextCloud passes on single sign on credentials to IMAP proxy
* IMAP proxy then uses those credentials to talk to something like BitWarden and pull down the users credentials for their email service.
* IMAP proxy acts as a client to that email service and forwards on.
