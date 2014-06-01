package main

//shouts to http://stackoverflow.com/questions/8490084/trying-to-write-an-irc-client-but-struggling-to-find-a-good-resource-regarding-c
const (
	RplNone string = "0"

	// Initial
	RplWelcome   string = "001" // :Welcome to the Internet Relay Network <nickname>
	RplYourHost  string = "002" // :Your host is <server>, running version <ver>
	RplCreated   string = "003" // :This server was created <datetime>
	RplMyInfo    string = "004" // <server> <ver> <usermode> <chanmode>
	RplMap       string = "005" // :map
	RplEndOfMap  string = "007" // :End of /MAP
	RplMotdStart string = "375" // :- server Message of the Day
	RplMotd      string = "372" // :- <info>
	RplMotdAlt   string = "377" // :- <info>                                                                        (some)
	RplMotdAlt2  string = "378" // :- <info>                                                                        (some)
	RplMotdEnd   string = "376" // :End of /MOTD command.
	RplUModeIs   string = "221" // <mode>

	// IsOn/UserHost
	RplUserHost string = "302" // :userhosts
	RplIsOn     string = "303" // :nicknames

	// Away
	RplAway    string = "301" // <nick> :away
	RplUnAway  string = "305" // :You are no longer marked as being away
	RplNowAway string = "306" // :You have been marked as being away

	// WHOIS/WHOWAS
	RplWhoisHelper   string = "310" // <nick> :looks very helpful                                                       DALNET
	RplWhoIsUser     string = "311" // <nick> <username> <address> * :<info>
	RplWhoIsServer   string = "312" // <nick> <server> :<info>
	RplWhoIsOperator string = "313" // <nick> :is an IRC Operator
	RplWhoIsIdle     string = "317" // <nick> <seconds> <signon> :<info>
	RplEndOfWhois    string = "318" // <request> :End of /WHOIS list.
	RplWhoIsChannels string = "319" // <nick> :<channels>
	RplWhoWasUser    string = "314" // <nick> <username> <address> * :<info>
	RplEndOfWhoWas   string = "369" // <request> :End of WHOWAS
	RplWhoReply      string = "352" // <channel> <username> <address> <server> <nick> <flags> :<hops> <info>
	RplEndOfWho      string = "315" // <request> :End of /WHO list.
	RplUserIPs       string = "307" // :userips                                                                         UNDERNET
	RplUserIP        string = "340" // <nick> :<nickname>=+<user>@<IP.address>                                          UNDERNET

	// List
	RplListStart  string = "321" // Channel :Users Name
	RplList       string = "322" // <channel> <users> :<topic>
	RplListEnd    string = "323" // :End of /LIST
	RplLinks      string = "364" // <server> <hub> :<hops> <info>
	RplEndOfLinks string = "365" // <mask> :End of /LINKS list.

	// Post-Channel Join
	RplUniqOpIs       string = "325"
	RplChannelModeIs  string = "324" // <channel> <mode>
	RplChannelUrl     string = "328" // <channel> :url                                                                   DALNET
	RplChannelCreated string = "329" // <channel> <time>
	RplNoTopic        string = "331" // <channel> :No topic is set.
	RplTopic          string = "332" // <channel> :<topic>
	RplTopicSetBy     string = "333" // <channel> <nickname> <time>
	RplNamReply       string = "353" // = <channel> :<names>
	RplEndOfNames     string = "366" // <channel> :End of /NAMES list.

	// Invitational
	RplInviting  string = "341" // <nick> <channel>
	RplSummoning string = "342"

	// Channel Lists
	RplInviteList      string = "346" // <channel> <invite> <nick> <time>                                                 IRCNET
	RplEndOfInviteList string = "357" // <channel> :End of Channel Invite List                                            IRCNET
	RplExceptList      string = "348" // <channel> <exception> <nick> <time>                                              IRCNET
	RplEndOfExceptList string = "349" // <channel> :End of Channel Exception List                                         IRCNET
	RplBanList         string = "367" // <channel> <ban> <nick> <time>
	RplEndOfBanList    string = "368" // <channel> :End of Channel Ban List

	// server/misc
	RplVersion      string = "351" // <version>.<debug> <server> :<info>
	RplInfo         string = "371" // :<info>
	RplEndOfInfo    string = "374" // :End of /INFO list.
	RplYoureOper    string = "381" // :You are now an IRC Operator
	RplRehashing    string = "382" // <file> :Rehashing
	RplYoureService string = "383"
	RplTime         string = "391" // <server> :<time>
	RplUsersStart   string = "392"
	RplUsers        string = "393"
	RplEndOfUsers   string = "394"
	RplNoUsers      string = "395"
	RplServList     string = "234"
	RplServListEnd  string = "235"
	RplAdminMe      string = "256" // :Administrative info about server
	RplAdminLoc1    string = "257" // :<info>
	RplAdminLoc2    string = "258" // :<info>
	RplAdminEMail   string = "259" // :<info>
	RplTryAgain     string = "263" // :Server load is temporarily too heavy. Please wait a while and try again.

	// tracing
	RplTraceLink       string = "200"
	RplTraceConnecting string = "201"
	RplTraceHandshake  string = "202"
	RplTraceUnknown    string = "203"
	RplTraceOperator   string = "204"
	RplTraceUser       string = "205"
	RplTraceServer     string = "206"
	RplTraceService    string = "207"
	RplTraceNewType    string = "208"
	RplTraceClass      string = "209"
	RplTraceReconnect  string = "210"
	RplTraceLog        string = "261"
	RplTraceEnd        string = "262"

	// stats
	RplStatsLinkInfo string = "211" // <connection> <sendq> <sentmsg> <sentbyte> <recdmsg> <recdbyte> :<open>
	RplStatsCommands string = "212" // <command> <uses> <bytes>
	RplStatsCLine    string = "213" // C <address> * <server> <port> <class>
	RplStatsNLine    string = "214" // N <address> * <server> <port> <class>
	RplStatsILine    string = "215" // I <ipmask> * <hostmask> <port> <class>
	RplStatsKLine    string = "216" // k <address> * <username> <details>
	RplStatsPLine    string = "217" // P <port> <??> <??>
	RplStatsQLine    string = "222" // <mask> :<comment>
	RplStatsELine    string = "223" // E <hostmask> * <username> <??> <??>
	RplStatsDLine    string = "224" // D <ipmask> * <username> <??> <??>
	RplStatsLLine    string = "241" // L <address> * <server> <??> <??>
	RplStatsuLine    string = "242" // :Server Up <num> days, <time>
	RplStatsoLine    string = "243" // o <mask> <password> <user> <??> <class>
	RplStatsHLine    string = "244" // H <address> * <server> <??> <??>
	RplStatsGLine    string = "247" // G <address> <timestamp> :<reason>
	RplStatsULine    string = "248" // U <host> * <??> <??> <??>
	RplStatsZLine    string = "249" // :info
	RplStatsYLine    string = "218" // Y <class> <ping> <freq> <maxconnect> <sendq>
	RplEndOfStats    string = "219" // <char> :End of /STATS report
	RplStatsUptime   string = "242"

	// GLINE
	RplGLineList      string = "280" // <address> <timestamp> <reason>                                                   UNDERNET
	RplEndOfGLineList string = "281" // :End of G-line List                                                              UNDERNET

	// Silence
	RplSilenceList      string = "271" // <nick> <mask>                                                                    UNDERNET/DALNET
	RplEndOfSilenceList string = "272" // <nick> :End of Silence List                                                      UNDERNET/DALNET

	// LUser
	RplLUserClient     string = "251" // :There are <user> users and <invis> invisible on <serv> servers
	RplLUserOp         string = "252" // <num> :operator(s) online
	RplLUserUnknown    string = "253" // <num> :unknown connection(s)
	RplLUserChannels   string = "254" // <num> :channels formed
	RplLUserMe         string = "255" // :I have <user> clients and <serv> servers
	RplLUserLocalUser  string = "265" // :Current local users: <curr> Max: <max>
	RplLUserGlobalUser string = "266" // :Current global users: <curr> Max: <max>

	// Errors
	ErrNoSuchNick        string = "401" // <nickname> :No such nick
	ErrNoSuchServer      string = "402" // <server> :No such server
	ErrNoSuchChannel     string = "403" // <channel> :No such channel
	ErrCannotSendToChan  string = "404" // <channel> :Cannot send to channel
	ErrTooManyChannels   string = "405" // <channel> :You have joined too many channels
	ErrWasNoSuchNick     string = "406" // <nickname> :There was no such nickname
	ErrTooManyTargets    string = "407" // <target> :Duplicate recipients. No message delivered
	ErrNoColors          string = "408" // <nickname> #<channel> :You cannot use colors on this channel. Not sent: <text>   DALNET
	ErrNoOrigin          string = "409" // :No origin specified
	ErrNoRecipient       string = "411" // :No recipient given (<command>)
	ErrNoTextToSend      string = "412" // :No text to send
	ErrNoTopLevel        string = "413" // <mask> :No toplevel domain specified
	ErrWildTopLevel      string = "414" // <mask> :Wildcard in toplevel Domain
	ErrBadMask           string = "415"
	ErrTooMuchInfo       string = "416" // <command> :Too many lines in the output, restrict your query                     UNDERNET
	ErrUnknownCommand    string = "421" // <command> :Unknown command
	ErrNoMotd            string = "422" // :MOTD File is missing
	ErrNoAdminInfo       string = "423" // <server> :No administrative info available
	ErrFileError         string = "424"
	ErrNoNicknameGiven   string = "431" // :No nickname given
	ErrErroneusNickname  string = "432" // <nickname> :Erroneus Nickname
	ErrNickNameInUse     string = "433" // <nickname> :Nickname is already in use.
	ErrNickCollision     string = "436" // <nickname> :Nickname collision KILL
	ErrUnAvailResource   string = "437" // <channel> :Cannot change nickname while banned on channel
	ErrNickTooFast       string = "438" // <nick> :Nick change too fast. Please wait <sec> seconds.                         (most)
	ErrTargetTooFast     string = "439" // <target> :Target change too fast. Please wait <sec> seconds.                     DALNET/UNDERNET
	ErrUserNotInChannel  string = "441" // <nickname> <channel> :They aren't on that channel
	ErrNotOnChannel      string = "442" // <channel> :You're not on that channel
	ErrUserOnChannel     string = "443" // <nickname> <channel> :is already on channel
	ErrNoLogin           string = "444"
	ErrSummonDisabled    string = "445" // :SUMMON has been disabled
	ErrUsersDisabled     string = "446" // :USERS has been disabled
	ErrNotRegistered     string = "451" // <command> :Register first.
	ErrNeedMoreParams    string = "461" // <command> :Not enough parameters
	ErrAlreadyRegistered string = "462" // :You may not reregister
	ErrNoPermForHost     string = "463"
	ErrPasswdMistmatch   string = "464"
	ErrYoureBannedCreep  string = "465"
	ErrYouWillBeBanned   string = "466"
	ErrKeySet            string = "467" // <channel> :Channel key already set
	ErrServerCanChange   string = "468" // <channel> :Only servers can change that mode                                     DALNET
	ErrChannelIsFull     string = "471" // <channel> :Cannot join channel (+l)
	ErrUnknownMode       string = "472" // <char> :is unknown mode char to me
	ErrInviteOnlyChan    string = "473" // <channel> :Cannot join channel (+i)
	ErrBannedFromChan    string = "474" // <channel> :Cannot join channel (+b)
	ErrBadChannelKey     string = "475" // <channel> :Cannot join channel (+k)
	ErrBadChanMask       string = "476"
	ErrNickNotRegistered string = "477" // <channel> :You need a registered nick to join that channel.                      DALNET
	ErrBanListFull       string = "478" // <channel> <ban> :Channel ban/ignore list is full
	ErrNoPrivileges      string = "481" // :Permission Denied- You're not an IRC operator
	ErrChanOPrivsNeeded  string = "482" // <channel> :You're not channel operator
	ErrCantKillServer    string = "483" // :You cant kill a server!
	ErrRestricted        string = "484" // <nick> <channel> :Cannot kill, kick or deop channel service                      UNDERNET
	ErrUniqOPrivsNeeded  string = "485" // <channel> :Cannot join channel (reason)
	ErrNoOperHost        string = "491" // :No O-lines for your host
	ErrUModeUnknownFlag  string = "501" // :Unknown MODE flag
	ErrUsersDontMatch    string = "502" // :Cant change mode for other users
	ErrSilenceListFull   string = "511" // <mask> :Your silence list is full                                                UNDERNET/DALNET
)
