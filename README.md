# next-semantic-version-cli
generate next semantic version string from given string.

## install
go
```bash
go install github.com/mkusaka/next-semantic-version-cli/cmd/nsv@latest
```

shell
```
bash <(curl https://raw.githubusercontent.com/mkusaka/next-semantic-version-cli/main/scripts/install.bash)
./nsv --version
```

## usage
```
❯ nvc --help   
Usage of nsv:
  -type string
        bump target string. (default "patch")
  -version
        print current version
```

## example
```bash
❯ echo "v$(nsv "v0.0.1")"
v0.0.2

❯ echo "v$(nsv -type major "v0.0.1")"
v1.0.0

❯ echo "v$(nsv -type minor "v0.0.1")"
v0.1.0

❯ echo "v0.0.1" | xargs nsv -type minor
0.1.0
```
