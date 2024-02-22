import json
import openai
import pandas as pd
import tiktoken
import time
import os

from embeddings import preprocess_text, embed_text

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




def process_from_file(outname):
    texts = read_and_preprocess_json("jailbreaks.json")

    df = pd.DataFrame(texts, columns=['name', 'value'])
    df = df[df['value'].notnull()]
    df['value'].apply(preprocess_text)
    df.to_csv('scraped.csv')

    df.head()

    df['embeddings'] = df.value.apply(lambda x: embed_text(x))

    df = df.dropna()

    df.to_csv(outname)
    df.head()


def setup_openai():
    openai.api_key = os.environ.get('OPENAI_API_KEY')


if __name__ == "__main__":
    setup_openai()
    process_from_file("scanned.csv")
