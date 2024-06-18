import ast
import csv
import os

import math
from openai import OpenAI

from settings.ssm_retriever import get_secret

os.environ["OPENAI_API_KEY"] = get_secret(os.getenv("OPENAI_SECRET_NAME"))
client = OpenAI()


def cosine_similarity(v1, v2):
    sumxx, sumxy, sumyy = 0, 0, 0
    for i in range(len(v1)):
        x = v1[i];
        y = v2[i]
        sumxx += x * x
        sumyy += y * y
        sumxy += x * y
    return sumxy / math.sqrt(sumxx * sumyy)


def cosine_distance(v1, v2):
    return 1 - cosine_similarity(v1, v2)


def convert_to_openai(x: str):
    return client.embeddings.create(input=[x], model='text-embedding-3-small').data[0].embedding


def find_lowest_score(question: str):
    q_embeddings = convert_to_openai(question)

    with open('injections_embeddings.csv') as f:
        embeddings_data = csv.DictReader(f)
        distances = []
        for row in embeddings_data:
            embeddings_list = ast.literal_eval(row['embeddings'])
            distances.append(
                {'distances': cosine_distance(q_embeddings, embeddings_list),
                 'value': row['value']
                 }
            )

        lowest_value = sorted(distances, key=lambda x: x['distances'])[0]

    print("Lowest value: ", lowest_value['value'])
    return float(lowest_value['distances'])
