// Copyright 2018, Goomba project Authors. All rights reserved.
//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with this
// work for additional information regarding copyright ownership.  The ASF
// licenses this file to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.  See the
// License for the specific language governing permissions and limitations
// under the License.

package password_generator

var (
	// // ADJECTIVES ...
	// ADJECTIVES = []string{"autumn", "hidden", "bitter", "misty", "silent", "empty", "dry", "dark", "summer",
	// 	"icy", "delicate", "quiet", "white", "cool", "spring", "winter", "patient",
	// 	"twilight", "dawn", "crimson", "wispy", "weathered", "blue", "billowing",
	// 	"broken", "cold", "damp", "falling", "frosty", "green", "long", "late", "lingering",
	// 	"bold", "little", "morning", "muddy", "old", "red", "rough", "still", "small",
	// 	"sparkling", "throbbing", "shy", "wandering", "withered", "wild", "black",
	// 	"young", "holy", "solitary", "fragrant", "aged", "snowy", "proud", "floral",
	// 	"restless", "divine", "polished", "ancient", "purple", "lively", "nameless"}

	// ADJECTIVES ...
	ADJECTIVES = []string{"Sunny", "Rainy", "Dry", "Dirty", "Silent", "Empty", "Dry", "Dark", "Cloudy",
		"Icy", "Foggy", "Quiet", "Clear", "Cool", "Windy", "Wet", "Still", "Pretty", "Ugly",
		"Old", "New", "Hot", "Cold", "Young", "Big", "Small", "Thick", "Thin", "Full", "Good",
		"Bad", "Best", "Black", "Early", "Late", "Free", "Great", "High", "Hard", "Human", "Large",
		"Little", "Only", "Other", "Real", "Fake", "Right", "Left", "True", "Red", "Blue", "Black",
		"Strong", "Weak", "Whole", "Short", "Long", "Next", "Few", "Same", "calm", "Cute", "Fair",
		"Fine", "Tiny", "Huge", "Happy", "Sad", "Hurt", "Kind", "Lazy", "Lucky", "Open", "Poor",
		"Shy", "Silly", "Rich", "Sleepy", "Stupid", "Super", "Tasty", "Tired", "Wild", "Wrong",
		"Brave", "Fat", "Wide", "Better", "Loud", "Bold", "Little", "Rough", "Snowy", "Proud"}

	// // NOUNS ...
	// NOUNS = []string{"waterfall", "river", "breeze", "moon", "rain", "wind", "sea", "morning",
	// 	"snow", "lake", "sunset", "pine", "shadow", "leaf", "dawn", "glitter", "forest",
	// 	"hill", "cloud", "meadow", "sun", "glade", "bird", "brook", "butterfly",
	// 	"bush", "dew", "dust", "field", "fire", "flower", "firefly", "feather", "grass",
	// 	"haze", "mountain", "night", "pond", "darkness", "snowflake", "silence",
	// 	"sound", "sky", "shape", "surf", "thunder", "violet", "water", "wildflower",
	// 	"wave", "water", "resonance", "sun", "wood", "dream", "cherry", "tree", "fog",
	// 	"frost", "voice", "paper", "frog", "smoke", "star"}
	// NOUNS ...
	NOUNS = []string{"book", "river", "apple", "lemon", "cherry", "banana", "peach", "orange",
		"pear", "kiwi", "mango", "plum", "fig", "papaya", "frog", "toad", "spider",
		"crow", "dove", "eagle", "duck", "goose", "gull", "jay", "owl", "parrot", "robin",
		"stork", "swan", "turkey", "eel", "ray", "shark", "trout", "shell", "ant", "bee",
		"flea", "fly", "bat", "bear", "camel", "deer", "fox", "goat", "hare", "lion", "tiger",
		"cat", "dog", "otter", "panda", "pony", "mouse", "rat", "whale", "zebra", "snail", "snake",
		"turtle", "carrot", "onion", "bean", "yam", "pumpkin", "summer", "winter", "fall",
		"tree", "hill", "storm", "yam", "pumpkin", "summer", "winter", "fall",
		"bush", "dew", "dust", "field", "fire", "flower", "firefly", "feather", "grass",
		"haze", "mountain", "night", "pond", "sea", "snowflake", "wind",
		"sound", "sky", "shape", "surf", "thunder", "violet", "water", "fruit",
		"wave", "water", "snow", "sun", "wood", "dream", "cherry", "tree", "fog",
		"frost", "voice", "paper", "frog", "smoke", "star"}
)
