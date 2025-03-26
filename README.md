# TurboTuffle
thin ergonomic wrapper around trufflehog secret scanning tool

using git clone --mirror to scan github repos to the fullest, all commit, all branches, all commit messages etc etc
somhow turfflehog will only process them if they are called .git ..


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

or
```
vim target && make build run
```

then see output `.html` files in `results`
