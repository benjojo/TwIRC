TwIRC
=====

#The IRC friendly twitter client

### Setup
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

### Making a tweet

A tweet can be done by simply posting in the `#twitterstream` room. Once posted you will see your own tweet appear again as a confirmation that it has been posted:


![](http://i.imgur.com/2R8QsMN.png)


### Removing the last tweet

Can be done by running `/undo`.

When you do this you will be PM'd by `SYS` to confirm that it has been undone.

The undo will be confimed as well by a `SYS` message being posted in the main room about you removing the tweet.

![](http://i.imgur.com/qGbKgxR.png)

### Replying to someone

#### If someone has tweeted to your directly

If someone has tweeted to you directly then you will get a message in both the `#twitterstream` room and also as a PM from their user.

Replying to them is as simple as responding back:

![](http://i.imgur.com/K6BpVw3.png)

When replying to a user it will automatically use the correct replyID and also add the "@user" part:

![](http://i.imgur.com/8ZSfJ2F.png)

#### If someone has tweeted and you want to respond

You can use `/all` and then PM them directly. The same thing will happen but the replyID will be fixed to whatever the last tweet of that person was. To change it back to the normal mode, use `/mention`

### Unfollowing someone

If you kick someone out of the `#twitterstream` room, they will be unfollowed:

![](http://i.imgur.com/5kkLdRE.png)
![](http://i.imgur.com/hPpAmvm.png)

