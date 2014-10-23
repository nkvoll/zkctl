zkctl
=====

`zkctl` is a command line client for [Apache ZooKeeper](http://zookeeper.apache.org/).

It's inspired by tools like [etcdctl](https://github.com/coreos/etcdctl/) and the interest of having an easy-to-install dependency-free client to access ZooKeeper.


Building
--------

Uses [godep](https://github.com/tools/godep).


Supported operations
--------------------

``ls``, ``set`` and ``get`` are currently supported. Watches and executing commands on changes are on the planning board. Input on how it should work is appreciated, just open a [issue](issues) for open discussion.


Usage
-----

    NAME:
       zkctl - A new cli application

    USAGE:
       zkctl [global options] command [command options] [arguments...]

    VERSION:
       0.0.1

    COMMANDS:
       ls		list a directory
       get		get the contents of a node
       set		set the contents of a node
       help, h	Shows a list of commands or help for one command

    GLOBAL OPTIONS:
       --zookeeper, -z 	a ZooKeeper connection string [$ZOOKEEPER_SERVERS]
       --format, -f 	output format [$ZKCTL_OUTPUT_FORMAT]
       --help, -h		show help
       --version, -v	print the version



License
-------

Apache Licensed
