# checkpot
A tool for checking a hash:pass pot file for hashes from a user:hash file. Checkpot also provides the
ability to create a file of matches using the format user:pass. In addition, it will calculate the 
percentage of user credentials found in the pot. 

This is helpful when you have a file containing a list 
of users and hashes and a pot file (e.g. hashcat.pot) with only the hash and cracked password, but you 
want a list of users and passwords, and/or the percentage of users with weak passwords.

### Download
Operating system specific packages can be [downloaded from here](https://github.com/ryhanson/checkpot/releases).

### Usage
```text
  Usage with hashcat.pot: checkpot -u user_hashes.txt -p hashcat.pot -o user_passes.txt

  Options:
    -h, --help      Show usage and exit.
    -v              Show version and exit.
    -u              The line separated list of user:hash entries.
    -p              The line separated list of hash:pass entries.
    -o              Path to save file containing user:pass line entries.
```

### Building
Run the included bash script, which will build it for Linux, OS X, and Windows.
```bash
$ ./release
```
