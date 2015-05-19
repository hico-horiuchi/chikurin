## chikurin v0.1.0

#### Requirements

  - [Golang](https://golang.org/) >= 1
  - [Sensu](http://sensuapp.org/) >= 0.13

#### Installation

    $ git clone git://github.com/hico-horiuchi/chikurin.git
    $ cd ohgi
    $ make gom
    $ sudo make install

#### Configuration

`~/.chikurin.json`

    {
      "datacenters": [
        {
          "name": "server-1",       // Required
          "host": "192.168.11.10",  // Required
          "port": 4567,             // Required
          "user": "sensu-1",        // Optional
          "password": "password"    // Optional
        },
        {
          "name": "server-2",
          "host": "192.168.11.20",
          "port": 4567
        }
      ],
      "timeout": 3                  // Optional
    }

#### Usage

#### License

chikurin is released under the [MIT license](https://raw.githubusercontent.com/hico-horiuchi/chikurin/master/LICENSE).
