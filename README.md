# gompose

Gompose is a configurable process management application for local development.

Gompose is inspired from following works:

* https://github.com/mattn/goreman
* https://github.com/kimmobrunfeldt/concurrently
* https://github.com/docker/compose

## Installation

### Requirements
* Go 1.9+

`$ go get github.com/keekun/gompose`

## Usage

```
$ touch .gompose.yaml
```

edit `.gompose.yaml` and add processes settings.

```yaml
processes:
  job1:
    name: Say-Yes
    command: "echo '{\"test\": 12344}' | jq -c ."
    spawn: ["/bin/bash", "-c"]
    format:
      fgcolor: "cyan" 
      bgcolor: "black"
      header: "[{{.Proc.Name}}|{{.Now.Format \"15:04:05\"}}] "
  job2:
    name: JSON
    command: "echo '{\"json\": 777}' | jq -c ."
    format:
      fgcolor: "red"
      bgcolor: "black"
```

run

```
$ gompose
```

## Features

- [x] Spawn processes according to `.gompose.yaml` file with colorized log.
- [ ] Cache logs into [BoltDB](https://github.com/boltdb/bolt), index logs locally with [bleve](http://www.blevesearch.com/).
- [ ] Add interactive cli mode and allow user to filter/search logs.
- [ ] Add cli tool to start/stop/restart process.
- [ ] More options, like environment variables, signal trapping, etc.
- [ ] Support docker-compose-like options for docker.
- [ ] Add HTTP UI.
- [ ] Add tests.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
