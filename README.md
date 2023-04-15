# Wikiextract

[![builds.sr.ht status](https://builds.sr.ht/~jamesponddotco/wikiextract.svg)](https://builds.sr.ht/~jamesponddotco/wikiextract?)

**Wikiextract** is a word extractor for Wikipedia articles. It can
extract words bigger than 4 characters from a given Wikipedia page or
list of pages and save them to a file you can later use as the source
for generating [diceware passwords](https://en.wikipedia.org/wiki/Diceware).

## Installation

### From source

First install the dependencies:

- Go 1.20 or above.
- make.
- [scdoc](https://git.sr.ht/~sircmpwn/scdoc).

Then compile and install:

```bash
make
sudo make install
```

## Usage

See _wikiextract(1)_

## Contributing

Anyone can help make `wikiextract` better. Check out [the contribution
guidelines](https://git.sr.ht/~jamesponddotco/wikiextract/tree/master/item/CONTRIBUTING.md)
for more information.

## Resources

The following resources are available:

- [Support and general discussions](https://lists.sr.ht/~jamesponddotco/wikiextract-discuss).
- [Patches and development related questions](https://lists.sr.ht/~jamesponddotco/wikiextract-devel).
- [Instructions on how to prepare patches](https://git-send-email.io/).
- [Feature requests and bug reports](https://todo.sr.ht/~jamesponddotco/wikiextract).

---

Released under the [MIT License](LICENSE.md).
