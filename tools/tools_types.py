from enum import Enum


class System(Enum):
    LINUX = "linux"
    DARWIN = "darwin"
    WINDOWS = "windows"

    ALL = "all"

class Architecture(Enum):
    AMD64 = "amd64"
    I386 = "i386"
    ARM64 = "arm64"
    ARM = "arm"

    ALL = "all"
