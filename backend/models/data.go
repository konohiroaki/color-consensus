package models

import "time"

/**
 * will be removed with real db.
 */

var Consensus = []*ColorConsensus{
	{Language: "en", Color: "red", Vote: 10, Colors: map[string]int{"ff0000": 10, "#ff007f": 3}},
	{Language: "en", Color: "green", Vote: 10, Colors: map[string]int{"00ff00": 10, "#00ff33": 3}},
	{Language: "ja", Color: "赤", Vote: 15, Colors: map[string]int{"ff0000": 15, "#ff007f": 5}},
}

var Votes = []*ColorVote{
	{Language: "en", Color: "red", User: "foo", Date: time.Now(), Colors: []string{"ff0000", "#ff007f"}},
	{Language: "en", Color: "red", User: "bar", Date: time.Now(), Colors: []string{"#ff0000", "#ff007f"}},
	{Language: "en", Color: "red", User: "baz", Date: time.Now(), Colors: []string{"#ff0000", "#ff007f"}},
	{Language: "en", Color: "red", User: "aaa", Date: time.Now(), Colors: []string{"#ff0000", "#ff007f"}},
}

var Users = []*User{
	{ID: "0da04f70-dc71-4674-b47b-365c3b0805c4", Nationality: "Japan", Gender: "Male", Birth: 1990, Date: time.Now()},
	{ID: "20af3406-8c7e-411a-851f-31732416fa83", Nationality: "Japan", Gender: "Male", Birth: 1991, Date: time.Now()},
}
