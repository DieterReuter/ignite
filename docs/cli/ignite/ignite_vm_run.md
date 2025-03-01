## ignite vm run

Create a new VM and start it

### Synopsis


Create and start a new VM immediately. The image (and kernel) is matched by
prefix based on its ID and name. This command accepts all flags used to
create and start a VM. The interactive flag (-i, --interactive) can be
specified to immediately attach to the started VM after creation.

Example usage:
	$ ignite run centos:7 \
		--interactive \
		--name my-vm \
		--cpus 2 \
		--ssh \
		--memory 2GB \
		--size 10G


```
ignite vm run <OCI image> [flags]
```

### Options

```
      --config string            Specify a path to a file with the API resources you want to pass
  -f, --copy-files strings       Copy files from the host to the created VM
      --cpus uint                VM vCPU count, 1 or even numbers between 1 and 32 (default 1)
  -d, --debug                    Debug mode, keep container after VM shutdown
  -h, --help                     help for run
  -i, --interactive              Attach to the VM after starting
      --kernel-args string       Set the command line for the kernel (default "console=ttyS0 reboot=k panic=1 pci=off ip=dhcp")
  -k, --kernel-image oci-image   Specify an OCI image containing the kernel at /boot/vmlinux and optionally, modules (default weaveworks/ignite-kernel:4.19.47)
      --memory size              Amount of RAM to allocate for the VM (default 512.0 MB)
  -n, --name string              Specify the name
      --net network-mode         Networking mode to use. Available options are: [cni docker-bridge] (default docker-bridge)
  -p, --ports strings            Map host ports to VM ports
  -s, --size size                VM filesystem size, for example 5GB or 2048MB (default 4.0 GB)
      --ssh[=<path>]             Enable SSH for the VM. If <path> is given, it will be imported as the public key. If just '--ssh' is specified, a new keypair will be generated. (default is unset, which disables SSH access to the VM)
```

### Options inherited from parent commands

```
      --log-level loglevel   Specify the loglevel for the program (default info)
  -q, --quiet                The quiet mode allows for machine-parsable output by printing only IDs
```

### SEE ALSO

* [ignite vm](ignite_vm.md)	 - Manage VMs

