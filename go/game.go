package main

type Game struct {
  currentLevelID string // the current level the player is on
  currentRoomID string // the current room the player is in
  inventory []string // items the player has collected
  roomData map[string]Room // data for each room
}

func NewGame(currentLevelID string, currentRoomID string, inventory []string, roomData map[string]Room) *Game {
	return &Game{
		currentLevelID: currentLevelID,
		currentRoomID: currentRoomID,
		inventory: inventory,
		roomData: roomData,
	}
}