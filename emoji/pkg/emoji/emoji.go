package emoji

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

func (r Range) Length() int {
	return r.end - r.start + 1
}

type EmojiDef struct {
	ranges     []Range
	emojiCount int
}

func parseDefLine(s string) (*Range, error) {
	hashIndex := strings.Index(s, "#")
	if hashIndex != -1 {
		s = s[0:hashIndex]
	}
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return nil, nil
	} else {
		re, err := regexp.Compile(`^(?P<start>[A-F0-9]+)(?:\.\.(?P<end>[A-F0-9]+))?`)
		if err != nil {
			return nil, err
		}

		submatches := re.FindStringSubmatch(s)

		if len(submatches) == 3 {
			n := submatches[1]
			start, err := strconv.ParseInt(n, 16, 32)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("failed parsing emoji '%s' in line '%s'", n, s))
			}

			end := start

			if len(submatches[2]) > 0 {
				n := submatches[2]
				end, err = strconv.ParseInt(n, 16, 32)
				if err != nil {
					return nil, errors.New(fmt.Sprintf("failed parsing emoji '%s' in line '%s'", n, s))
				}
			}

			return &Range{int(start), int(end)}, nil
		} else {
			return nil, errors.New(fmt.Sprintf("failed parsing line '%s'", s))
		}
	}
}

func parseDef(s string) EmojiDef {
	var ranges []Range
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		r, err := parseDefLine(scanner.Text())
		if err != nil {
			panic(err)
		}

		if r != nil {
			ranges = append(ranges, *r)
		}
	}

	count := 0
	for _, r := range ranges {
		count += r.Length()
	}

	return EmojiDef{ranges, count}
}

var emojiDef = parseDef(emojiDefString)

func (ed EmojiDef) RandomEmoji(rnd *rand.Rand) int {
	n := int(rnd.Int31n(int32(ed.emojiCount)))
	for _, curRange := range ed.ranges {
		if n < curRange.Length() {
			return curRange.start + n
		}
		n -= curRange.Length()
	}
	return -1
}

func RandomEmoji(rnd *rand.Rand) int {
	return emojiDef.RandomEmoji(rnd)
}

// from http://www.unicode.org/Public/emoji/2.0//emoji-data.txt
const emojiDefString = `
# Emoji Data for UTR #51
#
# File:    emoji-data.txt
# Version: 2.0
# Date:    2015-11-11
#
# Copyright (c) 2015 Unicode, Inc.
# For terms of use, see http://www.unicode.org/terms_of_use.html
# For documentation and usage, see http://www.unicode.org/reports/tr51/
#

# Warning: the format has changed from Version 1.0
# Format:
# codepoint(s) ; property=Yes # [count] (character(s)) name

# ================================================

# All omitted code points have Emoji=No
# @missing: 0000..10FFFF  ; Emoji ; No

0023          ; Emoji                #   [1] (#️)      NUMBER SIGN
002A          ; Emoji                #   [1] (*)       ASTERISK
0030..0039    ; Emoji                #  [10] (0️..9️)  DIGIT ZERO..DIGIT NINE
00A9          ; Emoji                #   [1] (©️)      COPYRIGHT SIGN
00AE          ; Emoji                #   [1] (®️)      REGISTERED SIGN
203C          ; Emoji                #   [1] (‼️)      DOUBLE EXCLAMATION MARK
2049          ; Emoji                #   [1] (⁉️)      EXCLAMATION QUESTION MARK
2122          ; Emoji                #   [1] (™️)      TRADE MARK SIGN
2139          ; Emoji                #   [1] (ℹ️)      INFORMATION SOURCE
2194..2199    ; Emoji                #   [6] (↔️..↙️)  LEFT RIGHT ARROW..SOUTH WEST ARROW
21A9..21AA    ; Emoji                #   [2] (↩️..↪️)  LEFTWARDS ARROW WITH HOOK..RIGHTWARDS ARROW WITH HOOK
231A..231B    ; Emoji                #   [2] (⌚️..⌛️)  WATCH..HOURGLASS
2328          ; Emoji                #   [1] (⌨)       KEYBOARD
23CF          ; Emoji                #   [1] (⏏)       EJECT SYMBOL
23E9..23F3    ; Emoji                #  [11] (⏩..⏳)    BLACK RIGHT-POINTING DOUBLE TRIANGLE..HOURGLASS WITH FLOWING SAND
23F8..23FA    ; Emoji                #   [3] (⏸..⏺)    DOUBLE VERTICAL BAR..BLACK CIRCLE FOR RECORD
24C2          ; Emoji                #   [1] (Ⓜ️)      CIRCLED LATIN CAPITAL LETTER M
25AA..25AB    ; Emoji                #   [2] (▪️..▫️)  BLACK SMALL SQUARE..WHITE SMALL SQUARE
25B6          ; Emoji                #   [1] (▶️)      BLACK RIGHT-POINTING TRIANGLE
25C0          ; Emoji                #   [1] (◀️)      BLACK LEFT-POINTING TRIANGLE
25FB..25FE    ; Emoji                #   [4] (◻️..◾️)  WHITE MEDIUM SQUARE..BLACK MEDIUM SMALL SQUARE
2600..2604    ; Emoji                #   [5] (☀️..☄)   BLACK SUN WITH RAYS..COMET
260E          ; Emoji                #   [1] (☎️)      BLACK TELEPHONE
2611          ; Emoji                #   [1] (☑️)      BALLOT BOX WITH CHECK
2614..2615    ; Emoji                #   [2] (☔️..☕️)  UMBRELLA WITH RAIN DROPS..HOT BEVERAGE
2618          ; Emoji                #   [1] (☘)       SHAMROCK
261D          ; Emoji                #   [1] (☝️)      WHITE UP POINTING INDEX
2620          ; Emoji                #   [1] (☠)       SKULL AND CROSSBONES
2622..2623    ; Emoji                #   [2] (☢..☣)    RADIOACTIVE SIGN..BIOHAZARD SIGN
2626          ; Emoji                #   [1] (☦)       ORTHODOX CROSS
262A          ; Emoji                #   [1] (☪)       STAR AND CRESCENT
262E..262F    ; Emoji                #   [2] (☮..☯)    PEACE SYMBOL..YIN YANG
2638..263A    ; Emoji                #   [3] (☸..☺️)   WHEEL OF DHARMA..WHITE SMILING FACE
2648..2653    ; Emoji                #  [12] (♈️..♓️)  ARIES..PISCES
2660          ; Emoji                #   [1] (♠️)      BLACK SPADE SUIT
2663          ; Emoji                #   [1] (♣️)      BLACK CLUB SUIT
2665..2666    ; Emoji                #   [2] (♥️..♦️)  BLACK HEART SUIT..BLACK DIAMOND SUIT
2668          ; Emoji                #   [1] (♨️)      HOT SPRINGS
267B          ; Emoji                #   [1] (♻️)      BLACK UNIVERSAL RECYCLING SYMBOL
267F          ; Emoji                #   [1] (♿️)      WHEELCHAIR SYMBOL
2692..2694    ; Emoji                #   [3] (⚒..⚔)    HAMMER AND PICK..CROSSED SWORDS
2696..2697    ; Emoji                #   [2] (⚖..⚗)    SCALES..ALEMBIC
2699          ; Emoji                #   [1] (⚙)       GEAR
269B..269C    ; Emoji                #   [2] (⚛..⚜)    ATOM SYMBOL..FLEUR-DE-LIS
26A0..26A1    ; Emoji                #   [2] (⚠️..⚡️)  WARNING SIGN..HIGH VOLTAGE SIGN
26AA..26AB    ; Emoji                #   [2] (⚪️..⚫️)  MEDIUM WHITE CIRCLE..MEDIUM BLACK CIRCLE
26B0..26B1    ; Emoji                #   [2] (⚰..⚱)    COFFIN..FUNERAL URN
26BD..26BE    ; Emoji                #   [2] (⚽️..⚾️)  SOCCER BALL..BASEBALL
26C4..26C5    ; Emoji                #   [2] (⛄️..⛅️)  SNOWMAN WITHOUT SNOW..SUN BEHIND CLOUD
26C8          ; Emoji                #   [1] (⛈)       THUNDER CLOUD AND RAIN
26CE..26CF    ; Emoji                #   [2] (⛎..⛏)    OPHIUCHUS..PICK
26D1          ; Emoji                #   [1] (⛑)       HELMET WITH WHITE CROSS
26D3..26D4    ; Emoji                #   [2] (⛓..⛔️)   CHAINS..NO ENTRY
26E9..26EA    ; Emoji                #   [2] (⛩..⛪️)   SHINTO SHRINE..CHURCH
26F0..26F5    ; Emoji                #   [6] (⛰..⛵️)   MOUNTAIN..SAILBOAT
26F7..26FA    ; Emoji                #   [4] (⛷..⛺️)   SKIER..TENT
26FD          ; Emoji                #   [1] (⛽️)      FUEL PUMP
2702          ; Emoji                #   [1] (✂️)      BLACK SCISSORS
2705          ; Emoji                #   [1] (✅)       WHITE HEAVY CHECK MARK
2708..270D    ; Emoji                #   [6] (✈️..✍)   AIRPLANE..WRITING HAND
270F          ; Emoji                #   [1] (✏️)      PENCIL
2712          ; Emoji                #   [1] (✒️)      BLACK NIB
2714          ; Emoji                #   [1] (✔️)      HEAVY CHECK MARK
2716          ; Emoji                #   [1] (✖️)      HEAVY MULTIPLICATION X
271D          ; Emoji                #   [1] (✝)       LATIN CROSS
2721          ; Emoji                #   [1] (✡)       STAR OF DAVID
2728          ; Emoji                #   [1] (✨)       SPARKLES
2733..2734    ; Emoji                #   [2] (✳️..✴️)  EIGHT SPOKED ASTERISK..EIGHT POINTED BLACK STAR
2744          ; Emoji                #   [1] (❄️)      SNOWFLAKE
2747          ; Emoji                #   [1] (❇️)      SPARKLE
274C          ; Emoji                #   [1] (❌)       CROSS MARK
274E          ; Emoji                #   [1] (❎)       NEGATIVE SQUARED CROSS MARK
2753..2755    ; Emoji                #   [3] (❓..❕)    BLACK QUESTION MARK ORNAMENT..WHITE EXCLAMATION MARK ORNAMENT
2757          ; Emoji                #   [1] (❗️)      HEAVY EXCLAMATION MARK SYMBOL
2763..2764    ; Emoji                #   [2] (❣..❤️)   HEAVY HEART EXCLAMATION MARK ORNAMENT..HEAVY BLACK HEART
2795..2797    ; Emoji                #   [3] (➕..➗)    HEAVY PLUS SIGN..HEAVY DIVISION SIGN
27A1          ; Emoji                #   [1] (➡️)      BLACK RIGHTWARDS ARROW
27B0          ; Emoji                #   [1] (➰)       CURLY LOOP
27BF          ; Emoji                #   [1] (➿)       DOUBLE CURLY LOOP
2934..2935    ; Emoji                #   [2] (⤴️..⤵️)  ARROW POINTING RIGHTWARDS THEN CURVING UPWARDS..ARROW POINTING RIGHTWARDS THEN CURVING DOWNWARDS
2B05..2B07    ; Emoji                #   [3] (⬅️..⬇️)  LEFTWARDS BLACK ARROW..DOWNWARDS BLACK ARROW
2B1B..2B1C    ; Emoji                #   [2] (⬛️..⬜️)  BLACK LARGE SQUARE..WHITE LARGE SQUARE
2B50          ; Emoji                #   [1] (⭐️)      WHITE MEDIUM STAR
2B55          ; Emoji                #   [1] (⭕️)      HEAVY LARGE CIRCLE
3030          ; Emoji                #   [1] (〰️)      WAVY DASH
303D          ; Emoji                #   [1] (〽️)      PART ALTERNATION MARK
3297          ; Emoji                #   [1] (㊗️)      CIRCLED IDEOGRAPH CONGRATULATION
3299          ; Emoji                #   [1] (㊙️)      CIRCLED IDEOGRAPH SECRET
1F004         ; Emoji                #   [1] (🀄️)      MAHJONG TILE RED DRAGON
1F0CF         ; Emoji                #   [1] (🃏)       PLAYING CARD BLACK JOKER
1F170..1F171  ; Emoji                #   [2] (🅰️..🅱️)  NEGATIVE SQUARED LATIN CAPITAL LETTER A..NEGATIVE SQUARED LATIN CAPITAL LETTER B
1F17E..1F17F  ; Emoji                #   [2] (🅾️..🅿️)  NEGATIVE SQUARED LATIN CAPITAL LETTER O..NEGATIVE SQUARED LATIN CAPITAL LETTER P
1F18E         ; Emoji                #   [1] (🆎)       NEGATIVE SQUARED AB
1F191..1F19A  ; Emoji                #  [10] (🆑..🆚)    SQUARED CL..SQUARED VS
1F1E6..1F1FF  ; Emoji                #  [26] (🇦..🇿)    REGIONAL INDICATOR SYMBOL LETTER A..REGIONAL INDICATOR SYMBOL LETTER Z
1F201..1F202  ; Emoji                #   [2] (🈁..🈂️)   SQUARED KATAKANA KOKO..SQUARED KATAKANA SA
1F21A         ; Emoji                #   [1] (🈚️)      SQUARED CJK UNIFIED IDEOGRAPH-7121
1F22F         ; Emoji                #   [1] (🈯️)      SQUARED CJK UNIFIED IDEOGRAPH-6307
1F232..1F23A  ; Emoji                #   [9] (🈲..🈺)    SQUARED CJK UNIFIED IDEOGRAPH-7981..SQUARED CJK UNIFIED IDEOGRAPH-55B6
1F250..1F251  ; Emoji                #   [2] (🉐..🉑)    CIRCLED IDEOGRAPH ADVANTAGE..CIRCLED IDEOGRAPH ACCEPT
1F300..1F321  ; Emoji                #  [34] (🌀..🌡)    CYCLONE..THERMOMETER
1F324..1F393  ; Emoji                # [112] (🌤..🎓)    WHITE SUN WITH SMALL CLOUD..GRADUATION CAP
1F396..1F397  ; Emoji                #   [2] (🎖..🎗)    MILITARY MEDAL..REMINDER RIBBON
1F399..1F39B  ; Emoji                #   [3] (🎙..🎛)    STUDIO MICROPHONE..CONTROL KNOBS
1F39E..1F3F0  ; Emoji                #  [83] (🎞..🏰)    FILM FRAMES..EUROPEAN CASTLE
1F3F3..1F3F5  ; Emoji                #   [3] (🏳..🏵)    WAVING WHITE FLAG..ROSETTE
1F3F7..1F4FD  ; Emoji                # [263] (🏷..📽)    LABEL..FILM PROJECTOR
1F4FF..1F53D  ; Emoji                #  [63] (📿..🔽)    PRAYER BEADS..DOWN-POINTING SMALL RED TRIANGLE
1F549..1F54E  ; Emoji                #   [6] (🕉..🕎)    OM SYMBOL..MENORAH WITH NINE BRANCHES
1F550..1F567  ; Emoji                #  [24] (🕐..🕧)    CLOCK FACE ONE OCLOCK..CLOCK FACE TWELVE-THIRTY
1F56F..1F570  ; Emoji                #   [2] (🕯..🕰)    CANDLE..MANTELPIECE CLOCK
1F573..1F579  ; Emoji                #   [7] (🕳..🕹)    HOLE..JOYSTICK
1F587         ; Emoji                #   [1] (🖇)       LINKED PAPERCLIPS
1F58A..1F58D  ; Emoji                #   [4] (🖊..🖍)    LOWER LEFT BALLPOINT PEN..LOWER LEFT CRAYON
1F590         ; Emoji                #   [1] (🖐)       RAISED HAND WITH FINGERS SPLAYED
1F595..1F596  ; Emoji                #   [2] (🖕..🖖)    REVERSED HAND WITH MIDDLE FINGER EXTENDED..RAISED HAND WITH PART BETWEEN MIDDLE AND RING FINGERS
1F5A5         ; Emoji                #   [1] (🖥)       DESKTOP COMPUTER
1F5A8         ; Emoji                #   [1] (🖨)       PRINTER
1F5B1..1F5B2  ; Emoji                #   [2] (🖱..🖲)    THREE BUTTON MOUSE..TRACKBALL
1F5BC         ; Emoji                #   [1] (🖼)       FRAME WITH PICTURE
1F5C2..1F5C4  ; Emoji                #   [3] (🗂..🗄)    CARD INDEX DIVIDERS..FILE CABINET
1F5D1..1F5D3  ; Emoji                #   [3] (🗑..🗓)    WASTEBASKET..SPIRAL CALENDAR PAD
1F5DC..1F5DE  ; Emoji                #   [3] (🗜..🗞)    COMPRESSION..ROLLED-UP NEWSPAPER
1F5E1         ; Emoji                #   [1] (🗡)       DAGGER KNIFE
1F5E3         ; Emoji                #   [1] (🗣)       SPEAKING HEAD IN SILHOUETTE
1F5E8         ; Emoji                #   [1] (🗨)       LEFT SPEECH BUBBLE
1F5EF         ; Emoji                #   [1] (🗯)       RIGHT ANGER BUBBLE
1F5F3         ; Emoji                #   [1] (🗳)       BALLOT BOX WITH BALLOT
1F5FA..1F64F  ; Emoji                #  [86] (🗺..🙏)    WORLD MAP..PERSON WITH FOLDED HANDS
1F680..1F6C5  ; Emoji                #  [70] (🚀..🛅)    ROCKET..LEFT LUGGAGE
1F6CB..1F6D0  ; Emoji                #   [6] (🛋..🛐)    COUCH AND LAMP..PLACE OF WORSHIP
1F6E0..1F6E5  ; Emoji                #   [6] (🛠..🛥)    HAMMER AND WRENCH..MOTOR BOAT
1F6E9         ; Emoji                #   [1] (🛩)       SMALL AIRPLANE
1F6EB..1F6EC  ; Emoji                #   [2] (🛫..🛬)    AIRPLANE DEPARTURE..AIRPLANE ARRIVING
1F6F0         ; Emoji                #   [1] (🛰)       SATELLITE
1F6F3         ; Emoji                #   [1] (🛳)       PASSENGER SHIP
1F910..1F918  ; Emoji                #   [9] (🤐..🤘)    ZIPPER-MOUTH FACE..SIGN OF THE HORNS
1F980..1F984  ; Emoji                #   [5] (🦀..🦄)    CRAB..UNICORN FACE
1F9C0         ; Emoji                #   [1] (🧀)       CHEESE WEDGE

# Total code points: 1051

# ================================================

# All omitted code points have Emoji_Presentation=No
# @missing: 0000..10FFFF  ; Emoji_Presentation ; No

231A..231B    ; Emoji_Presentation   #   [2] (⌚️..⌛️)  WATCH..HOURGLASS
23E9..23EC    ; Emoji_Presentation   #   [4] (⏩..⏬)    BLACK RIGHT-POINTING DOUBLE TRIANGLE..BLACK DOWN-POINTING DOUBLE TRIANGLE
23F0          ; Emoji_Presentation   #   [1] (⏰)       ALARM CLOCK
23F3          ; Emoji_Presentation   #   [1] (⏳)       HOURGLASS WITH FLOWING SAND
25FD..25FE    ; Emoji_Presentation   #   [2] (◽️..◾️)  WHITE MEDIUM SMALL SQUARE..BLACK MEDIUM SMALL SQUARE
2614..2615    ; Emoji_Presentation   #   [2] (☔️..☕️)  UMBRELLA WITH RAIN DROPS..HOT BEVERAGE
2648..2653    ; Emoji_Presentation   #  [12] (♈️..♓️)  ARIES..PISCES
267F          ; Emoji_Presentation   #   [1] (♿️)      WHEELCHAIR SYMBOL
2693          ; Emoji_Presentation   #   [1] (⚓️)      ANCHOR
26A1          ; Emoji_Presentation   #   [1] (⚡️)      HIGH VOLTAGE SIGN
26AA..26AB    ; Emoji_Presentation   #   [2] (⚪️..⚫️)  MEDIUM WHITE CIRCLE..MEDIUM BLACK CIRCLE
26BD..26BE    ; Emoji_Presentation   #   [2] (⚽️..⚾️)  SOCCER BALL..BASEBALL
26C4..26C5    ; Emoji_Presentation   #   [2] (⛄️..⛅️)  SNOWMAN WITHOUT SNOW..SUN BEHIND CLOUD
26CE          ; Emoji_Presentation   #   [1] (⛎)       OPHIUCHUS
26D4          ; Emoji_Presentation   #   [1] (⛔️)      NO ENTRY
26EA          ; Emoji_Presentation   #   [1] (⛪️)      CHURCH
26F2..26F3    ; Emoji_Presentation   #   [2] (⛲️..⛳️)  FOUNTAIN..FLAG IN HOLE
26F5          ; Emoji_Presentation   #   [1] (⛵️)      SAILBOAT
26FA          ; Emoji_Presentation   #   [1] (⛺️)      TENT
26FD          ; Emoji_Presentation   #   [1] (⛽️)      FUEL PUMP
2705          ; Emoji_Presentation   #   [1] (✅)       WHITE HEAVY CHECK MARK
270A..270B    ; Emoji_Presentation   #   [2] (✊..✋)    RAISED FIST..RAISED HAND
2728          ; Emoji_Presentation   #   [1] (✨)       SPARKLES
274C          ; Emoji_Presentation   #   [1] (❌)       CROSS MARK
274E          ; Emoji_Presentation   #   [1] (❎)       NEGATIVE SQUARED CROSS MARK
2753..2755    ; Emoji_Presentation   #   [3] (❓..❕)    BLACK QUESTION MARK ORNAMENT..WHITE EXCLAMATION MARK ORNAMENT
2757          ; Emoji_Presentation   #   [1] (❗️)      HEAVY EXCLAMATION MARK SYMBOL
2795..2797    ; Emoji_Presentation   #   [3] (➕..➗)    HEAVY PLUS SIGN..HEAVY DIVISION SIGN
27B0          ; Emoji_Presentation   #   [1] (➰)       CURLY LOOP
27BF          ; Emoji_Presentation   #   [1] (➿)       DOUBLE CURLY LOOP
2B1B..2B1C    ; Emoji_Presentation   #   [2] (⬛️..⬜️)  BLACK LARGE SQUARE..WHITE LARGE SQUARE
2B50          ; Emoji_Presentation   #   [1] (⭐️)      WHITE MEDIUM STAR
2B55          ; Emoji_Presentation   #   [1] (⭕️)      HEAVY LARGE CIRCLE
1F004         ; Emoji_Presentation   #   [1] (🀄️)      MAHJONG TILE RED DRAGON
1F0CF         ; Emoji_Presentation   #   [1] (🃏)       PLAYING CARD BLACK JOKER
1F18E         ; Emoji_Presentation   #   [1] (🆎)       NEGATIVE SQUARED AB
1F191..1F19A  ; Emoji_Presentation   #  [10] (🆑..🆚)    SQUARED CL..SQUARED VS
1F1E6..1F1FF  ; Emoji_Presentation   #  [26] (🇦..🇿)    REGIONAL INDICATOR SYMBOL LETTER A..REGIONAL INDICATOR SYMBOL LETTER Z
1F201         ; Emoji_Presentation   #   [1] (🈁)       SQUARED KATAKANA KOKO
1F21A         ; Emoji_Presentation   #   [1] (🈚️)      SQUARED CJK UNIFIED IDEOGRAPH-7121
1F22F         ; Emoji_Presentation   #   [1] (🈯️)      SQUARED CJK UNIFIED IDEOGRAPH-6307
1F232..1F236  ; Emoji_Presentation   #   [5] (🈲..🈶)    SQUARED CJK UNIFIED IDEOGRAPH-7981..SQUARED CJK UNIFIED IDEOGRAPH-6709
1F238..1F23A  ; Emoji_Presentation   #   [3] (🈸..🈺)    SQUARED CJK UNIFIED IDEOGRAPH-7533..SQUARED CJK UNIFIED IDEOGRAPH-55B6
1F250..1F251  ; Emoji_Presentation   #   [2] (🉐..🉑)    CIRCLED IDEOGRAPH ADVANTAGE..CIRCLED IDEOGRAPH ACCEPT
1F300..1F320  ; Emoji_Presentation   #  [33] (🌀..🌠)    CYCLONE..SHOOTING STAR
1F32D..1F335  ; Emoji_Presentation   #   [9] (🌭..🌵)    HOT DOG..CACTUS
1F337..1F37C  ; Emoji_Presentation   #  [70] (🌷..🍼)    TULIP..BABY BOTTLE
1F37E..1F393  ; Emoji_Presentation   #  [22] (🍾..🎓)    BOTTLE WITH POPPING CORK..GRADUATION CAP
1F3A0..1F3CA  ; Emoji_Presentation   #  [43] (🎠..🏊)    CAROUSEL HORSE..SWIMMER
1F3CF..1F3D3  ; Emoji_Presentation   #   [5] (🏏..🏓)    CRICKET BAT AND BALL..TABLE TENNIS PADDLE AND BALL
1F3E0..1F3F0  ; Emoji_Presentation   #  [17] (🏠..🏰)    HOUSE BUILDING..EUROPEAN CASTLE
1F3F4         ; Emoji_Presentation   #   [1] (🏴)       WAVING BLACK FLAG
1F3F8..1F43E  ; Emoji_Presentation   #  [71] (🏸..🐾)    BADMINTON RACQUET AND SHUTTLECOCK..PAW PRINTS
1F440         ; Emoji_Presentation   #   [1] (👀)       EYES
1F442..1F4FC  ; Emoji_Presentation   # [187] (👂..📼)    EAR..VIDEOCASSETTE
1F4FF..1F53D  ; Emoji_Presentation   #  [63] (📿..🔽)    PRAYER BEADS..DOWN-POINTING SMALL RED TRIANGLE
1F54B..1F54E  ; Emoji_Presentation   #   [4] (🕋..🕎)    KAABA..MENORAH WITH NINE BRANCHES
1F550..1F567  ; Emoji_Presentation   #  [24] (🕐..🕧)    CLOCK FACE ONE OCLOCK..CLOCK FACE TWELVE-THIRTY
1F595..1F596  ; Emoji_Presentation   #   [2] (🖕..🖖)    REVERSED HAND WITH MIDDLE FINGER EXTENDED..RAISED HAND WITH PART BETWEEN MIDDLE AND RING FINGERS
1F5FB..1F64F  ; Emoji_Presentation   #  [85] (🗻..🙏)    MOUNT FUJI..PERSON WITH FOLDED HANDS
1F680..1F6C5  ; Emoji_Presentation   #  [70] (🚀..🛅)    ROCKET..LEFT LUGGAGE
1F6CC         ; Emoji_Presentation   #   [1] (🛌)       SLEEPING ACCOMMODATION
1F6D0         ; Emoji_Presentation   #   [1] (🛐)       PLACE OF WORSHIP
1F6EB..1F6EC  ; Emoji_Presentation   #   [2] (🛫..🛬)    AIRPLANE DEPARTURE..AIRPLANE ARRIVING
1F910..1F918  ; Emoji_Presentation   #   [9] (🤐..🤘)    ZIPPER-MOUTH FACE..SIGN OF THE HORNS
1F980..1F984  ; Emoji_Presentation   #   [5] (🦀..🦄)    CRAB..UNICORN FACE
1F9C0         ; Emoji_Presentation   #   [1] (🧀)       CHEESE WEDGE

# Total code points: 838

# ================================================

# All omitted code points have Emoji_Modifier=No
# @missing: 0000..10FFFF  ; Emoji_Modifier ; No

1F3FB..1F3FF  ; Emoji_Modifier       #   [5] (🏻..🏿)    EMOJI MODIFIER FITZPATRICK TYPE-1-2..EMOJI MODIFIER FITZPATRICK TYPE-6

# Total code points: 5

# ================================================

# All omitted code points have Emoji_Modifier_Base=No
# @missing: 0000..10FFFF  ; Emoji_Modifier_Base ; No

261D          ; Emoji_Modifier_Base  #   [1] (☝️)      WHITE UP POINTING INDEX
26F9          ; Emoji_Modifier_Base  #   [1] (⛹)       PERSON WITH BALL
270A..270D    ; Emoji_Modifier_Base  #   [4] (✊..✍)    RAISED FIST..WRITING HAND
1F385         ; Emoji_Modifier_Base  #   [1] (🎅)       FATHER CHRISTMAS
1F3C3..1F3C4  ; Emoji_Modifier_Base  #   [2] (🏃..🏄)    RUNNER..SURFER
1F3CA..1F3CB  ; Emoji_Modifier_Base  #   [2] (🏊..🏋)    SWIMMER..WEIGHT LIFTER
1F442..1F443  ; Emoji_Modifier_Base  #   [2] (👂..👃)    EAR..NOSE
1F446..1F450  ; Emoji_Modifier_Base  #  [11] (👆..👐)    WHITE UP POINTING BACKHAND INDEX..OPEN HANDS SIGN
1F466..1F469  ; Emoji_Modifier_Base  #   [4] (👦..👩)    BOY..WOMAN
1F46E         ; Emoji_Modifier_Base  #   [1] (👮)       POLICE OFFICER
1F470..1F478  ; Emoji_Modifier_Base  #   [9] (👰..👸)    BRIDE WITH VEIL..PRINCESS
1F47C         ; Emoji_Modifier_Base  #   [1] (👼)       BABY ANGEL
1F481..1F483  ; Emoji_Modifier_Base  #   [3] (💁..💃)    INFORMATION DESK PERSON..DANCER
1F485..1F487  ; Emoji_Modifier_Base  #   [3] (💅..💇)    NAIL POLISH..HAIRCUT
1F4AA         ; Emoji_Modifier_Base  #   [1] (💪)       FLEXED BICEPS
1F575         ; Emoji_Modifier_Base  #   [1] (🕵)       SLEUTH OR SPY
1F590         ; Emoji_Modifier_Base  #   [1] (🖐)       RAISED HAND WITH FINGERS SPLAYED
1F595..1F596  ; Emoji_Modifier_Base  #   [2] (🖕..🖖)    REVERSED HAND WITH MIDDLE FINGER EXTENDED..RAISED HAND WITH PART BETWEEN MIDDLE AND RING FINGERS
1F645..1F647  ; Emoji_Modifier_Base  #   [3] (🙅..🙇)    FACE WITH NO GOOD GESTURE..PERSON BOWING DEEPLY
1F64B..1F64F  ; Emoji_Modifier_Base  #   [5] (🙋..🙏)    HAPPY PERSON RAISING ONE HAND..PERSON WITH FOLDED HANDS
1F6A3         ; Emoji_Modifier_Base  #   [1] (🚣)       ROWBOAT
1F6B4..1F6B6  ; Emoji_Modifier_Base  #   [3] (🚴..🚶)    BICYCLIST..PEDESTRIAN
1F6C0         ; Emoji_Modifier_Base  #   [1] (🛀)       BATH
1F918         ; Emoji_Modifier_Base  #   [1] (🤘)       SIGN OF THE HORNS

# Total code points: 64
`
