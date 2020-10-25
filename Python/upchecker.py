import requests
import argparse
from pathlib import Path
import yaml
import socket


parser = argparse.ArgumentParser(description='Check if list of servers are available')
parser.add_argument("yaml_file", type=Path)
args = parser.parse_args()


with open(args.yaml_file) as yaml_file:
    targets = yaml.load(yaml_file)

for name in targets:
    target = targets[name]
    s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    print(f"Probing target {name} -- Host {target['host']} on TCP port {target['port']} ... ", end="")
    try:
        s.connect((target["host"], int(target["port"])))
        s.shutdown(2)
    except socket.error:
        print("FAIL")
    else:
        print("OK")



