import sys
from tools_header import make_header
from tools_types import Architecture, System
from tools_os import execute
import os


def make(system: System, architecture: Architecture, debug: bool = False):
    if system == System.ALL:
        make(System.LINUX, architecture, debug)
        make(System.DARWIN, architecture, debug)
        make(System.WINDOWS, architecture, debug)
        return
    
    if architecture == Architecture.ALL:
        make(system, Architecture.AMD64, debug)
        make(system, Architecture.ARM64, debug)
        make(system, Architecture.ARM, debug)
        if system != System.DARWIN:
            make(system, Architecture.I386, debug)
        return

    with open("VERSION", "r") as f:
        version = f.read().strip()
    sources = "lab.draklowell.net/routine-runtime/wrapper/"
    name = f"routine-runtime-{version}"

    if system == System.WINDOWS:
        compiler = "x86_64-w64-mingw32-gcc"
        extension = "dll"
    elif system == System.LINUX:
        compiler = "gcc"
        extension = "so"
    elif system == System.DARWIN:
        compiler = "gcc"
        extension = "so"

    target_name = f"{name}-{system.value}-{architecture.value}"

    goos = system.value
    goarch = architecture.value

    flags = ""
    if not debug:
        flags += '-ldflags="-s -w"'

    execute(f"CC={compiler} GOOS={goos} GOARCH={goarch} go build {flags} -buildmode=c-shared -o /build/{target_name}.{extension} {sources}")

    with open(f"/build/{target_name}.h", "r") as f:
        header = f.read()

    with open(f"wrapper/types.h", "r") as f:
        header_types = f.read()
    
    print(f"python:tools_header.make_header /build/{target_name}.h")
    with open(f"/build/{target_name}.h", "w") as f:
        f.write(make_header(header, header_types, system, architecture, debug))

def main():
    system = os.getenv("OS", "all")
    architecture = os.getenv("ARCH", "all")

    debug = os.getenv("DEBUG")
    if debug == None:
        debug = False
    else:
        debug = True

    make(System(system), Architecture(architecture), debug)

if __name__ == "__main__":
    try:
        main()
    except SystemExit as e:
        print("Exit with code:", e.code)
        sys.exit(min(e.code, 127))
