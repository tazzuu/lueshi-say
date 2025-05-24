# lueshi-say

Like [`cowsay`](https://en.wikipedia.org/wiki/Cowsay) but with our favorite Italian Plumber and his trusty dinosaur.

Run `lueshi-say` to hear LUEshi recite famous LUElinks lore and memes! Run the program repeatedly to hear what else he has to say :)

# Examples

<img width="1028" alt="Image" src="https://github.com/user-attachments/assets/ee65504f-9b5a-4777-8820-9c98836a5320" />

# Usage

## Command Line Mode

Run the program from the command line like this

```bash
./lueshi-say
```

LUEshi will output a random message from the LUElinks lore archives.

## Server Mode

Start the program in server mode like this

```bash
# defaults to port 4242
./lueshi-say -s
```

You can also start the Docker container version in server mode like this

```bash
docker run --rm -p 4242:4242 tazzuu/lueshi-say:latest -s
```

In both cases, in a separate terminal session, you can query the server to get a LUEshi response like this

```bash
curl localhost:4242

# or

wget localhost:4242 -q -O -
```

# Build & Download

Build from source with [Go 1.23+](https://go.dev/doc/install)

```bash
make build

./lueshi-say
```

Or download and run one of the pre-built executable binaries: https://github.com/tazzuu/lueshi-say/releases

Or run it from Docker

```bash
docker run --rm tazzuu/lueshi-say:latest
```

# Credits

Thanks to Patamon for the original version of the LUEshi ASCII art

- https://web.archive.org/web/20101023064231/http://gamefaqsarchive.com/index.php?pg=385

More ASCII arts are available here

- https://gfascii.art/

A history of LUEshi is here 

- https://web.archive.org/web/20101202145806/http://wikifaqs.net/index.php?title=LUEshi

Copy pasta taken from the LUE and LUElinks community lore archives
