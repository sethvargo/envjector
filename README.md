# Envjector

A small, single static binary for populating an environment from a file and
executing a child process with that environment.


## Usage

```
$ envjector -file /path/to/envfile -- myapp
```


## FAQ

**How is this different than [`env(1)`][env1]?**
<br>
In functionality, its quite similar. `env(1)` probably has more functionality
related to changing directories and signal handling. Unlike `env(1)`, envjector
is a single static binary with builds for Linux, Darwin, and Windows. Envjector
accepts an environment file, and uses execve(2) to call the process instead of
monitoring it as a child.


[env1]: http://man7.org/linux/man-pages/man1/env.1.html
