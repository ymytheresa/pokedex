# Pokedex CLI

This is a command-line interface (CLI) Pokedex application written in Go. It allows users to explore Pokémon in different location areas, catch Pokémon, and view their caught Pokémon.

## Building and Running

To build and run the Pokedex CLI, follow these steps:

1. Navigate to the root directory of the project.
2. Build the application using the following command:
   ```
   go build
   ```
3. Run the application:
   ```
   ./Pokedex
   ```

## Usage

Once you've started the Pokedex CLI, you'll see a prompt that looks like this:

```
Pokedex >
```

Here are the available commands and their usage:

### Help

To see a list of available commands, type:

```
Pokedex > help
```

This will display the following message:

```
Welcome to the Pokedex!
Usage:

explore : explore <location-area> to see list of pokemon of that location
catch   : catch <pokemon name>, pokemon might be caught or escaped
pokedex : show all the caught pokemon
help    : Displays a help message
exit    : Exit the Pokedex
map     : Show page of map
mapb    : Show last page of map
```

### Explore

To explore Pokémon in a specific location area:

```
Pokedex > explore <location-area>
```

Replace `<location-area>` with the name of the location you want to explore.

### Catch

To attempt to catch a Pokémon:

```
Pokedex > catch <pokemon name>
```

Replace `<pokemon name>` with the name of the Pokémon you want to catch. The Pokémon might be caught or escape.

Example:
```
Pokedex > catch mewtwo
Throwing a Pokeball at mewtwo...
mewtwo escaped!
Pokedex > catch dialga
Throwing a Pokeball at dialga...
dialga was caught!
```

### View Pokedex

To view all the Pokémon you've caught:

```
Pokedex > pokedex
```

This command will list all the Pokémon you've successfully caught.

Example:
```
Pokedex > pokedex
blissey
pikachu
arecus
dialga
```

### Map

To view the current page of the map:

```
Pokedex > map
```

To view the previous page of the map:

```
Pokedex > mapb
```

### Exit

To exit the Pokedex CLI:

```
Pokedex > exit
```

Enjoy exploring and catching Pokémon with your new Pokedex CLI!