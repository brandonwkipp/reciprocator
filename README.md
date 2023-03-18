# reciprocator

The `reciprocator` is a tool that can quickly reciprocate the notes within a `midi` file with respect a user-defined tonic in accordance with the concept of "Negative Harmony". This concept was first pioneered by twentieth century Musicologist [Ernst Levy](https://en.wikipedia.org/wiki/Ernst_Levy) and explored in his book, *A Theory of Harmony*.

## TODO
- Update build process
- See if I can rebuild in C (libsmf)?

### Build process
You will need to clone this repo down:
```
git clone https://github.com/brandonwkipp/reciprocator
cd reciprocator

# Build for Linux
make build-linux

# Build for MacOS
make build-darwin

# Build for Windows
make build-windows

cp ./bin/reciprocator ~/bin/reciprocator
```
