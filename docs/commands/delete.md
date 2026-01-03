# pt delete

Appends a delete event to the event log for the last event or the specified event ID. This effectively hides an event from being processed while preserving the event log's forward immutability

  `pt delete [<event_id>]`

## Examples

```sh
$ pt note Hello world!

# delete the most recent event
$ pt delete

# delete the first event
$ pt delete 1

# a delete event itself can be deleted to undo
$ pt delete 3
```
