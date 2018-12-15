# Title

> Bringing golang's sync.WaitGroup to the terminal!

`termwg` allows easy synchronisation of execution across multiple terminals.

## Install

```bash
go get -u github.com/Nekroze/termwg/...
```

## Usage

Synchronise execution of commands on multiple terminals, borrowing from the golang `sync.WorkGroup` abstraction.

```bash
termwg wait
```

The above command will block forever if we do nothing. Now in another terminal we may execute:

```bash
termwg done
```

Now both terminal's commands will exit successfully as the `default` work group counter was lowered to 0 by the `termwg done` command.

This provides opportunities to synchronise execution in multiple terminals easily with one liners.

Say we have our logs in one tmux pane and we build in another. I may need to do a rebuild and restart it along with my log viewer after it is complete.

So in my logviewer terminal I execute:

```bash
termwg wait build && tail -f *.log
```

Now over in my build terminal I can execute the following and go get a coffee as both will execute in the correct order.

```bash
make build start && termwg done build
```

## Contributing

PRs welcome!

## License

GPLv3 © 2018 Taylor Lawson
