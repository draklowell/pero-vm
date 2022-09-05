import os
import subprocess
import sys
from time import sleep
from typing import Dict, List

def execute(cmd: List[str], environ: Dict[str, str]):
    print(" ".join(cmd))

    env_text = []
    for name, value in environ.items():
        env_text.append(f"{name}={value}")
    env_text = " ".join(env_text)
    print(f"[ENV: {env_text}]")

    sys.stdout.flush()
    sleep(0.1)
    env = os.environ.copy()
    env.update(environ)
    process = subprocess.Popen(cmd, stdout=sys.stdout, stderr=sys.stderr, env=env)
    retcode = process.wait()
    if retcode != 0:
        print("Not null exit code")
        exit(retcode)
