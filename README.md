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
cop /source/path/ /destination/path --from <from_datetime> --to <to_datetime>
```

**Note: Dates must be in the following format: Year-Month-DayTHour:Minute
where T is the separator between date and time.

## Example

To move all files created between 01-01-2018 midnight & 01-02-2018 midnight
from /home/ubunutu to /home/ubuntu/myfiles do:

```shell
cop /home/ubuntu /home/ubuntu/myfiles --from 2018-01-01:00:00 --to 2018-02-01:00:00
```

## Contributing

1. Fork it ( https://github.com/[my-github-username]/cop/fork )
2. Create your feature branch (git checkout -b feature/my_new_feature)
3. Commit your changes (git commit -am 'Add some feature')
4. Push to the branch (git push origin my-new-feature)
5. Create a new Pull Request
