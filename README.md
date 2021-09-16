
test (memo) case for two problems
* https://github.com/firebase/firebase-tools/issues/3556
* to write metadata and object data at the same time does not works

```bash
$ go run main.go
file data: "123\r\n"
2021/09/16 23:13:14 failed to get the attrs: storage: object doesn't exist
exit status 1
```

emulator startup log

```log
⚠  emulators: You are not currently authenticated so some features may not work correctly. Please run firebase login to authenticate the CLI.
i  emulators: Starting emulators: storage
i  storage: downloading cloud-storage-rules-runtime-v1.0.1.jar...

i  ui: downloading ui-v1.6.3.zip...

i  ui: Emulator UI logging to ui-debug.log

┌─────────────────────────────────────────────────────────────┐
│ ✔  All emulators ready! It is now safe to connect your app. │
│ i  View Emulator UI at http://0.0.0.0:4000                  │
└─────────────────────────────────────────────────────────────┘

┌──────────┬──────────────┬─────────────────────────────┐
│ Emulator │ Host:Port    │ View in Emulator UI         │
├──────────┼──────────────┼─────────────────────────────┤
│ Storage  │ 0.0.0.0:9199 │ http://0.0.0.0:4000/storage │
└──────────┴──────────────┴─────────────────────────────┘
  Emulator Hub running at localhost:4400
  Other reserved ports: 4500

Issues? Report them at https://github.com/firebase/firebase-tools/issues and attach the *-debug.log files.
```
