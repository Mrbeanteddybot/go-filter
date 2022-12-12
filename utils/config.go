// (c) Jisin0

package utils

import "github.com/PaulSonOfLars/gotgbot/v2"

var TEXT map[string]string = map[string]string{

	"START": `
𝖧𝖾𝗒 %v 𝖭𝗂𝖼𝖾 𝗍𝗈 𝗆𝖾𝖾𝗍 𝗒𝗈𝗎 🙌
𝗂𝗆 𝖺𝗇 𝖺𝗐𝖾𝗌𝗈𝗆𝖾 𝖿𝗂𝗅𝗍𝖾𝗋 𝖻𝗈𝗍 𝗐𝗂𝗍𝗁 𝗀𝗅𝗈𝖻𝖺𝗅 𝖿𝗂𝗅𝗍𝖾𝗋 𝗌𝗎𝗉𝗉𝗈𝗋𝗍

𝖨 𝖼𝖺𝗇 𝗌𝖺𝗏𝖾 𝖺 𝖼𝗎𝗌𝗍𝗈𝗆 𝗋𝖾𝗉𝗅𝗒 𝖿𝗈𝗋 𝖺 𝗐𝗈𝗋𝖽 𝗂𝗇 𝖺𝗇𝗒 𝖼𝗁𝖺𝗍. 𝖢𝗁𝖾𝖼𝗄 𝗆𝗒 𝗁𝖾𝗅𝗉 𝗆𝖾𝗇𝗎 𝖿𝗈𝗋 𝗆𝗈𝗋𝖾 𝖽𝖾𝗍𝖺𝗂𝗅𝗌
	`,
	"ABOUT": `
<b>➥ 𝖫𝖺𝗇𝗀𝗎𝖺𝗀𝖾</b> : <a href='https://go.dev'>𝗚𝗼</a>
<b>➥ 𝖥𝗋𝖺𝗆𝖾𝗐𝗈𝗄</b> : <a href='github.com/PaulSonOfLars/gotgbot'>𝗚𝗼𝘁𝗴𝗯𝗼𝘁</a>
<b>➥ 𝖲𝖾𝗋𝗏𝖾𝗋</b> : <a href='heroku.com'>𝗛𝗲𝗿𝗼𝗸𝘂</a>
<b>➥ 𝖣𝖺𝗍𝖺𝖻𝖺𝗌𝖾</b> : <a href='mongodb.org'>𝗠𝗼𝗻𝗴𝗼𝗗𝗕</a>
<b>➥ 𝖣𝖾𝗏𝖾𝗅𝗈𝗉𝖾𝗋</b> : <a href='t.me/About_beantg'>𝗕𝗲𝗮𝗻</a>

	`,

	"MF": `
<b>Mᴀɴᴜᴀʟ ғɪʟᴛᴇʀs ᴀʟʟᴏᴡ ʏᴏᴜ ᴛᴏ sᴀᴠᴇ ᴄᴜsᴛᴏᴍ ғɪʟᴛᴇʀs ᴏᴛʜᴇʀ ᴛʜᴀɴ ᴛʜᴇ ᴀᴜᴛᴏᴍᴀᴛɪᴄ ᴏɴᴇs. Fɪʟᴛᴇʀs ᴄᴀɴ ʙᴇ ᴏғ ᴛᴇxᴛ/ᴘʜᴏᴛᴏ/ᴅᴏᴄᴜᴍᴇɴᴛ/ᴀᴜᴅɪᴏ/ᴀɴɪᴍᴀᴛɪᴏɴ/ᴠɪᴅᴇᴏ .</b>

<b><u>Nᴇᴡ ғɪʟᴛᴇʀ :</u></b>

<u>Fᴏʀᴍᴀᴛ</u>
<code>/filter "keyword" text</code> or
Rᴇᴘʟʏ ᴛᴏ ᴀ ᴍᴇssᴀɢᴇ -><code>/filter "keyword"</code>
<u>Usᴀɢᴇ</u>
<code>/filter "hi" hello</code>
[<code>hello</code>] -> Reply -> <code>/filter hi</code>

<b><u>Sᴛᴏᴘ ғɪʟᴛᴇʀ :</u></b>

<u>Fᴏʀᴍᴀᴛ</u>
<code>/stop "keyword"</code>
<u>Usᴀɢᴇ</u>
<code>/stop "hi"</code>

<b><u>Vɪᴇᴡ ғɪʟᴛᴇʀs :</u></b>
<code>/filters</code>
`,

	"GF": `
<b>Gʟᴏʙᴀʟ ғɪʟᴛᴇʀs ᴀʀᴇ ᴍᴀɴᴜᴀʟ ғɪʟᴛᴇʀs sᴀᴠᴇᴅ ʙʏ ʙᴏᴛ ᴀᴅᴍɪɴs ᴛʜᴀᴛ ᴡᴏʀᴋ ɪɴ ᴀʟʟ ᴄʜᴀᴛs. Tʜᴇʏ ᴘʀᴏᴠɪᴅᴇ ʟᴀᴛᴇsᴛ ᴍᴏᴠɪᴇs ɪɴ ᴀ ᴇᴀsʏ ᴛᴏ ᴜsᴇ ғᴏʀᴍᴀᴛ.</b>

<b><u>Sᴛᴏᴘ ғɪʟᴛᴇʀ :</u></b>

<u>Fᴏʀᴍᴀᴛ</u>
<code>/stop "keyword"</code>
<u>Usᴀɢᴇ</u>
<code>/stop "et"</code>

<b><u>Vɪᴇᴡ ғɪʟᴛᴇʀs :</u></b>
/gfilters
`,
	"CONNECT": `
<b>Cᴏɴɴᴇᴄᴛɪᴏɴs ᴀʟʟᴏᴡ ʏᴏᴜ ᴛᴏ ᴍᴀɴᴀɢᴇ ʏᴏᴜʀ ɢʀᴏᴜᴘ ʜᴇʀᴇ ɪɴ ᴘᴍ ɪɴsᴛᴇᴀᴅ ᴏғ sᴇɴᴅɪɴɢ ᴛʜᴏsᴇ ᴄᴏᴍᴍᴀɴᴅs ᴘᴜʙʟɪᴄʟʏ ɪɴ ᴛʜᴇ ɢʀᴏᴜᴘ ⠘⁾</b>

<b><u>Cᴏɴɴᴇᴄᴛ :</u></b>
-> Fɪʀsᴛ ɢᴇᴛ ʏᴏᴜʀ ɢʀᴏᴜᴘ's ɪᴅ ʙʏ sᴇɴᴅɪɴɢ /id ɪɴ ʏᴏᴜʀ ɢʀᴏᴜᴘ
-> <code>/connect [group_id]</code>

<b><u>Dɪsᴄᴏɴɴᴇᴄᴛ :</u></b>
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
		{{Text: "ʜᴇʟᴩ", CallbackData: "edit(HELP)"},
			
		}, {
			{Text: "ᴜᴩᴅᴀᴛᴇꜱ", Url: "t.me/Hell_Botz"}, {Text: "ꜱᴜᴩᴩᴏʀᴛ", Url: "t.me/Hellbotsupport"},
		},
		{{Text: "ᴀʙᴏᴜᴛ", CallbackData: "edit(ABOUT)"}},
	},
	"ABOUT": {
		{
			{Text: "ʜᴏᴍᴇ", CallbackData: "edit(START)"},
			{Text: "ꜱᴛᴀᴛꜱ", CallbackData: "stats"},
		},
	},
	"STATS": {
		{
			{Text: "ʙᴀᴄᴋ", CallbackData: "edit(ABOUT)"},
			{Text: "ʀᴇꜰʀᴇꜱʜ", CallbackData: "stats"},
		},
	},
	"HELP": {
		{{Text: "ꜰɪʟᴛᴇʀ", CallbackData: "edit(MF)"},
			{Text: "ɢʟᴏʙᴀʟ", CallbackData: "edit(GF)"},
		}, {
			{Text: "ᴄᴏɴɴᴇᴄᴛ", CallbackData: "edit(CONNECT)"}, {Text: "ʙʀᴏᴀᴅᴄᴀꜱᴛ", CallbackData: "edit(BROADCAST)"},
		},
		{{Text: "◁", CallbackData: "edit(START)"}},
	},
}
