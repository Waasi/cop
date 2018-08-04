# Cop

Cop is a cli for copying files from a directory to another filtered by dates

## Installation

#### MacOS

`git clone https://github.com/Waasi/cop.git`

`make install-macos`

#### Linux (Coming Soon)

## Usage

To move all the files from a source directory to a destination directory
filtered by the from datetime and to datetime type the following in the
terminal:

```shell
cop /source/path/ /destination/path --from hour:minute:second --to hour:minute:second
```

**Note: Times must be in the following format: hour:minute:second where `:` is the separator of units

## Example

To move all files created between midnight 00:00:00 and noon 12:00:00
from /home/ubunutu to /home/ubuntu/myfiles do:

```shell
cop /home/ubuntu /home/ubuntu/myfiles --from 00:00:00 --to 12:00:00
```

## Contributing

1. Fork it ( https://github.com/[my-github-username]/cop/fork )
2. Create your feature branch (git checkout -b feature/my_new_feature)
3. Commit your changes (git commit -am 'Add some feature')
4. Push to the branch (git push origin my-new-feature)
5. Create a new Pull Request
