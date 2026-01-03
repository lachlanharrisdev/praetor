# pt capture

Capture tool output from STDIN or from a filepath. If Praetor officially supports the tool and is recognised, the tool will be parsed through its corresponding module *(NOT IMPLEMENTED YET)*

  `pt capture [options] <filepath|->`

## Examples

```sh
# capture output from STDIN
$ nmap -p- 127.0.0.1 | pt capture`

# capture output from a filepath
$ pt capture nmap_results.txt
```
