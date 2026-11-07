import json

def get_levels():
    try:
        with open("level_catalog/index.json", "r", encoding="utf-8") as file:
            levels = json.load(file)
        return levels
    except FileNotFoundError as e:
        print(f"Error loading levels: {e}")
        return None
    except json.JSONDecodeError as e:
        print(f"Error loading levels: {e}")
        return None

def load_level(file_path):
    # check that the file reading works and that the file is valid JSON and that the file exists
    try:
        with open(file_path, "r", encoding="utf-8") as file:
            level = json.load(file)
        return level
    except FileNotFoundError as e:
        print(f"Error loading level data: {e}")
        return None
    except json.JSONDecodeError as e:
        print(f"Error loading level data: {e}")
        return None

def validate_level(level):
    if level is None:
        print("Unable to load level data. See error messages above.")
        return False
    try:
        if "rooms" not in level:
            raise ValueError("Rooms key is required")
        if "start" not in level:
            raise ValueError("Start key is required")
        if "title" not in level:
            raise ValueError("Title key is required")
        # check that the start is a valid room id
        if level["start"] not in level["rooms"]:
            raise ValueError("Start key must be a valid room id")
        return True
    except ValueError as e:
        print(e)
        return False

def validate_room(room):
    try:
        if "name" not in room:
            raise ValueError("Name key is required")
        if "description" not in room:
            raise ValueError("Description key is required")
        if "exits" not in room:
            raise ValueError("Exits key is required")
        return True
    except ValueError as e:
        print(e)
        return False

def validate_exits(exits_obj, room_name):
    try:
        # check that the exits are a dictionary of strings - the key is the direction and the value is the room id
        if not isinstance(exits_obj, dict) or not all(isinstance(key, str) for key in exits_obj) or not all(isinstance(value, str) for value in exits_obj.values()):
            raise ValueError(f"Exits must be a dictionary of strings - the key is the direction and the value is the room id in room {room_name}")
        return True
    except ValueError as e:
        print(e)
        return False


# this is the code that will be executed when the script is run. it calls the functions defined above as part of validation.
levels = get_levels()
for level_id, level_path in levels.items():
    level = load_level("level_catalog/" + level_path)
    if not validate_level(level):
        print("level data is invalid. See error messages above.")
        exit(1)
    for room in level["rooms"].values():
        if not validate_room(room):
            print("Room data is invalid. See error messages above.")
            exit(1)
        if not validate_exits(room["exits"], room["name"]):
            print("Exit data is invalid. See error messages above.")
            exit(1)
print("Level data is valid")