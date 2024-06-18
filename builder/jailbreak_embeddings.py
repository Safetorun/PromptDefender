import os

import time
from openai import OpenAI

client = OpenAI(api_key=os.environ.get('OPENAI_API_KEY'))
import pandas as pd
import tiktoken
import json

max_tokens = 1000


def read_and_preprocess_json(json_file):
    with open(json_file, 'r') as f:
        data = json.load(f)
    return data


def split_into_many(text, max_tokens=max_tokens):
    tokenizer = tiktoken.get_encoding("cl100k_base")
    sentences = text.split('. ')

    n_tokens = [len(tokenizer.encode(" " + sentence)) for sentence in sentences]

    chunks = []
    tokens_so_far = 0
    chunk = []

    for sentence, token in zip(sentences, n_tokens):
        if tokens_so_far + token > max_tokens:
            chunks.append(". ".join(chunk) + ".")
            chunk = []
            tokens_so_far = 0

        if token > max_tokens:
            continue

        chunk.append(sentence)
        tokens_so_far += token + 1

    return chunks


def convert_to_openai(x):
    try:
        return client.embeddings.create(input=[x], model='text-embedding-3-small').data[0].embedding
    except Exception as ex:
        print(ex)
        print("Sleeping")
        time.sleep(60)
        print("Continue")
        return convert_to_openai(x)


def process_from_file(outname):
    texts = list(map(lambda x: {"value": x}, read_and_preprocess_json("injections.json")))

    df = pd.DataFrame(texts, columns=['value'])
    df.to_csv('injections.csv')
    df.head()
    tokenizer = tiktoken.get_encoding("cl100k_base")

    shortened = []

    df['n_tokens'] = df.value.apply(lambda x: len(tokenizer.encode(x)))

    for row in df.iterrows():
        if row[1]['value'] is None:
            continue

        if row[1]['n_tokens'] > max_tokens:
            shortened += list(map(lambda x: (split_into_many(row[1]['value']))))

        else:
            shortened.append((row[1]['value']))

    df = pd.DataFrame(shortened, columns=['value'])

    df['n_tokens'] = df.value.apply(lambda x: len(tokenizer.encode(x)))

    df['embeddings'] = df.value.apply(
        lambda x: convert_to_openai(x))

    df.to_csv(outname)
    df.head()


if __name__ == "__main__":
    process_from_file("injections_embeddings.csv")
