from tools_types import Architecture, System


def make_header(header: str, types: str, system: System, architecture: Architecture, debug: bool) -> str:
    result = ""

    if system == System.WINDOWS:
        result += "#define DllExport __declspec(dllexport)\n"

    result += '#ifdef __cplusplus\nextern "C" {\n#endif\n'

    if_stack = 0
    for line in types.split("\n"):
        line = line.strip()
        if not line:
            continue

        if line.startswith("static"):
            continue

        if system == System.WINDOWS:
            line = line.replace("__declspec(dllexport)", "DllExport")

        result += line + "\n"

    result += "#ifndef __rrt_methods\n"

    skip = True
    skip_cplusplus = False
    if_stack = 0
    for line in header.split("\n"):
        line = line.strip()
        if not line:
            continue

        if skip:
            if line == "/* End of boilerplate cgo prologue.  */":
                skip = False
            continue

        if skip_cplusplus:
            if line == "#endif":
                if if_stack == 0:
                    skip_cplusplus = False
                else:
                    if_stack -= 1
            elif line.startswith("#if"):
                if_stack += 1
            continue
        if line == "#ifdef __cplusplus":
            skip_cplusplus = True
            continue

        if system == System.WINDOWS:
            line = line.replace("__declspec(dllexport)", "DllExport")

        result += line + "\n"

    result += "#define __rrt_methods\n#endif\n#ifdef __cplusplus\n}\n#endif\n"

    return result
