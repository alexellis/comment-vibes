package emoji

// emoticonMap is a map of emoji aliases to emoticon counterparts.
var emoticonMap = map[string][]string{
	"angry":                        {">:(", ">:-("},
	"anguished":                    {"D:"},
	"broken_heart":                 {"</3"},
	"confused":                     {":/", ":-/", `:\`, `:-\`},
	"disappointed":                 {":(", "):", ":-("},
	"heart":                        {"<3"},
	"kiss":                         {":*", ":-*"},
	"laughing":                     {":>", ":->"},
	"monkey_face":                  {":o)"},
	"neutral_face":                 {":|"},
	"open_mouth":                   {":o", ":O", ":-o", ":-O"},
	"slightly_smiling_face":        {":)", "(:", ":-)"},
	"smile":                        {":D", ":-D"},
	"smiley":                       {"=)", "=-)"},
	"stuck_out_tongue":             {":p", ":P", ":-p", ":-P", ":b", ":-b"},
	"stuck_out_tongue_winking_eye": {";p", ";P", ";-p", ";-P", ";b", ";-b"},
	"sunglasses":                   {"8)"},
	"wink":                         {";)", ";-)"},
}
