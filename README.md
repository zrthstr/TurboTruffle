# TurboTruffle

Thin ergonomic wrapper around the trufflehog secret scanning tool.

## About

Given a GitHub PAT, TurboTruffle fetches all repos for a given org.
Using `git clone --mirror` and [Trufflehog](https://github.com/trufflesecurity/trufflehog) to exhaustively scan GitHub repositories, including all commits, branches, and commit messages.

While Trufflehog is a great tool, it can be somewhat unergonomic at times.
Trufflehog requires repositories to be named with a .git extension for processing.

A .html file is produced that links to the file with its findings on GitHub at the exact commit message.


## setup
```
git clone https://github.com/zrthstr/TurboTuffle
make build
```

## usage
see `Makefile` for details

edit `target` file ..
```
make run
```

then see output `.html` files in `results`
