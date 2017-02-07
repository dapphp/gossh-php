## Name:

**GoSSH-php** - An SSH connection manager written in PHP

### Version:

**1.0.0**

### Author:

Drew Phillips <drew@drew-phillips.com>

### Requirements:

* PHP 5.2 or greater

## Description:

GoSSH-php is an SSH connection manager written in PHP.  It's purpose is to let
you easily define SSH connection shortcuts in a PHP config file so you can
connect to frequently used servers with minimal typing.

For example, once we've defined a connection, rather than having to type:

    ssh a.username@gw-1.remote.example-server.net -A -p 20222

With a connection shortcut, the above is as simple as:

    go work

As many connections as needed can be defined and managed from a single config,
each with their own parameters and options.

PHP is only used to parse options, manage connections, and build an SSH command.
A shell script is responsible for invoking PHP, and then SSH for connecting.
Once the connection is established, the only remaining process from calling `go`
is SSH.  Both PHP and the shell script terminate before connecting.

## Installation:

***INSTRUCTIONS HERE FOR DOWNLOADING AND INSTALLING***

If the program isn't installed to your default path, it's recommended to set up
a symlink, or alias.

One possibility is:

`ln -s /opt/gossh-php/go /usr/local/bin/go`

This creates a symlink from `/usr/local/bin/go` to our launcher in the install
directory.  Now, you can run `go` from any path to launch a connection.

Alternatively, you can create an alias in your `.bash_aliases` or `.bashrc`
file to define the command and alias:

`alias go=/opt/gossh-php/go`


## Host Configuration:

The purpose of this program is to give you a quick way to connect to commonly
used SSH hosts.  While not required, host information is kept in a config file
that use PHP syntax to define connections.

The default name of the configuration file is `.gohosts.cfg.php`.

GoSSH-php first tries to load it's configuration from your home directory, and
if not found will check in the same directory it is installed to.

Here's an example configuration file that describes the syntax:

    <?php

    // GoSSH-php configuration file

    return array(
        'hosts' => array(
            'work' => array('username', 'ssh.myworkplace.org', null, 1),
            'home' => array('me',       'myhouse.no-ip.org',   40200, 0),
        ),
    );

As you can see, the config file should *return* an array that contains an index
called hosts, that defines an array of connections.  Each host entry's array
key is the name of the shortcut used when connecting, and the connection itself
is an array where the first entry defines the username for connecting, the
second is the hostname to connect to, following by the port (or `null` for 22),
and then a true/false value to define whether or not to forward the
authentication agent connection (equivalent to `ssh -A`).

In the above example, we could type `go work` from the command line to execute:
`ssh -A root@ssh.myworkplace.org`.  The SSH authentication information is
forwarded so we can then SSH to other hosts from work using our local machine's
identity, rather than the identity on the remote server.

Running `go home` would execute `ssh me@myhouse.no-ip.org -p 40200`


## Using:

Once connections are defined in the config file, connecting is easy.  Just run
`go connection` where *connection* is the name of the connection to use.

Alternatively, you can use this to connect to any arbitrary host:

    go user@host.org

Properties given on the command line take precedence over connection properties.
For example, if we have defined a connection called `work`, with the username
`admin`, it's possible to connect as a different user by running
`go -u otheruser work`.

Type `go` to see advanced usage and list connections.

## TODO:

* Support for passing additional/unrecognized arguments directly through to the
SSH command.

## Copyright:

    Copyright (c) 2017 Drew Phillips
    All rights reserved.

    Redistribution and use in source and binary forms, with or without
    modification, are permitted provided that the following conditions are met:

    - Redistributions of source code must retain the above copyright notice,
      this list of conditions and the following disclaimer.
    - Redistributions in binary form must reproduce the above copyright notice,
      this list of conditions and the following disclaimer in the documentation
      and/or other materials provided with the distribution.

    THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
    AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
    IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
    ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
    LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
    CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
    SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
    INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
    CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
    ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
    POSSIBILITY OF SUCH DAMAGE.
