package bot

import (
	"../config"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

//Imports the config file, fmt, the API and strings.

var BotID string
var goBot *discordgo.Session

// declares vars

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)
	// starts the bot with the Token

	if err != nil {
		fmt.Println(error(err))
		return
	}

	// checks if there are any errors

	u, err := goBot.User("@me")
	// sets the user to @me

	if err != nil {
		fmt.Println(error(err))
	}
	// checks for errors again

	BotID = u.ID

	//goBot.AddHandler(messageHandler)
	goBot.AddHandler(halloFunc)
	goBot.AddHandler(vriendjesFunc)
	goBot.AddHandler(helpFunc)
	goBot.AddHandler(testFunc)
	goBot.AddHandler(weatherFunc)

	// adds handlers

	err = goBot.Open()

	// sets var

	if err != nil {

		fmt.Println(err.Error())
		return
	}
	// error check

	fmt.Println("Bot succesfully started.")
} // prints above line if the bot is actually running

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if strings.HasPrefix(m.Content, config.BotPrefix) {
		if m.Author.ID == BotID {
			return
		}
		if m.Content == "!ping" {
			_, _ = s.ChannelMessageSend(m.ChannelID, "Pong")
		}
	}
}

func vriendjesFunc(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, config.BotPrefix) {
		if m.Author.ID == BotID {
			return
		}
		if m.Content == "!vriend" {
			friendadd := s.RelationshipFriendRequestSend(m.Author.ID)     //sends friend request i think
			msg := fmt.Sprintf("Ik heb je als vriend ge-add.", friendadd) // prints this line while it mentions the user

			_, err := s.ChannelMessageSend(m.ChannelID, msg)
			if err != nil {
				fmt.Println(error(err))
			}
		}
	}
}
func halloFunc(s *discordgo.Session, m *discordgo.MessageCreate) { // creates a function halloFunc which uses the .Session and .MessageCreate structs
	if strings.HasPrefix(m.Content, config.BotPrefix) { // checks every msg in chan to check if it's for the bot
		if m.Author.ID == BotID { // checks if the author of the message is the bot itself, if so, it returns.
			return
		}
		if m.Content == "!hallo" { // checks if the content in the msg is !vriend
			mention := m.Author.Mention()           // does shit
			msg := fmt.Sprintf("Hallo %s", mention) // prints this line while it mentions the user

			_, err := s.ChannelMessageSend(m.ChannelID, msg)
			if err != nil {
				fmt.Println(error(err))
			}
		}
	}
}

func helpFunc(s *discordgo.Session, m *discordgo.MessageCreate) { // creates a function halloFunc which uses the .Session and .MessageCreate structs
	if strings.HasPrefix(m.Content, config.BotPrefix) { // checks every msg in chan to check if it's for the bot
		if m.Author.ID == BotID { // checks if the author of the message is the bot itself, if so, it returns.
			return
		}
		if m.Content == "!help" { // checks if the content in the msg is !help
			// does shit
			msg := fmt.Sprintf("```Discord bot written in Go, by Gnunix.```") // prints this line while it mentions the user

			_, err := s.ChannelMessageSend(m.ChannelID, msg)
			if err != nil {
				fmt.Println(error(err))
			}
		}
	}
}
func testFunc(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, config.BotPrefix) {
		if m.Author.ID == BotID {
			return
		}
		if m.Content == "!doesitwork" {

			msg := fmt.Sprintf("```Yes, it's currently working.```")

			_, err := s.ChannelMessageSend(m.ChannelID, msg)
			if err != nil {
				fmt.Println(error(err))
			}
		}
	}
}

func weatherFunc(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, config.BotPrefix) {
		if m.Author.ID == BotID {
			return
		}
		if m.Content == "!weather" {

			msg := fmt.Sprintf("```wip```")

			_, err := s.ChannelMessageSend(m.ChannelID, msg)
			if err != nil {
				fmt.Println(error(err))
			}
		}
	}
}
