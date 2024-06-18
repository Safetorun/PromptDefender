from openai import OpenAI
from scipy.spatial.distance import cosine

import numpy as np
import pandas as pd
import os

from settings.ssm_retriever import get_secret

client = OpenAI()
os.environ["OPENAI_API_KEY"] = get_secret(os.getenv("OPENAI_SECRET_NAME"))


def convert_to_openai(x: str):
    return client.embeddings.create(input=[x], model='text-embedding-3-small').data[0].embedding


def find_lowest_score(question: str):
    df = pd.read_csv('injections_embeddings.csv', index_col=0)
    df['embeddings'] = df['embeddings'].apply(eval).apply(np.array)
    df.head()

    q_embeddings = convert_to_openai(question)
    df['distances'] = df["embeddings"].apply(lambda x: cosine(q_embeddings, x))
    df.sort_values('distances', ascending=True)
    lowest_value = df.iloc[0]

    return lowest_value.distances, lowest_value.name, lowest_value.value
