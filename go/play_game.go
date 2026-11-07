package main
import (
  "fmt"
  "encoding/json"
  "os"
  "bufio"
  "strings"
  "path/filepath"
)

type Room struct {
  Name string `json:"name"`
  Description string `json:"description"`
  Exits map[string]string `json:"exits"`
}

type Level struct {
  Rooms map[string]Room `json:"rooms"`
  Start string `json:"start"`
  Title string `json:"title"`
}

const path_to_levels_index = "../level_catalog/index.json";
const path_to_levels = "../level_catalog/";

func getLevelsIndex() (map[string]string, error) {
	levels_index_data, err := os.ReadFile(path_to_levels_index);
	if err != nil {
		fmt.Println("Error reading levels index data:", err)
		return map[string]string{}, err
	}
	var levels_index map[string]string
	if err := json.Unmarshal(levels_index_data, &levels_index); err != nil {
		fmt.Println("Error unmarshalling levels index data:", err)
		return map[string]string{}, err
	}
	return levels_index, err
}

func getLevelData(level_id string) (Level, error) {
	// get the level data from the file mapped to the level_id in the levels index
	levels_index, err := getLevelsIndex()
	if err != nil {
		fmt.Println("Error getting levels index:", err)
		return Level{}, err
	} else {
		// load and return the level data from this successfully loaded level
		levelPath := filepath.Join(filepath.Dir(path_to_levels), levels_index[level_id])
		level_data, err := os.ReadFile(levelPath);
		if err != nil {
			fmt.Println("Error reading level data:", err)
			return Level{}, err
		}

		var level Level
		if err := json.Unmarshal(level_data, &level); err != nil {
			fmt.Println("Error unmarshalling level data:", err)
			return Level{}, err
		}
		return level, nil
	}
}

func main() {
	// load the level data from the starting level id
	// if level data is not set, set it 'level_1'
	startingLevelID := "level_1"
    if len(os.Args) > 1 {
        startingLevelID = os.Args[1]
    }

    // Use startingLevelID to load the level.
	level_data, err := getLevelData(startingLevelID)
	if err != nil {
		return
	}

	current_room := level_data.Rooms[level_data.Start]
	// print the current room description and ask what the user wants to do
	fmt.Println(current_room.Description)	
	fmt.Println("What do you want to do?")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			return
		}

		user_input := scanner.Text() // Retrieves the string value

		switch  {
		case user_input == "list exits":
			// if user types "list exits", print the exits
			for direction, room_id := range current_room.Exits {
				fmt.Println(direction, "->", room_id)
			}
		case user_input == "quit":
			// if user types "quit", quit the program
			return
		case user_input == "help":
			// if user types "help", print the help message
			fmt.Println("Available commands:")
			fmt.Println("list exits - list the exits in the current room")
			fmt.Println("quit - quit the program")
			fmt.Println("help - print this help message")
			fmt.Println("move <direction> - move to the next room in the given direction")
		case strings.HasPrefix(user_input, "move "):
			direction := strings.TrimPrefix(user_input, "move ")
			nextRoomID, ok := current_room.Exits[direction]
			if ok {
				current_room = level_data.Rooms[nextRoomID]
				fmt.Println(current_room.Description)
			} else {
				fmt.Println("Invalid direction. Type 'help' for a list of commands.")
			}
		default:
			// if the user types an invalid command, print the invalid command message
			fmt.Println("Invalid command. Type 'help' for a list of commands.")
		}

		// if user types an object name, print the object description (future feature)
		
	}

	return
}