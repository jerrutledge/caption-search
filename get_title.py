# -*- coding: utf-8 -*-

# Sample Python code for youtube.videos.list
# See instructions for running these code samples locally:
# https://developers.google.com/explorer-help/code-samples#python

import json
import os

import google_auth_oauthlib.flow
import googleapiclient.discovery
import googleapiclient.errors

input_file = "yt_ids.txt"
output_file = "yt_metadata.json"

scopes = ["https://www.googleapis.com/auth/youtube.readonly"]


def main():
    # Disable OAuthlib's HTTPS verification when running locally.
    # *DO NOT* leave this option enabled in production.
    os.environ["OAUTHLIB_INSECURE_TRANSPORT"] = "1"

    api_service_name = "youtube"
    api_version = "v3"
    client_secrets_file = "client_secret_.json"

    # Get credentials and create an API client
    flow = google_auth_oauthlib.flow.InstalledAppFlow.from_client_secrets_file(
        client_secrets_file, scopes)
    credentials = flow.run_console()
    youtube = googleapiclient.discovery.build(
        api_service_name, api_version, credentials=credentials)

    # perform requests
    responses = []

    with open(input_file) as file:
        for line in file:
            new_id = line.strip()
            request = youtube.videos().list(
                part="snippet",
                id=new_id
            )
            response = request.execute()

            responses.append(response)
            print("recieved response for id", new_id)

    with open(output_file, 'w') as file:
        json.dump(responses, file)


if __name__ == "__main__":
    main()
