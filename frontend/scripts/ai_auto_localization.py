# /// script
# requires-python = ">=3.12"
# dependencies = [
#     "dotenv",
#     "openai",
# ]
# ///

import asyncio
import os
import json
from openai import AsyncOpenAI
from dotenv import load_dotenv
from pydantic import BaseModel, Field
from typing import Annotated


class Item(BaseModel):
    key: Annotated[
        str, Field(description="The key part in the JSON file")
    ]  # The key part in the JSON file
    value: Annotated[
        str, Field(description="The value part in the JSON file")
    ]  # The value part in the JSON file


class Response(BaseModel):
    items: list[Item]


def get_prompt(src_json: str, target_lang: str) -> str:
    return f"""
You are a localization expert, translating the provided localization JSON file into the corresponding language.
That the key names in the JSON file do not need to be translated; only the value parts should be translated.
Please translate the following JSON file into {target_lang}:
---
{src_json}
"""


async def translate_file(
    client: AsyncOpenAI, src_json: str, target_lang: str
) -> str:
    response = await client.beta.chat.completions.parse(
        model="gpt-4.1-nano",
        messages=[
            {
                "role": "user",
                "content": get_prompt(src_json, target_lang),
            },
        ],
        response_format=Response,
    )

    if response and response.choices:
        try:
            parse_response = response.choices[0].message.parsed

            # dumps json string with utf-8 encoding
            return json.dumps(
                {item.key: item.value for item in parse_response.items},
                ensure_ascii=False,
                indent=2,
            )
        except json.JSONDecodeError:
            print("Response: ", response)
            raise Exception("Response is not valid JSON.")
    else:
        raise Exception("Failed to get a valid response from the API.")


async def main():
    current_dir = os.path.dirname(os.path.abspath(__file__))
    locales_dir = os.path.join(current_dir, "..", "src", "i18n", "locales")

    load_dotenv(os.path.join(current_dir, ".env"))

    # Check if the OpenAI API key is set
    if not os.getenv("OPENAI_API_KEY"):
        print("Please set the OPENAI_API_KEY environment variable.")
        exit(1)

    client = AsyncOpenAI(
        api_key=os.getenv("OPENAI_API_KEY"),
        base_url=os.getenv("OPENAI_BASE_URL", "https://api.openai.com/v1"),
    )

    source_locale_file = os.path.join(locales_dir, "zh.json")
    if not os.path.exists(source_locale_file):
        print(f"Source locale file {source_locale_file} does not exist.")
        exit(1)

    with open(source_locale_file, "r", encoding="utf-8") as f:
        source_locale = f.read()

        # check if the source locale is valid JSON
        try:
            json.loads(source_locale)
        except json.JSONDecodeError as e:
            print(f"Source locale file is not valid JSON: {e}")
            exit(1)

    print("Files to be translated:")
    tasks = []
    files_to_process = []
    for file in os.listdir(locales_dir):
        if file.endswith(".json") and file != "zh.json":
            # print(f" - {file}")
            files_to_process.append(file)
            tasks.append(
                translate_file(client, source_locale, os.path.basename(file))
            )

    if tasks:
        try:
            results = await asyncio.gather(*tasks, return_exceptions=True)
        except Exception as e:
            print(f"Error translating files: {e}")
            results = [e] * len(tasks)

        for file, result in zip(files_to_process, results):
            if isinstance(result, Exception):
                print(f"Error translating {file}: {result}")
                continue
            # write the translated content to the file
            with open(
                os.path.join(locales_dir, file), "w", encoding="utf-8"
            ) as f:
                f.write(result)
                print(f"- {file}")
    print("Translation completed.")


if __name__ == "__main__":
    asyncio.run(main())

    # print(json.dumps(Response.model_json_schema(mode="validation"), indent=2))
