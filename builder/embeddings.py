import os
import pandas as pd
import re
import numpy as np
import requests
from nltk.corpus import stopwords
import nltk
from nltk.stem import WordNetLemmatizer
from sklearn.feature_extraction.text import TfidfVectorizer
from sklearn.metrics.pairwise import cosine_similarity
import time
import openai

MODEL_ID = 'sentence-transformers/bert-base-nli-mean-tokens'
HF_TOKEN = os.getenv("HUGGINGFACE_TOKEN")

api_url = f"https://api-inference.huggingface.co/pipeline/feature-extraction/{MODEL_ID}"
headers = {"Authorization": f"Bearer {HF_TOKEN}"}

MODE = "huggingface"  # openai


def convert_to_openai(x):
    try:
        return openai.Embedding.create(input=x[0], engine='text-embedding-ada-002')['data'][0]['embedding']
    except Exception as ex:
        print(ex)
        print("Sleeping")
        time.sleep(60)
        print("Continue")
        return convert_to_openai(x)


def query(texts):
    if MODE == "openai":
        return convert_to_openai(texts)
    else:
        response = requests.post(api_url, headers=headers, json={"inputs": texts, "options": {"wait_for_model": True}})
        if response.status_code == 200:
            return response.json()
        else:
            raise Exception(f"Failed to retrieve embeddings: {response.status_code}, {response.text}")


def preprocess_text(text):
    nltk.download('stopwords')
    nltk.download('wordnet')

    text = text.lower()
    text = re.sub(r'[^\w\s]', '', text)
    stop_words = set(stopwords.words('english'))
    words = text.split()
    words = [word for word in words if word not in stop_words]

    lemmatizer = WordNetLemmatizer()
    words = [lemmatizer.lemmatize(word) for word in words]
    return ' '.join(words)


def feature_engineering(df):
    # Use TF-IDF to create features
    vectorizer = TfidfVectorizer()
    features = vectorizer.fit_transform(df['texts'])
    return pd.DataFrame(features.toarray(), columns=vectorizer.get_feature_names_out())


def embed_text(text):
    return query([preprocess_text(text)])


def get_embeddings(texts):
    outputs = query(texts)
    embeddings = pd.DataFrame(outputs)
    embeddings['texts'] = texts
    embeddings['texts'] = embeddings['texts'].apply(preprocess_text)
    return embeddings


def save_embeddings_to_csv(embeddings, filename):
    embeddings.to_csv(filename, index=False)


def load_embeddings_from_csv(filename):
    return pd.read_csv(filename)


def calculate_similarity(prompt_embeddings, df):
    embeddings = np.concatenate(df['embeddings'].to_list())
    similarities = cosine_similarity(prompt_embeddings, embeddings)
    df['similarity'] = similarities[0]
    return df


def sort_by_similarity(df):
    return df.sort_values(by='similarity', ascending=False)
