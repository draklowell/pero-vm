import os

def execute(cmd: str):
    print(cmd)
    ret = os.system(cmd)
    if ret != 0:
        print("Not null exit code")
        exit(ret)
