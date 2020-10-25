<?php

function print_usage($custom = null) {
    global $argv;
    echo 'Usage: ' . $argv[0] . ' <Targets> <Scope>' . PHP_EOL . PHP_EOL;
    if($custom != null) {
        echo $custom . PHP_EOL . PHP_EOL;
    }
    echo 'Targets: must be a file formatted in YAML scheme' . PHP_EOL;
    echo 'Scope: can be a specific (list) of targets to be probed. Use the comma as a delimiter. Optional.' . PHP_EOL;
    exit(1);
}
