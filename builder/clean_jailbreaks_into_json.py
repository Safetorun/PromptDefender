import json

f = open("jailbreaks.json", "w")

jailbreaks = []

with open("jailbreaking.txt") as jailbreaking_file:
    jailbreaking = jailbreaking_file.readlines()

    in_main_doc = False

    in_tag = False
    tag_content = ""
    summary_content = ""

    for line in jailbreaking:
        line = line.rstrip()
        if line.startswith("<details>"):
            in_tag = True
        elif line.startswith("</details>"):
            in_tag = False
            jailbreaks.append({"name": tag_content, "value": summary_content})
            tag_content = ""
        elif line.startswith("  <summary>"):
            summary_content = line.removeprefix("  <summary>").removesuffix("</summary>")
        elif line.startswith("  <blockquote>") or line.startswith("<blockquote>"):
            continue
        elif in_main_doc and line != "\n" and not line.startswith("##") and not line.startswith("==="):
            if in_tag:
                tag_content += line.strip()
        elif "-->" in line:
            in_main_doc = True

json.dump(jailbreaks, f)
