
# Install

```
go get -u github.com/madogiwa/trigger
```

# Usage

```
trigger exec [file or directory] -- [command...]
```

## ex) wc -l

Start the `trigger` command for `docs` directory.

```console
$ trigger exec docs -- wc -l
watching for changes in docs ...
```

Add a line to `doc1.txt` in `docs` directory.

```
$ echo "line" >> /Users/Shared/docs/doc1.txt
```

Then `wc -l` is executed.

```
2019/11/17 16:26:16 exec: wc -l /Users/Shared/docs/doc1.txt
       1 /Users/Shared/docs/doc1.txt
```

Press `Ctrl-c` to shutdown.

```
^Cshutdown... wc docs
```
