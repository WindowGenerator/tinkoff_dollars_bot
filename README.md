## Overview

Tinkoff Dollars Bot might help you find out about ATMs that carry your currency.

## Roadmap

- [x] - Write an MVP that should contain: 
* documentation
* normal binding with Makefile
* Commands: (help, start, get_atms (with filters))

- [ ] - Making CI work
- [ ] - Simplify the solution
- [ ] - Make gorutines for ATMs request
- [ ] - Get Cache for Cities
- [ ] - Make background task for have update cache
- [ ] - Make background notification task (mvp)
- [ ] - Make command for setup notification task
- [ ] - Make a normal interface for communication with the bot

## Documentation

#### Installation: 
`make -f Makefile install-deps`
`make -f Makefile local-build`

#### Running (local):
create .env in workdir and type your telegram bot token
`./_output/bin/linux/amd64/tinkoff_dollars_bot`

#### Contibuting (I dont know why upu want to contribute this):
`make -f Makefile install-deps`
`make -f Makefile install-instruments`

After edits:

`make -f Makefile local-build`
`make -f Makefile local-ci`