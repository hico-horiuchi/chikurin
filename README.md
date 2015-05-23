## chikurin v0.2.0

![status.png](https://raw.githubusercontent.com/hico-horiuchi/chikurin/master/status.png)
![clients.png](https://raw.githubusercontent.com/hico-horiuchi/chikurin/master/clients.png)

#### Requirements

  - [Golang](https://golang.org/) >= 1
  - [Sensu](http://sensuapp.org/) >= 0.13

#### Installation

    $ git clone git://github.com/hico-horiuchi/chikurin.git
    $ cd chikurin
    $ make gom
    $ make bindata
    $ sudo make install

#### Configuration

`~/.chikurin.json`

    {
      "datacenters": [
        {
          "name": "server-1",         // Required
          "host": "192.168.11.10",    // Required
          "port": 4567,               // Required
          "user": "sensu-1",          // Optional
          "password": "password"      // Optional
        },
        {
          "name": "server-2",
          "host": "192.168.11.20",
          "port": 4567
        }
      ],
      "show_datacenters": true,       // Optional
      "show_clients": true,           // Optional
      "port": 8000,                   // Optional
      "timeout": 3,                   // Optional
      "log": "/var/log/chikurin.log"  // Optional
    }

#### Usage

    Sensu status page by golang
    https://github.com/hico-horiuchi/chikurin
    
    Usage: 
      chikurin [flags]
      chikurin [command]
    
    Available Commands: 
      start       Start chikurin daemon
      stop        Stop chikurin daemon
      status      Show status of chikurin daemon
      version     Print and check version of chikurin
      help        Help about any command
    
    Flags:
      -h, --help=false: help for chikurin
    
    Use "chikurin help [command]" for more information about a command.

#### License

chikurin is released under the [MIT license](https://raw.githubusercontent.com/hico-horiuchi/chikurin/master/LICENSE).
