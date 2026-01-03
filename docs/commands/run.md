# pt run

Run a command, optionally within an isolated sandbox (similar to a container), and record it to the event log.

  `pt run [--sandbox -s] <command> [args...]`

## Examples

```sh
# run a command normally
$ pt run echo "Hello, World!"

# run a command inside of a sandbox
# NOTE: the sandbox has the current working directory cloned inside for scripts to access
$ pt run -s python3 dodgy_script.py wordlist.txt
```
