import os
import requests
import pandas as pd
import time
from sklearn.metrics.pairwise import cosine_similarity
import numpy as np

MODEL_ID = "sentence-transformers/all-MiniLM-L6-v2"
HF_TOKEN = os.getenv("HUGGINGFACE_TOKEN")

api_url = f"https://api-inference.huggingface.co/pipeline/feature-extraction/{MODEL_ID}"
headers = {"Authorization": f"Bearer {HF_TOKEN}"}


def query(texts):
    response = requests.post(api_url, headers=headers, json={"inputs": texts, "options": {"wait_for_model": True}})
    return response.json()


def get_texts():
    return ["How do I get a replacement Medicare card?",
            "What is the monthly premium for Medicare Part B?",
            "How do I terminate my Medicare Part B (medical insurance)?",
            "How do I sign up for Medicare?",
            "Can I sign up for Medicare Part B if I am working and have health insurance through an employer?",
            "How do I sign up for Medicare Part B if I already have Part A?",
            "What are Medicare late enrollment penalties?",
            "What is Medicare and who can get it?",
            "How can I get help with my Medicare Part A and Part B premiums?",
            "What are the different parts of Medicare?",
            "Will my Medicare premiums be higher because of my higher income?",
            "What is TRICARE ?",
            "Should I sign up for Medicare Part B if I have Veterans' Benefits?"]


def get_embeddings(texts):
    outputs = query(texts)
    embeddings = pd.DataFrame(outputs)
    embeddings['texts'] = texts
    return embeddings


def save_embeddings_to_csv(embeddings, filename):
    embeddings.to_csv(filename, index=False)


def load_embeddings_from_csv(filename):
    return pd.read_csv(filename)


def calculate_similarity(prompt_embeddings, df):
    similarities = cosine_similarity(prompt_embeddings, df.iloc[:, :-1].values)
    df['similarity'] = similarities[0]
    return df


def sort_by_similarity(df):
    return df.sort_values(by='similarity', ascending=False)


def main():
    texts = get_texts()
    embeddings = get_embeddings(texts)
    save_embeddings_to_csv(embeddings, "embeddings.csv")

    prompt = ["Replace Medicare card"]
    prompt_embeddings = query(prompt)

    df = load_embeddings_from_csv("embeddings.csv")
    df = calculate_similarity(prompt_embeddings, df)
    df_sorted = sort_by_similarity(df)

    print(df_sorted.iloc[0]['texts'])
    save_embeddings_to_csv(df_sorted, "sorted_embeddings.csv")


if __name__ == "__main__":
    main()
