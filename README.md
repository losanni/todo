# Todo

A simple cli todo app written in go

## Todos

Todos saved to ```~/.local/share/todo/```

Each file is a standard text file. The name of the file is the "thing to do" and
the contents of the file are more info about the "thing to do".

## Support platforms

- Anything Unix supported by [Go](https://golang.org)

## Install

```
git clone https://github.com/losanni/todo
cd todo
make install
```

## Docs
```
Use the following comands. Wrap "todonames" and "texttoadd" in quotes

List todos:
	todo

Add todo:
	todo add todoname

Read todo:
	todo read todoname

Update todo description:
	todo edit todoname texttoadd

Remove todo:
	todo remove todoname

```
