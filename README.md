# discord-bot-golang

- install golang ("sudo snap install go --classic")
- git clone ("git clone https://github.com/tmk11/discord-bot-golang.git")
- "cd discord-bot-golang"
- "go mod tidy"
- set variable Environment for discord token ("export DISCORD_TOKEN='your-discord-token'") 
- "go run main.go"

Discord Bot command: "!bot ...", "!bot weather [location]"


OR


- install Docker
- get my image ("sudo docker pull mkong4/golang-discord-bot")
- run container ("sudo docker run -d -e DISCORD_TOKEN='your-discord-token' --name golang-discord-bot-container mkong4/golang-discord-bot")
