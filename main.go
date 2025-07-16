package main

import (
	"encoding/base64"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"runtime"
	"strings"
)

// overwrite this at build time ;
// -ldflags="-X 'main.Version=someversion'"
var Version = "latest"

func PrintVersionAndQuit() {
	fmt.Println(Version)
	os.Exit(0)
}

func Debug() {
	fmt.Printf("%s %s %s %d\n", runtime.Version(), runtime.GOOS, runtime.GOARCH, runtime.NumCPU())
	os.Exit(0)
}

func getPhrase() string {
	// pick a phrase from the list

	phrases := make([]string, 0)
	phrases = append(phrases,
		`Coming in on a bad note. Don't know ya and don't know if I care to given who ya sight as a source. Name's Saley, the GodKing. I've been accused of being a bit of a peacekeeper and ring leader of our little rag tag group. Ninja is a trusted second, and Roman is my official go to boy on vibes. Just as a fair warning, don't bring that name up again. You won't like what you get. That being said, strike 1, and welcome to the group.`,

		`Furious about an incident at work yesterday.

everyone gets out, group goes to the restroom

i am at the urinal trying to piss with everybody around talking, couple of farts come out

they laugh, im embarrassed but whatever, even more farts come, the kind that are quick fast pops

this douche bag who makes fun of me every chance he gets is like OH OH IF YOURE GIVING AWAY FREE GAS FILL ME UP FILL ME UP gets on his fucking knees and gets right behind me and opens his mouth going aaaahhh like at a dentist office and fake begs me to fart on him by saying he was running on E and shit

people were laughing, i just fucking finished and went out to my car to smoke

it was legit humiliating.`,

		`I'd just like to interject for a moment. What you're referring to as Linux, is in fact, GNU/Linux, or as I've recently taken to calling it, GNU plus Linux. Linux is not an operating system unto itself, but rather another free component of a fully functioning GNU system made useful by the GNU corelibs, shell utilities and vital system components comprising a full OS as defined by POSIX.

Many computer users run a modified version of the GNU system every day, without realizing it. Through a peculiar turn of events, the version of GNU which is widely used today is often called "Linux", and many of its users are not aware that it is basically the GNU system, developed by the GNU Project.

There really is a Linux, and these people are using it, but it is just a part of the system they use. Linux is the kernel: the program in the system that allocates the machine's resources to the other programs that you run. The kernel is an essential part of an operating system, but useless by itself; it can only function in the context of a complete operating system. Linux is normally used in combination with the GNU operating system: the whole system is basically GNU with Linux added, or GNU/Linux. All the so-called "Linux" distributions are really distributions of GNU/Linux.
`,
		`i am a heron. i ahev a long neck and i pick fish out of the water w/ my beak. if you dont repost this comment on 10 other pages i will fly into your kitchen tonight and make a mess of your pots and pans`,
		`I was shooting heroin and reading "The Fountainhead" in the front seat of my privately owned police cruiser when a call came in. I put a quarter in the radio to activate it. It was the chief.

"Bad news, detective. We got a situation on our hands."

"What? Is the mayor trying to ban trans fats again?"

"Worse. Somebody just stole four hundred and forty-seven million dollars' worth of bitcoins."

The heroin needle practically fell out of my arm. "What kind of monster would do something like that? Bitcoins are the ultimate currency: virtual, anonymous, stateless. They represent true economic freedom, not subject to arbitrary manipulation by any government. Do we have any leads?"

"Not yet. But mark my words: we're going to figure out who did this and we're going to take them down … provided someone pays us a fair market rate to do so."

"Easy, chief," I said. "Any rate the market offers is, by definition, fair."

He laughed. "That's why you're the best I got, Lisowski. Now you get out there and find those bitcoins."

"Don't worry," I said. "I'm on it."

I put a quarter in the siren. Ten minutes later, I was on the scene. It was a normal office building, strangled on all sides by public sidewalks. I hopped over them and went inside.

"Home Depot™ Presents the Police!®" I said, flashing my badge and my gun and a small picture of Ron Paul. "Nobody move unless you want to!" They didn't.

"Now, which one of you punks is going to pay me to investigate this crime?" No one spoke up.

"Come on," I said. "Don't you all understand that the protection of private property is the foundation of all personal liberty?"

It didn't seem like they did.

"Seriously, guys. Without a strong economic motivator, I'm just going to stand here and not solve this case. Cash is fine, but I prefer being paid in gold bullion or autographed Penn Jillette posters."
Nothing. These people were stonewalling me. It almost seemed like they didn't care that a fortune in computer money invented to buy drugs was missing.

I figured I could wait them out. I lit several cigarettes indoors. A pregnant lady coughed, and I told her that secondhand smoke is a myth. Just then, a man in glasses made a break for it.

"Subway™ Eat Fresh and Freeze, Scumbag!®" I yelled.

Too late. He was already out the front door. I went after him.

"Stop right there!" I yelled as I ran. He was faster than me because I always try to avoid stepping on public sidewalks. Our country needs a private-sidewalk voucher system, but, thanks to the incestuous interplay between our corrupt federal government and the public-sidewalk lobby, it will never happen.

I was losing him. "Listen, I'll pay you to stop!" I yelled. "What would you consider an appropriate price point for stopping? I'll offer you a thirteenth of an ounce of gold and a gently worn 'Bob Barr '08' extra-large long-sleeved men's T-shirt!"

He turned. In his hand was a revolver that the Constitution said he had every right to own. He fired at me and missed. I pulled my own gun, put a quarter in it, and fired back. The bullet lodged in a U.S.P.S. mailbox less than a foot from his head. I shot the mailbox again, on purpose.

"All right, all right!" the man yelled, throwing down his weapon. "I give up, cop! I confess: I took the bitcoins."

"Why'd you do it?" I asked, as I slapped a pair of Oikos™ Greek Yogurt Presents Handcuffs® on the guy.

"Because I was afraid."

"Afraid?"

"Afraid of an economic future free from the pernicious meddling of central bankers," he said. "I'm a central banker."

I wanted to coldcock the guy. Years ago, a central banker killed my partner. Instead, I shook my head.
"Let this be a message to all your central-banker friends out on the street," I said. "No matter how many bitcoins you steal, you'll never take away the dream of an open society based on the principles of personal and economic freedom."

He nodded, because he knew I was right. Then he swiped his credit card to pay me for arresting him.`,
		`Greetings.👵🏿😴👙⛄️ My🐕 name is🈚️ Jebuiz Y'har. If my🐯 calculations🎸 are👮🏿 correct,🚁 you should🚦 be♈️💅💓🐈 receiving🎅🏾 this📱🔍 transmission◽️ in💑 the year💆🏼 2018 AD.👨‍👨‍👦🚶🏾 It amuses👯✨ me🐐🔠📎 that you💗👑💇🏼💈 used to🛅 calculate🚣🏻🈺⏰ your dates👒♥️ in👨🏿 relation🔘 to🕝 the life of🚷👌🏿 an ancient🎈 man.🛄 You🈺 see,📄🐤 we💼🍅⌚️ have🍙 a😑 slightly💲 different timescale.🐱💋💁🏻🔶 But😪 to🙉 make🔬 things🚳 simple,💁🏼 I am🎳 writing🏢☀️🐯📱 from📱 the🕙 year⚾️ 49,170🏇🏿💑🚣🏾👃🏿 AD.🏭🐬🏊🏽
`,
		`Greetings. My name is Jebuiz y'har. If my calculations are correct, you should be receiving this transmission in the year 2013 AD. It amuses me that you used to calculate your dates in relation to the life of an ancient man. You see, we have a slightly different timescale. But to make things simple, I am writing from the year 49,170 AD.
`,
		`Do a barrel roll!`,
		`Mechanic: Somebody set up us the bomb.

Operator: Main screen turn on.

CATS: All your base are belong to us.

CATS: You have no chance to survive make your time.

Captain: Move 'ZIG'.

Captain: For great justice.
`,
		`Here I was minding my own business fapping, headphones on... when I think I hear a noise. I take the headphones off, turn the monitor off, and already have my clothes off so I don't know what to do. I could try putting clothes back on, but she might open my door while I'm doing that and I'd get caught for sure.

I hear the huge beast coming, my fucking mom; fat, insane, bitchy, deep hoarse voice. I don't know what the fuck to do so I hope to god that she doesn't come in.

Then, SHE COMES IN. I fucking kneel down and peek from behind the bed, and just stare in total awkwardness. She says "WUT U DOIN STILL UP." I just stare not knowing what the fuck to do or say. "U NAKED??" I hesitate and say "no." "WAT U DOIN" I don't know what the fuck to say. She leaves the room and bitches about everything she can possibly imagine, with some "FUCKIN PERVERT"s and "DON NO WUT U DOIN, FUCKIN PLAYING WITCHA SELF"s thrown in. I quickly put pants on and walk out the room to the kitchen acting like I'm getting something, hoping to subtly "convince" her that I wasn't naked and that I was simply not wearing a shirt. I go back into the room, ignoring her constant stream of bitching. She apparently wasn't convinced since I still heard a few "FUCKIN WEIRDO"s.

She then sometimes in (while I'm typing this amusing topic), saying "U ON COMPUTA? WAT U DOIN UP? WAS U MASTABATIN!!!!" I hesitate, and confusedly say "I just woke up..." She continues bitching and goes to her room, and she's still bitching a stream of bitchiness to this moment.

Now I'm totally out of the fucking mood so I just spent almost an hour working at it for NOTHING, since I never even got to finish up. Fucking ugly fat bitch, why'd it have to wake up.

If I would have known I'd NOT GET TO FINISH, I would have went to sleep nearly 2 fucking hours ago.
`,
		`I can guarantee you are not actually married and probably live alone like a schizo in your apartment. You have delusions of grandeur, and probably aren't even 30. All you are is a Trump-loving loser who wastes all your time playing League and bitching and whining about Biden in politics channel. How embarrassing is it being a grown man spending all your free time in discord that you could be using spending time with your imaginary wife and kids? I hope you reevaluate your life and start to focus on the things in life that actually matter, like reading a good book and having a nice laugh with REAL LIFE friends, not just your discord virtual friends. Rather than whine like a little bitch about political issues that you don't have the balls to do anything about. Good luck ✌️
`,
		`ive been trying to get on that site for so long its not even possible. i have contacted people i suspect to be members, used advanced data-mining techniques, and even corresponded via snail-mail with a moderator. It's more or less like "The Matrix" in that you cannot understand it untill you are selected to. Your site is merely spreading dis-info and throwing people off the trail. This site is an endless web of classified information concerning every subject imaginable. The things that users have access to is literally impossible to comprehend. You would have more luck becoming the president of the entire earth than being able to become a member. Its not possible ive tried too many times for that to be a fact that you could carry out. The web of truths and lies surrounding luelinks has circulated beyond what i have even found, and ive found basically all their is to find about this fact. Luelinks has not ever escaped its own enigma, and neither can you, especially someone like you trying to become a moderating admin. Ive talked to LlamaMan about getting an invite and he hasnt ever responded to me. Ive dug up some info telling me that some guy called LargeCow or something like that has moderating powers and he can grant anyone access, but I dont think youll be able to talk to him about it. It was tough enough for me to try to do it. Trust me man youre not gonna be able to get on Luelinks. i'd be willing to give away my wii and all of my nintendo games for an invite. I want to be able to watch all the movies off the site, and be able to find all the newest youtube videos that ive heard they make. Apparently the site is even funnier than collegehumor LOL. I odnt know man im gonna keep trying but you better just stop right now since youre not gonna be able to get in. youll never be an admin. Stop asking. A funnier video that they made was the one about barbecue sauce in the fridge and that guy talking about getting girls with it. The variety that he had in his fridge was hiLOLrious. Lol. Only a internet webpage like luelinks could make a video like that. I digress, I suppose, but my main point is this, my friend: LUElinks is the ultimate enigma of the internet. They have less than 100 members, all of whom treat each other as brothers. Each one of them commands god-like levels of processing power and RAM and interntet speed and whatnot which they get illegally through contacts on the site and the US government. If the government was to know about the government and computer contacts they compiled they would be with out a dowbt arrested on site. It is even better than gaia online and the other fast paced exciting message boards. Access to computers would be banned and no member would be able to use the telephone because of the hacking ability that is available. This is a fact that was proven by a bust of a a luelinks narc. But donbt get the wrong idea about ME. I would never do that of course. I would be a good contribLUEter (lol that's what they say) of links to LUElinks. Think of it this way: There is only one tiger left in the world (of the internet) and its LUElinks and its wild and beautiful and invisible to any non member, and to even get a general idea of what this majestic animal spirit tyger looks like I have had to dedicate hundreds of man hours and over sixty dollars (US) to informants from within the site and former members (and moderators) who have been affiliated with the site but kicked out because they tried to help me to get inside. //T/ his ibasically just stop trying to get in. if a luelinks pro like me couldn't get in, you don't have a single chance. Sorry man but you just cant do it but I will and then ill send you an invite if I want but I may s honestly the thing I think about most getting into this site of all sites and all possible data and info you could ever want to need or even imagine even if you were dreaming and on drugs from the future. So be too busy moderating once I get the entrance.
`,
		`Seriously, mods are only chosen a few times a year, and the minimum requirements for applying are very steep. Even so, hundreds apply, but only a few are selected. If you seriously want to be a mod, you'll likely (but not always) need to have a history on GameFAQs that makes you recognizable to the majority of the user-base for several years to even be considered. Pestering the current moderators or admins essentially guarantees that you won't be selected.
`,
		`Ok, I'm sitting on the toilet pooping, thinking about hot girls, reading this months EGM, and smelling my own gas all at once but at the same time trying not to intertwine each of these things in my mind because one can really ruin the other. Anywho, I shot out an enormous crap log and at that moment I realized that I didn't wipe myself the last time i pooped. The old crap melded with my ass pubes and dried to form an ass-pube-net, if you will. So today when I shot out that big log of human goodness(poop) It got trapped and entangled in my ass-pube-net. It was pretty messy and it hasn't dried yet. I got a fork out of the kitchen and tried to break the net, but no dice. Any ideas?

I tried the shower thing and the hot water didn't melt my ass-pube-net. And BTW, I wiped the poop off the fork with a pair of my uncle's dirty underpants which were sitting on the kitchen table before putting it back.

I had my girlfriend come over to help me with my ass-pube net. She tried to remove it with a comb, a pair of her aunt's nail clippers, and an icecube, but it didn't get rid of my asspube net and only ended up getting her dirty and smelly, so she gave up and went home.

My ass-pube-net problem was resolved by the way. My uncle had to drive me to the hospital where a proctologist had to go in a slice a layer of my butt-flesh off, although now I have extreme bleeding hemmorhoids out the wazoo!!!! And becuase we haven't yet upgraded the toilets in our house to ones that flush (yes my parents are living in the 1980's still...) we got an ear full from the hospitals psychotherapist. Heheh! There was one funny thing that happened out of this whole mess. When the doctor was inspecting my net I blasted(farted) gas in her face! The look in her eyes was fantastic, I thought she was going to cry!
`,
		`y helo thar buttsecks?`,
		`how strangely erotic`,
		`Its a trap!`,
		`Dont exist foo!`,
		`The cookies she didn't even want!`,
		`You just lost The Game`,
		`the gamefaqs spinoff luelinks "ETI"`,
		`Rumor's of LUE's demise are greatly exaggerated. As of this moment, it has over 30,000 active posts, making it one of the largest boards on GameFAQs. The rumors are mostly spread by disaffected people who did not or could not sign up for LUE before it went "excLUEsive."

However, rumors of how bad LUE used to be are not at all exaggerated in the least. It really was that bad. Even still it flirts with the obscene. I'm a member and I still drop by every now and then, but for the most part I'm done with GameFAQs.

But whatever else LUE is or was, "garbage" is not a good way to describe it. Aside from the stinking mounds of human depravity, it is one of the wittiest, most risque, and interesting boards I ever saw on GameFAQs.`,
		`hi this is simba from the lion king and uh i want to tell you about my latest business it's simba's shit pit simba's shit pit ah do you live in the jungle and if you live in the jungle are you constantly looking for places to shit well you're in luck cause it's simba's shit pit right off of root 80 simba shit pit ah are you an antelope ah who are when you're taking a shit a lion jumps out and bites your head off that happens to you all the time doesn't it that's why you have to go to simba shit pit and uh at simba's shit pit our slogan is simba shit pit we ain't lying`,
		`went to a local Mexican restaurant that has $1 tacos for 'Taco Tuesday.' It's a Chipotle style line where you choose hard or soft taco, chicken or beef, and you can add up to three "toppings" from the following: lettuce, salsa, corn, rice and beans, sour cream, or spicy sour cream. I think there were a few others but they are of no use to this story.

So I walk up to the start of the line and she asks me what I want. I say: "Hi. I'll have six tacos. Three hard and three soft. Three with chicken, and three with beef, but not so that all the chicken or beef are on the same style taco. I'll have any three toppings between lettuce, salsa, rice and beans, sour cream and spicy sour cream such that salsa and any style of sour cream is on all of them and that no two tacos have the exact same contents."

She looks at my face for a few seconds with a blank stare before saying, "Fuck this" and just walks to the back where the customer can't see, and then like 15 seconds later walks down the hallway and out the door and gets into her car. Then some guy comes out from the back, who I assume is the manager, with a concerned look on his face, apologizes and takes my order. Other employees were really confused.

Tacos were okay.
`,
		`i have a tiny dick and even smaller balls. the balls are above the dick also. the balls hand down around the dick like the ears on a basset hound. the dick stinks and i hate it. the balls do not smell. however, i also hate the balls. thank you
`,
		`Imagine 2 guys hanging out, outside of a Walmart, on average 6 hours a day. They just stand out there, randomly showing passersby photos of gay porn. Whether or not this is illegal is a gray area, and plus they have the town sheriff on their side so they don't get in trouble. If they ever do, they come back to that Walmart with a different mask on.`,
		`Light Under Earth`, `Love Unites Everything`, `Linger Until Eternity`, `Lost Under Echoes`, `Life Unfolding Endlessly`, `Library of Unknown Eras`,
		`Light Unfolding Everywhere`, `Looming Under Eternity`, `Lands Untouched Evermore`, `Longing Underneath Everything`, `Love Unfolds Elegantly`,
	)

	return phrases[rand.Intn(len(phrases))]
}

func makeBubble() string {
	// get a phrase
	phrase := getPhrase()

	// hold a list of lines here
	var allLines []string
	var currentLine []rune
	var charCount int

	// split the phrase into lines of 76 char each
	interval := 76

	for _, char := range phrase {
		// if its a newline, reset
		if char == '\n' {
			// append the empty line because we are gonna join on newline later anyway
			allLines = append(allLines, string(currentLine))
			currentLine = nil
			charCount = 0
		} else {
			// otherwise
			// add char to current line
			currentLine = append(currentLine, char)
			charCount++
		}

		// reset the line if we're at the length limit
		if charCount == interval {
			allLines = append(allLines, string(currentLine))
			currentLine = nil
			charCount = 0
		}
	}

	// add the last line contents
	if len(currentLine) > 0 {
		allLines = append(allLines, string(currentLine))
	}

	mergedLines := strings.Join(allLines, "\n")

	// left pad border
	var re = regexp.MustCompile(`(?m)^`)
	leftBorderLines := re.ReplaceAllString(mergedLines, `| `)
	// fmt.Println(leftBorderLines)

	// right pad border
	re = regexp.MustCompile(`(?m)^.*$`)
	rightBorderLines := re.ReplaceAllStringFunc(leftBorderLines, func(line string) string {
		runes := []rune(line) // account for unicode
		if len(runes) < interval+3 {
			// pad with spaces up to 78 chars, then append "|"
			padding := strings.Repeat(" ", interval+3-len(runes))
			return line + padding + " |"
		}
		// leave longer or equal lines alone
		return line
	})

	var finalString string

	finalString = finalString +
		`/-------------------------------------------------------------------------------\` + "\n" +
		rightBorderLines + "\n" +
		`\------------------------------------------------------------------------------/` + "\n" +
		`                           /` + "\n" +
		`                          /` + "\n"

	return finalString
}

func makeLUESHI() string {
	b64 := `ICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICBfXyAgIF9fCiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIC/igJlgwrdgXCwtLS1gIOKAngogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIC98LC3igJhgwq/Cr2BcKG8pX1wsLS0tLSwsLF8KICAgICAgICAgICAgICAgICAg4oCefirCr8KvYOKAnVwsICAgICAgICAgICAgIF/igJ5fICAgICAgICAgICAgKCBgXChvKSwsXy9gIMKvIMK3IG8gwrcgwrcgwrdvIGAtLAogICAgICAgICAgICAgICAgIC8gICAgICAo4oCcLCB+OyrigJnCr8Kvwq/igJ1cLCAoXywtLSBgYOKAnX4sICAgICBcIMK3IMK3IMK3IMK3IMK3IMK3IMK3IMK3IMK3IMK3IMK3IMK3IMK3IMK3IGBcCiAgICAgICAgICAgICAgICB8ICAgICAgLC9gLC0qfjt+ICAt4oCeLC8gKOKAmGAgYGApLyAgLCAgICBcICwvYCDCtyDCtyDCt1/igJ7igJ4swrcgwrcgwrcgwrcgwrcgwrcgwrcgwrcgwrcgwrcgXAogICAgICAgICAgICAgICAgfCAgICAgLyAsL2AsLS1cIFzigJlgY1wsLS0t4oCeMSDigLnigJlgLS0oXyAsLyAvIMK3IMK3IMK3LC9gICAgICApwrcgwrcgwrcgwrcgwrcgwrcgwrcgwrcgwrcgwrcgfAogICAgICAgICAgICAgICAgICkgICDCteKAmWAgIFwgKGMpIGDCryAgICAgIGApLCAgICwtfmAgICBcIMK3IMK3IMK3fCAgICAgJ+KAnVwsIMK3IMK3IMK3IMK3IMK3IMK3IMK3IMK3IMK3IMK3LwogICAgICAgICAgICAgICAvYCAsL19+LSwgIGA7O2DigJ4t4oCeLF9fLCAvLCBgYGAvICAgICAgICBgXMK3IMK3XCAgICAgICAgIGAqLSAsXyDCtyDCtyDCtyDCtyDCtyDCtywt4oCYCiAgICAgICAgICAgICAgIGAtL8KoOy0tO37igJkgYOKAnSotPSw9LV9g4oCdICwpICwvYCAgICAgICAgICBg4oC6IMK3IGBcLCAgICAgICAgICAgICDCr+KAnX4tLS0sLS1gCiAgIOKAnuKAniAgICAgICAgIF9fXCwgICDigJgsICAgICAgICAgYFzigJ5fLC8gLC87LTtfICAgICAgICAgIC9gIGAgfCAvICAgICAgICAgICAgIOKAni3igJwKICggICBgXCwtfipgwq8gICAgICBgwq9gIGB+LS1+Kn4tLS1+Oy9gLC1+KmBgYCotLSwgICAgICAvIMK3ICAgfCAgICAgICAgICwtLS1+KuKAnWAKICBcICAgICDigJ0qfi0sLOKAnuKAnl9fX1/igJ7igJ4sIC1+4oCdYMKvwq/Cr8KvLyAvICAgICAgICAgICDigJhcLCkgLCAgLyDCtyDCt3wgICAgICAgICBcCiAgIGDigJ3igJl+LSzigJ7igJ7igJ7igJ7igJ7igJ4sLH7igJhgYCAgICggICAgICAgLCBfX3wgfCAgICAgICAgLCBgXOKAnuKAni8gLC9gIC8gICAgICAgICB8CiAgICAgICAgICAgICAgICAgICAgICAgICBcYOKAneKAneKAnWAgICAgICBg4oCZfjstLOKAniwsXynigJ1gXy3igJggwrcgwrcgLyAgICAgICAgICB8CiAgICAgICAgICAoYCrigJ0tLOKAniwt4oCdwq/Cr+KAnWAtO+KAniAgICAgICAgICAgICwgLydgYCwtfuKAnWDCr8K3IMK3IMK3IMK3LyAgICAgICAgICB8CiAgICAgICAgICAgfCDCtyDCtygsOy09PT0t4oCeLCBgXCwgICAgICAsLWB8ICAgIC8gwrcgwrcgwrcgfCDCtywvYCAgICAgICAgICAgfAogICAgICAgICAgICBcIMK3IMK3XCwgICAgICAgICBcXCAgIGBcICAgKSAgLyAgIC8gwrcgwrdcIMK3IMK3YH4sXyAgICAgICAgIC8KICAgICAgICAgICAgIFxcLF9gfiAsX+KAniwgLSpcXCwgYC8sL+KAni9gICwvIMK3IMK3IMK3YOKAmTstIOKAnl8gwrcgwq8tLCAgICAgICAvLAogICAgICAgICAgICAgIGBcLCxg4oCdfCDCtyDCt2AtLOKAnl/igJ4pKeKAmWAiYCAgLC9gX+KAnix+KuKAmWAgICAgICggICAsICxgKSAsLTsgYOKAmVwsCiAgICAgICAgICAgICAgICAgYCotXCDCtyDCtyDCtyBgfi0tLS1+KmAgLyAgICAgICAgICAgICAgICDigJzigJ1+4oCdYCAgL2AgKCBfIMK3KQogICAgICAgICAgICAgICwgwqwtLC0tXCDCtyDCtyDCtyDCtyDCtyDCtyDCtyDCtyAvICAgICAgICAgICAgICAgICwsLeKAnGAgICAg4oCZLeKAniwt4oCYCiAgICAgICAgICAgICAgfCB8wrcgwrcqwrdcIMK3IMK3IMK3IMK3IMK3IMK3IMK3Xy8g4oCeX19fX+KAnuKAnuKAnuKAniwtLS1+KmAKICAgICAgICAgICAgIC8gLyDCtyDCtyDCtyBgfi3igJ7igJ7igJ7igJ7igJ4sIDs7YCwsIC0tYAogICAgICAgICAgICB8IHwgwrcgwrcgwrcgwrd8wq8gwrcgLC8gwq8KICAgICAgICAgICAgIFwsXCxfLOKAniAvLS0tfmAK`
	data, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		log.Fatal(err)
	}

	textBubble := makeBubble()
	lueshiStr := string(data)
	finalLUEshi := textBubble + lueshiStr

	return finalLUEshi
}

func printLUEshi() {
	lueshiStr := makeLUESHI()
	fmt.Println(lueshiStr)
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	lueshiStr := makeLUESHI()
	io.WriteString(w, lueshiStr)
}

func runLUEshiServer(port string) {
	// curl http://localhost:4242
	fmt.Printf("running LUEshi server on port %s\n", port)

	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	// add more endpoints here

	err := http.ListenAndServe(":"+port, mux)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func main() {
	printVersion := flag.Bool("v", false, "print version information")
	debug := flag.Bool("debug", false, "run debug")
	runServer := flag.Bool("s", false, "run in server mode")
	serverPort := flag.String("p", "4242", "port to use in server mode")
	flag.Parse()
	// cliArgs := flag.Args() // all positional args passed

	if *printVersion {
		PrintVersionAndQuit()
	}

	if *debug {
		Debug()
	}

	if *runServer {
		runLUEshiServer(*serverPort)
	} else {
		// default action if no args are passed
		printLUEshi()
	}

}
