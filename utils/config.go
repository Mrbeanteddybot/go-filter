// (c) Jisin0

package utils

import "github.com/PaulSonOfLars/gotgbot/v2"

var TEXT map[string]string = map[string]string{

	"START": `
๐ง๐พ๐ %v ๐ญ๐๐ผ๐พ ๐๐ ๐๐พ๐พ๐ ๐๐๐ ๐
๐๐ %v ๐บ๐ ๐บ๐๐พ๐๐๐๐พ ๐ฟ๐๐๐๐พ๐ ๐ป๐๐ ๐๐๐๐ ๐๐๐๐ป๐บ๐ ๐ฟ๐๐๐๐พ๐ ๐๐๐๐๐๐๐

๐จ ๐ผ๐บ๐ ๐๐บ๐๐พ ๐บ ๐ผ๐๐๐๐๐ ๐๐พ๐๐๐ ๐ฟ๐๐ ๐บ ๐๐๐๐ฝ ๐๐ ๐บ๐๐ ๐ผ๐๐บ๐. ๐ข๐๐พ๐ผ๐ ๐๐ ๐๐พ๐๐ ๐๐พ๐๐ ๐ฟ๐๐ ๐๐๐๐พ ๐ฝ๐พ๐๐บ๐๐๐
	`,
	"ABOUT": `
<b>โฅ ๐ซ๐บ๐๐๐๐บ๐๐พ</b> : <a href='https://go.dev'>๐๐ผ</a>
<b>โฅ ๐ฅ๐๐บ๐๐พ๐๐๐</b> : <a href='github.com/PaulSonOfLars/gotgbot'>๐๐ผ๐๐ด๐ฏ๐ผ๐</a>
<b>โฅ ๐ฒ๐พ๐๐๐พ๐</b> : <a href='heroku.com'>๐๐ฒ๐ฟ๐ผ๐ธ๐</a>
<b>โฅ ๐ฃ๐บ๐๐บ๐ป๐บ๐๐พ</b> : <a href='mongodb.org'>๐ ๐ผ๐ป๐ด๐ผ๐๐</a>
<b>โฅ ๐ฃ๐พ๐๐พ๐๐๐๐พ๐</b> : <a href='t.me/About_beantg'>๐๐ฒ๐ฎ๐ป</a>

	`,

	"MF": `
<b>Mแดษดแดแดส าษชสแดแดสs แดสสแดแดก สแดแด แดแด sแดแด แด แดแดsแดแดแด าษชสแดแดสs แดแดสแดส แดสแดษด แดสแด แดแดแดแดแดแดแดษชแด แดษดแดs. Fษชสแดแดสs แดแดษด สแด แดา แดแดxแด/แดสแดแดแด/แดแดแดแดแดแดษดแด/แดแดแดษชแด/แดษดษชแดแดแดษชแดษด/แด ษชแดแดแด .</b>

<b><u>Nแดแดก าษชสแดแดส :</u></b>

<u>Fแดสแดแดแด</u>
<code>/filter "keyword" text</code> or
Rแดแดสส แดแด แด แดแดssแดษขแด -><code>/filter "keyword"</code>
<u>Usแดษขแด</u>
<code>/filter "hi" hello</code>
[<code>hello</code>] -> Reply -> <code>/filter hi</code>

<b><u>Sแดแดแด าษชสแดแดส :</u></b>

<u>Fแดสแดแดแด</u>
<code>/stop "keyword"</code>
<u>Usแดษขแด</u>
<code>/stop "hi"</code>

<b><u>Vษชแดแดก าษชสแดแดสs :</u></b>
<code>/filters</code>
`,

	"GF": `
<b>Gสแดสแดส าษชสแดแดสs แดสแด แดแดษดแดแดส าษชสแดแดสs sแดแด แดแด สส สแดแด แดแดแดษชษดs แดสแดแด แดกแดสแด ษชษด แดสส แดสแดแดs. Tสแดส แดสแดแด ษชแดแด สแดแดแดsแด แดแดแด ษชแดs ษชษด แด แดแดsส แดแด แดsแด าแดสแดแดแด.</b>

<b><u>Sแดแดแด าษชสแดแดส :</u></b>

<u>Fแดสแดแดแด</u>
<code>/stop "keyword"</code>
<u>Usแดษขแด</u>
<code>/stop "et"</code>

<b><u>Vษชแดแดก าษชสแดแดสs :</u></b>
/gfilters
`,
	"CONNECT": `
<b>Cแดษดษดแดแดแดษชแดษดs แดสสแดแดก สแดแด แดแด แดแดษดแดษขแด สแดแดส ษขสแดแดแด สแดสแด ษชษด แดแด ษชษดsแดแดแดแด แดา sแดษดแดษชษดษข แดสแดsแด แดแดแดแดแดษดแดs แดแดสสษชแดสส ษชษด แดสแด ษขสแดแดแด โ โพ</b>

<b><u>Cแดษดษดแดแดแด :</u></b>
-> Fษชสsแด ษขแดแด สแดแดส ษขสแดแดแด's ษชแด สส sแดษดแดษชษดษข /id ษชษด สแดแดส ษขสแดแดแด
-> <code>/connect [group_id]</code>

<b><u>Dษชsแดแดษดษดแดแดแด :</u></b>
<code>/disconnect</code>
`,

	"BROADCAST": `
<b>The broadcast feature allows bot admins to broadcast a message to all of the bot's users.</b>

<I>Broadcasts are of two types :</i>
 ~ Full Broadcast - Broadcast to all of the bot users <code>/broadcast</code>
 ~ Concast - Broadcast to only users who are connected to a chat <code>/concast</code>
`,

	"HELP": `
<b>To know how to use my modules use the buttons below to get all my commands with usage examples !</b>
`,
}

var BUTTONS map[string][][]gotgbot.InlineKeyboardButton = map[string][][]gotgbot.InlineKeyboardButton{
	"START": {
		{{Text: "สแดสแดฉ", CallbackData: "edit(HELP)"},
			
		}, {
			{Text: "แดแดฉแดแดแดแด๊ฑ", Url: "t.me/Hell_Botz"}, {Text: "๊ฑแดแดฉแดฉแดสแด", Url: "t.me/Hellbotsupport"},
		},
		{{Text: "แดสแดแดแด", CallbackData: "edit(ABOUT)"}},
	},
	"ABOUT": {
		{
			{Text: "สแดแดแด", CallbackData: "edit(START)"},
			{Text: "๊ฑแดแดแด๊ฑ", CallbackData: "stats"},
		},
	},
	"STATS": {
		{
			{Text: "สแดแดแด", CallbackData: "edit(ABOUT)"},
			{Text: "สแด๊ฐสแด๊ฑส", CallbackData: "stats"},
		},
	},
	"HELP": {
		{{Text: "๊ฐษชสแดแดส", CallbackData: "edit(MF)"},
			{Text: "ษขสแดสแดส", CallbackData: "edit(GF)"},
		}, {
			{Text: "แดแดษดษดแดแดแด", CallbackData: "edit(CONNECT)"}, {Text: "สสแดแดแดแดแด๊ฑแด", CallbackData: "edit(BROADCAST)"},
		},
		{{Text: "โ", CallbackData: "edit(START)"}},
	},
}
