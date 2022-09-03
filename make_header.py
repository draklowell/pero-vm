import os

with open(f"build/{os.environ['NAME']}-{os.environ['OS']}-{os.environ['ARCH']}.h") as f:
    generated_lines = f.readlines()

with open(f"wrapper/types.h") as f:
    types_lines = f.readlines()

result = '#ifdef __cplusplus\nextern "C" {\n#endif\n'

if_stack = 0
for line in types_lines:
    line = line.strip()
    if not line:
        continue

    if line.startswith("static"):
        continue

    result += line + "\n"

result += "#ifndef __rrt_methods\n"

skip = True
skip_cplusplus = False
if_stack = 0
for line in generated_lines:
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

    result += line + "\n"

result += "#define __rrt_methods\n#endif\n#ifdef __cplusplus\n}\n#endif\n"

with open(
    f"build/{os.environ['NAME']}-{os.environ['OS']}-{os.environ['ARCH']}.h", "w"
) as f:
    f.write(result)
