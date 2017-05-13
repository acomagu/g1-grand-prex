# g1-grand-prix

This is my implementation of [G1 GRAND PRIX 五目並べ Hackathon](http://peatix.com/event/260374).

## Input/Output Rule

The client only inputs or outputs 2 numbers separated by one space, like `12 5`.

The inputs means the other client's placing, and the outputs means the your client's placing.

The client always outputs one placing after the other's inputs. If the client must place first, must got `-1 -1`.

## Usage

```bash
$ go run main.go
```

This will wait input from STDIN.

Put `-v` flag to print the game board to STDERR for debugging.
