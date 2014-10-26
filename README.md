TwIRC
=====

###The IRC friendly twitter client

![](http://i.imgur.com/m75cnMK.png)

TwIRC is a service that you connect to though a IRC client to use twitter on.

It requires a Twitter Application API key pair, Those are easy to obtain though.

Once you have the keypair you will need to put that in a `twitterauth.cfg` file, There is a sample in this repo of what it is expecting.

After that you can start the server up, point your IRC client to it and connect.

If you connect with out auth, then you will be PM'd by a "SYS" user giving you instructions on how to use the TwIRC server.

![](http://i.imgur.com/BqsXwKJ.png)

After you have followed those instructions, a JSON blob will be PM'd to you. You should then set that as the "server password" in your client:

![](http://i.imgur.com/mmYY0nR.png)

After you have done all of that, connect and then join the `#twitterstream` room, The people in that room should be all of the people you are following.