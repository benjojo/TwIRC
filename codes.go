package main

//shouts to http://stackoverflow.com/questions/8490084/trying-to-write-an-irc-client-but-struggling-to-find-a-good-resource-regarding-c
const (
	RplNone int = 0

	// Initial
	RplWelcome   int = 001 // :Welcome to the Internet Relay Network <nickname>
	RplYourHost  int = 002 // :Your host is <server>, running version <ver>
	RplCreated   int = 003 // :This server was created <datetime>
	RplMyInfo    int = 004 // <server> <ver> <usermode> <chanmode>
	RplMap       int = 005 // :map
	RplEndOfMap  int = 007 // :End of /MAP
	RplMotdStart int = 375 // :- server Message of the Day
	RplMotd      int = 372 // :- <info>
	RplMotdAlt   int = 377 // :- <info>                                                                        (some)
	RplMotdAlt2  int = 378 // :- <info>                                                                        (some)
	RplMotdEnd   int = 376 // :End of /MOTD command.
	RplUModeIs   int = 221 // <mode>

	// IsOn/UserHost
	RplUserHost int = 302 // :userhosts
	RplIsOn     int = 303 // :nicknames

	// Away
	RplAway    int = 301 // <nick> :away
	RplUnAway  int = 305 // :You are no longer marked as being away
	RplNowAway int = 306 // :You have been marked as being away

	// WHOIS/WHOWAS
	RplWhoisHelper   int = 310 // <nick> :looks very helpful                                                       DALNET
	RplWhoIsUser     int = 311 // <nick> <username> <address> * :<info>
	RplWhoIsServer   int = 312 // <nick> <server> :<info>
	RplWhoIsOperator int = 313 // <nick> :is an IRC Operator
	RplWhoIsIdle     int = 317 // <nick> <seconds> <signon> :<info>
	RplEndOfWhois    int = 318 // <request> :End of /WHOIS list.
	RplWhoIsChannels int = 319 // <nick> :<channels>
	RplWhoWasUser    int = 314 // <nick> <username> <address> * :<info>
	RplEndOfWhoWas   int = 369 // <request> :End of WHOWAS
	RplWhoReply      int = 352 // <channel> <username> <address> <server> <nick> <flags> :<hops> <info>
	RplEndOfWho      int = 315 // <request> :End of /WHO list.
	RplUserIPs       int = 307 // :userips                                                                         UNDERNET
	RplUserIP        int = 340 // <nick> :<nickname>=+<user>@<IP.address>                                          UNDERNET

	// List
	RplListStart  int = 321 // Channel :Users Name
	RplList       int = 322 // <channel> <users> :<topic>
	RplListEnd    int = 323 // :End of /LIST
	RplLinks      int = 364 // <server> <hub> :<hops> <info>
	RplEndOfLinks int = 365 // <mask> :End of /LINKS list.

	// Post-Channel Join
	RplUniqOpIs       int = 325
	RplChannelModeIs  int = 324 // <channel> <mode>
	RplChannelUrl     int = 328 // <channel> :url                                                                   DALNET
	RplChannelCreated int = 329 // <channel> <time>
	RplNoTopic        int = 331 // <channel> :No topic is set.
	RplTopic          int = 332 // <channel> :<topic>
	RplTopicSetBy     int = 333 // <channel> <nickname> <time>
	RplNamReply       int = 353 // = <channel> :<names>
	RplEndOfNames     int = 366 // <channel> :End of /NAMES list.

	// Invitational
	RplInviting  int = 341 // <nick> <channel>
	RplSummoning int = 342

	// Channel Lists
	RplInviteList      int = 346 // <channel> <invite> <nick> <time>                                                 IRCNET
	RplEndOfInviteList int = 357 // <channel> :End of Channel Invite List                                            IRCNET
	RplExceptList      int = 348 // <channel> <exception> <nick> <time>                                              IRCNET
	RplEndOfExceptList int = 349 // <channel> :End of Channel Exception List                                         IRCNET
	RplBanList         int = 367 // <channel> <ban> <nick> <time>
	RplEndOfBanList    int = 368 // <channel> :End of Channel Ban List

	// server/misc
	RplVersion      int = 351 // <version>.<debug> <server> :<info>
	RplInfo         int = 371 // :<info>
	RplEndOfInfo    int = 374 // :End of /INFO list.
	RplYoureOper    int = 381 // :You are now an IRC Operator
	RplRehashing    int = 382 // <file> :Rehashing
	RplYoureService int = 383
	RplTime         int = 391 // <server> :<time>
	RplUsersStart   int = 392
	RplUsers        int = 393
	RplEndOfUsers   int = 394
	RplNoUsers      int = 395
	RplServList     int = 234
	RplServListEnd  int = 235
	RplAdminMe      int = 256 // :Administrative info about server
	RplAdminLoc1    int = 257 // :<info>
	RplAdminLoc2    int = 258 // :<info>
	RplAdminEMail   int = 259 // :<info>
	RplTryAgain     int = 263 // :Server load is temporarily too heavy. Please wait a while and try again.

	// tracing
	RplTraceLink       int = 200
	RplTraceConnecting int = 201
	RplTraceHandshake  int = 202
	RplTraceUnknown    int = 203
	RplTraceOperator   int = 204
	RplTraceUser       int = 205
	RplTraceServer     int = 206
	RplTraceService    int = 207
	RplTraceNewType    int = 208
	RplTraceClass      int = 209
	RplTraceReconnect  int = 210
	RplTraceLog        int = 261
	RplTraceEnd        int = 262

	// stats
	RplStatsLinkInfo int = 211 // <connection> <sendq> <sentmsg> <sentbyte> <recdmsg> <recdbyte> :<open>
	RplStatsCommands int = 212 // <command> <uses> <bytes>
	RplStatsCLine    int = 213 // C <address> * <server> <port> <class>
	RplStatsNLine    int = 214 // N <address> * <server> <port> <class>
	RplStatsILine    int = 215 // I <ipmask> * <hostmask> <port> <class>
	RplStatsKLine    int = 216 // k <address> * <username> <details>
	RplStatsPLine    int = 217 // P <port> <??> <??>
	RplStatsQLine    int = 222 // <mask> :<comment>
	RplStatsELine    int = 223 // E <hostmask> * <username> <??> <??>
	RplStatsDLine    int = 224 // D <ipmask> * <username> <??> <??>
	RplStatsLLine    int = 241 // L <address> * <server> <??> <??>
	RplStatsuLine    int = 242 // :Server Up <num> days, <time>
	RplStatsoLine    int = 243 // o <mask> <password> <user> <??> <class>
	RplStatsHLine    int = 244 // H <address> * <server> <??> <??>
	RplStatsGLine    int = 247 // G <address> <timestamp> :<reason>
	RplStatsULine    int = 248 // U <host> * <??> <??> <??>
	RplStatsZLine    int = 249 // :info
	RplStatsYLine    int = 218 // Y <class> <ping> <freq> <maxconnect> <sendq>
	RplEndOfStats    int = 219 // <char> :End of /STATS report
	RplStatsUptime   int = 242

	// GLINE
	RplGLineList      int = 280 // <address> <timestamp> <reason>                                                   UNDERNET
	RplEndOfGLineList int = 281 // :End of G-line List                                                              UNDERNET

	// Silence
	RplSilenceList      int = 271 // <nick> <mask>                                                                    UNDERNET/DALNET
	RplEndOfSilenceList int = 272 // <nick> :End of Silence List                                                      UNDERNET/DALNET

	// LUser
	RplLUserClient     int = 251 // :There are <user> users and <invis> invisible on <serv> servers
	RplLUserOp         int = 252 // <num> :operator(s) online
	RplLUserUnknown    int = 253 // <num> :unknown connection(s)
	RplLUserChannels   int = 254 // <num> :channels formed
	RplLUserMe         int = 255 // :I have <user> clients and <serv> servers
	RplLUserLocalUser  int = 265 // :Current local users: <curr> Max: <max>
	RplLUserGlobalUser int = 266 // :Current global users: <curr> Max: <max>

	// Errors
	ErrNoSuchNick        int = 401 // <nickname> :No such nick
	ErrNoSuchServer      int = 402 // <server> :No such server
	ErrNoSuchChannel     int = 403 // <channel> :No such channel
	ErrCannotSendToChan  int = 404 // <channel> :Cannot send to channel
	ErrTooManyChannels   int = 405 // <channel> :You have joined too many channels
	ErrWasNoSuchNick     int = 406 // <nickname> :There was no such nickname
	ErrTooManyTargets    int = 407 // <target> :Duplicate recipients. No message delivered
	ErrNoColors          int = 408 // <nickname> #<channel> :You cannot use colors on this channel. Not sent: <text>   DALNET
	ErrNoOrigin          int = 409 // :No origin specified
	ErrNoRecipient       int = 411 // :No recipient given (<command>)
	ErrNoTextToSend      int = 412 // :No text to send
	ErrNoTopLevel        int = 413 // <mask> :No toplevel domain specified
	ErrWildTopLevel      int = 414 // <mask> :Wildcard in toplevel Domain
	ErrBadMask           int = 415
	ErrTooMuchInfo       int = 416 // <command> :Too many lines in the output, restrict your query                     UNDERNET
	ErrUnknownCommand    int = 421 // <command> :Unknown command
	ErrNoMotd            int = 422 // :MOTD File is missing
	ErrNoAdminInfo       int = 423 // <server> :No administrative info available
	ErrFileError         int = 424
	ErrNoNicknameGiven   int = 431 // :No nickname given
	ErrErroneusNickname  int = 432 // <nickname> :Erroneus Nickname
	ErrNickNameInUse     int = 433 // <nickname> :Nickname is already in use.
	ErrNickCollision     int = 436 // <nickname> :Nickname collision KILL
	ErrUnAvailResource   int = 437 // <channel> :Cannot change nickname while banned on channel
	ErrNickTooFast       int = 438 // <nick> :Nick change too fast. Please wait <sec> seconds.                         (most)
	ErrTargetTooFast     int = 439 // <target> :Target change too fast. Please wait <sec> seconds.                     DALNET/UNDERNET
	ErrUserNotInChannel  int = 441 // <nickname> <channel> :They aren't on that channel
	ErrNotOnChannel      int = 442 // <channel> :You're not on that channel
	ErrUserOnChannel     int = 443 // <nickname> <channel> :is already on channel
	ErrNoLogin           int = 444
	ErrSummonDisabled    int = 445 // :SUMMON has been disabled
	ErrUsersDisabled     int = 446 // :USERS has been disabled
	ErrNotRegistered     int = 451 // <command> :Register first.
	ErrNeedMoreParams    int = 461 // <command> :Not enough parameters
	ErrAlreadyRegistered int = 462 // :You may not reregister
	ErrNoPermForHost     int = 463
	ErrPasswdMistmatch   int = 464
	ErrYoureBannedCreep  int = 465
	ErrYouWillBeBanned   int = 466
	ErrKeySet            int = 467 // <channel> :Channel key already set
	ErrServerCanChange   int = 468 // <channel> :Only servers can change that mode                                     DALNET
	ErrChannelIsFull     int = 471 // <channel> :Cannot join channel (+l)
	ErrUnknownMode       int = 472 // <char> :is unknown mode char to me
	ErrInviteOnlyChan    int = 473 // <channel> :Cannot join channel (+i)
	ErrBannedFromChan    int = 474 // <channel> :Cannot join channel (+b)
	ErrBadChannelKey     int = 475 // <channel> :Cannot join channel (+k)
	ErrBadChanMask       int = 476
	ErrNickNotRegistered int = 477 // <channel> :You need a registered nick to join that channel.                      DALNET
	ErrBanListFull       int = 478 // <channel> <ban> :Channel ban/ignore list is full
	ErrNoPrivileges      int = 481 // :Permission Denied- You're not an IRC operator
	ErrChanOPrivsNeeded  int = 482 // <channel> :You're not channel operator
	ErrCantKillServer    int = 483 // :You cant kill a server!
	ErrRestricted        int = 484 // <nick> <channel> :Cannot kill, kick or deop channel service                      UNDERNET
	ErrUniqOPrivsNeeded  int = 485 // <channel> :Cannot join channel (reason)
	ErrNoOperHost        int = 491 // :No O-lines for your host
	ErrUModeUnknownFlag  int = 501 // :Unknown MODE flag
	ErrUsersDontMatch    int = 502 // :Cant change mode for other users
	ErrSilenceListFull   int = 511 // <mask> :Your silence list is full                                                UNDERNET/DALNET
)
