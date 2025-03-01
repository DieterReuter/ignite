## ignite rmi

Remove VM base images

### Synopsis


Remove one or multiple VM base images. Images are matched by prefix based on
their ID and name. To remove multiple images, chain the matches separated by spaces.
The force flag (-f, --force) kills and removes any running VMs using the image.


```
ignite rmi <image> [flags]
```

### Options

```
  -f, --force   Force this operation. Warning, use of this mode may have unintended consequences.
  -h, --help    help for rmi
```

### Options inherited from parent commands

```
      --log-level loglevel   Specify the loglevel for the program (default info)
  -q, --quiet                The quiet mode allows for machine-parsable output by printing only IDs
```

### SEE ALSO

* [ignite](ignite.md)	 - ignite: easily run Firecracker VMs

