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

    platform = f"{system.value}-{architecture.value}"
    environ = {}
    if system == System.WINDOWS:
        if architecture == Architecture.AMD64:
            platform = "win64"
            compiler = "x86_64-w64-mingw32-gcc"
        else:
            print(f"Unsupported platform: {system.value}, {architecture.value}")
            return
        extension = "dll"
    elif system == System.LINUX:
        if architecture == Architecture.AMD64:
            compiler = "gcc"
        elif architecture == Architecture.ARM64:
            compiler = "aarch64-linux-gnu-gcc"
        elif architecture == Architecture.ARM:
            compiler = "arm-linux-gnueabi-gcc"
        else:
            print(f"Unsupported platform: {system.value}, {architecture.value}")
            return
        extension = "so"
    else:
        print(f"Unsupported platform: {system.value}, {architecture.value}")
        return

    target_name = f"{name}-{platform}"

    goos = system.value
    if architecture == Architecture.I386:
        goarch = "386"
    else:
        goarch = architecture.value

    flags = []
    if not debug:
        flags.append('-ldflags=-s -w')

    execute(
        ["go", "build", *flags, "-buildmode=c-shared", "-o", f"/build/{target_name}.{extension}", sources],
        {"CC": compiler, "GOOS": goos, "GOARCH": goarch, "CGO_ENABLED": "1", **environ}
    )

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
