#!/bin/bash

###############################################################################
#
#  Simple SSH launcher with shortcuts
#  See go.php for usage and to add shortcuts
#
#  This bash script is used for efficiency, using passthru() in PHP leaves the
#  PHP process running until the shell session is terminated.  This script uses
#  PHP to build the command, and then calls exec() from the shell to spawn a
#  new process and terminate.
#
###############################################################################

# get directory of this program
dir=$( realpath "${BASH_SOURCE[0]}" )
cwd=$( dirname "$dir" )

# construct php command from our args
cmd=$cwd"/go.php $@"

# execute php script and get output
res=$($cmd)
ret=$?

if [ $ret -eq 0 ] ; then
    # successful, exec the output of php
    echo "$res"
    exec $res
    exit $?
else
    # php returned non zero, error or show usage
    echo "$res"
    echo
    exit $ret
fi
