# Yeah

Output a string repeatedly until killed.

Yet Another [`yes`](https://github.com/coreutils/coreutils/blob/master/src/yes.c) clone but in Golang.

## Usage

Just like [`yes`](<https://en.wikipedia.org/wiki/Yes_(Unix)>):

```bash
yeah
```

> This will print "y" until the process is canceled.

You can also pass a custom string to repeat:

```bash
yeah custom string
```

> This will print "custom string" until the process is canceled.

### Installation

```bash
go install github.com/UltiRequiem/yeah@latest
```

### License

This project is licensed under the [MIT License](./LICENSE.md).
