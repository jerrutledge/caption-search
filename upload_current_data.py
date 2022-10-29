# -*- coding: utf-8 -*-

import json
from os import listdir
from os.path import isfile, join


directory_path = "transcripts"
json_filename = "yt_metadata.json"
output_filename = "match.tsv"


def get_files():
    onlyfiles = [f for f in listdir(
        directory_path) if isfile(join(directory_path, f))]

    return onlyfiles


def find_title(title_dict, desired_title):
    for key in title_dict:
        if desired_title in title_dict[key]:
            return key


if __name__ == "__main__":
    filenames = get_files()
    if '.DS_Store' in filenames:
        filenames.remove('.DS_Store')
    data = []
    with open(json_filename) as file:
        data = json.load(file)

    id_full_title = {}
    for entry in data:
        # try to find a matching srt file
        ep_id = entry["items"][0]["id"]
        episode_title = entry["items"][0]["snippet"]["title"]
        episode_title = episode_title.replace('.', '').replace(
            '/', '').replace("”", '').replace("“", '').replace('"', '')
        episode_title = episode_title.replace(':', "").lower()
        episode_title = episode_title.replace("episode", "ep")
        episode_title = episode_title.replace(" ep ", " ").replace(" - ", " ")
        id_full_title[ep_id] = episode_title
        print(episode_title)

    skips = []
    with open(output_filename, 'w') as file:
        for filename in filenames:
            if ".srt" not in filename:
                skips.append(filename)
                continue
            title = filename.replace(".srt", "")
            title = title.replace('.', '').lower().replace("episode", "ep")
            title = title.replace(" ep ", " ").replace(" - ", " ")
            found_title = find_title(id_full_title, title)
            if found_title == None:
                skips.append(filename)
                continue
            file.write(found_title + "\t" + filename + "\n")

    if len(skips):
        print("Had to skip:", skips)
