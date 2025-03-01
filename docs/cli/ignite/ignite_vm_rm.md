## ignite vm rm

Remove VMs

### Synopsis


Remove one or multiple VMs. The VMs are matched by prefix based
on their ID and name. To remove multiple VMs, chain the matches
separated by spaces. The force flag (-f, --force) kills running
VMs before removal instead of throwing an error.


```
ignite vm rm <vm>... [flags]
```

### Options

```
  -f, --force   Force this operation. Warning, use of this mode may have unintended consequences.
  -h, --help    help for rm
```

### Options inherited from parent commands

```
      --log-level loglevel   Specify the loglevel for the program (default info)
  -q, --quiet                The quiet mode allows for machine-parsable output by printing only IDs
```

### SEE ALSO

* [ignite vm](ignite_vm.md)	 - Manage VMs

