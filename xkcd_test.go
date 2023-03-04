package main_test

import (
	"strings"
	"testing"
	"fmt"

	"github.com/namikaze-dev/xkcd"
)

const comicsJSON = `
[
	{
		"month": "4",
		"num": 571,
		"link": "",
		"year": "2009",
		"news": "",
		"safe_title": "Can't Sleep",
		"transcript": "[[Someone is in bed, presumably trying to sleep. The top of each panel is a thought bubble showing sheep leaping over a fence.]]\n1 ... 2 ...\n<<baaa>>\n[[Two sheep are jumping from left to right.]]\n\n... 1,306 ... 1,307 ...\n<<baaa>>\n[[Two sheep are jumping from left to right. The would-be sleeper is holding his pillow.]]\n\n... 32,767 ... -32,768 ...\n<<baaa>> <<baaa>> <<baaa>> <<baaa>> <<baaa>>\n[[A whole flock of sheep is jumping over the fence from right to left. The would-be sleeper is sitting up.]]\nSleeper: ?\n\n... -32,767 ... -32,766 ...\n<<baaa>>\n[[Two sheep are jumping from left to right. The would-be sleeper is holding his pillow over his head.]]\n\n{{Title text: If androids someday DO dream of electric sheep, don't forget to declare sheepCount as a long int.}}",
		"alt": "If androids someday DO dream of electric sheep, don't forget to declare sheepCount as a long int.",
		"img": "https://imgs.xkcd.com/comics/cant_sleep.png",
		"title": "Can't Sleep",
		"day": "20"
	},
	{
		"month": "11",
		"num": 500,
		"link": "",
		"year": "2008",
		"news": "",
		"safe_title": "Election",
		"transcript": "[[Character sits at his computer desk, staring at his computer.]]\nIt's over.\nAfter twenty months it's finally over.\nI don't have to be an election junkie anymore.\n[[Closeup of character's face and screen.]]\nI don't have to care about opinion polls, exit polls, margins of error, attack ads, game-changers, tracking polls, swing states, swing votes, the Bradley effect, or <name> the <occupation>.\nI'm free.\n[[Character staring at his computer screen, full shot.]]\n[[Character types on his computer.]] <<Tap Tap>>\n[[On screen]]Google  \"2012 polling statistics\"\n{{Title text: \"Someday I'll be rich enough to hire Nate Silver to help make all my life decisions.  'Should I sleep with her?'  'Well, I'm showing a 35% chance it will end badly.' \"}}",
		"alt": "Someday I'll be rich enough to hire Nate Silver to help make all my life decisions.  'Should I sleep with her?'  'Well, I'm showing a 35% chance it will end badly.'",
		"img": "https://imgs.xkcd.com/comics/election.png",
		"title": "Election",
		"day": "5"
	}
]`

func TestSearch(t *testing.T) {
	// check invalid json source
	comicsRd := strings.NewReader(".")
	_, err := main.NewSearcher(comicsRd)
	if err == nil {
		t.Fatal("create new searcher error expected, got nil")
	}

	// check valid json source
	comicsRd = strings.NewReader(comicsJSON)
	searcher, err := main.NewSearcher(comicsRd)
	if err != nil {
		t.Fatalf("create new searcher error: %v", err)
	}

	// check comics count
	got := searcher.Total()
	want := 2
	if want != got {
		t.Errorf("got %v, want %v", got, want)
	}

	// check single queries
	var cases = []struct {
		in   string
		want int
	}{
		{
			"Decisions",
			1,
		},
		{
			"SLEEP",
			2,
		},
		{
			"bakery",
			0,
		},
	}

	for _, c := range cases {
		matched := searcher.Search(c.in)
		got := len(matched)
		if c.want != got {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	// check multiple queries
	var cases2 = []struct {
		in   []string
		want int
	}{
		{
			[]string{"cargo", "sheep"},
			1,
		},
		{
			[]string{"Day", "SHEEPCOUNt"},
			2,
		},
		{
			[]string{"edgelord", "avarice", "pun"},
			0,
		},
	}

	for _, c := range cases2 {
		matched := searcher.Search(c.in...)
		got := len(matched)
		if c.want != got {
			for _, m := range matched {
				fmt.Println(m.Title)
			}
			t.Errorf("got %v, want %v", got, c.want)
		}
	}
}
