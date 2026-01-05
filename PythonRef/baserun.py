import os
import subprocess

cProberCache=os.path.exists("./.cache")
if not cProberCache:
    os.makedirs("./.cache")
if os.path.exists("./.cache/FunctionInfo.db"):
    os.remove("./.cache/FunctionInfo.db")
if os.path.exists("./.cache/PacketInfo.db"):
    os.remove("./.cache/PacketInfo.db")
# bpftool -j btf dump file /sys/kernel/btf/vmlinux > ./btf.json
subprocess.run(["rm","-f","./.cache/btf.json"])
fo=open("./.cache/btf.json","x")
subprocess.run(["bpftool","-j","btf","dump","file","/sys/kernel/btf/vmlinux"],stdout=fo)