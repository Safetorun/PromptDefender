import hashlib
from typing import Dict, Optional, Callable


def hash_string(s):
    hash_object = hashlib.sha256(s.encode())
    hex_dig = hash_object.hexdigest()
    return hex_dig


PrepareFunction = Callable[[str], str]
RetrieveFunction = Callable[[str], Dict]
StoreFunction = Callable[[str, dict], None]


def retrieve_item_if_exists(key: str,
                            retrieve_function: RetrieveFunction,
                            prepare_function: PrepareFunction = hash_string) -> \
        Optional[Dict]:
    """
    Retrieve an item from cache if it exists already

    :param key: key to retrieve - this will be "prepared" (i.e. hashed) before being used

    :return: the item if it exists, otherwise None
    """
    prepared_key = prepare_function(key)
    return retrieve_function(prepared_key)


def store_item(key: str, item: any,
               store_function: StoreFunction,
               prepare_function: PrepareFunction = hash_string
               ):
    """
    Store an item in cache

    :param key: key to store - this will be "prepared" (i.e. hashed) before being used
    :param item: the item to store
    """
    prepared_key = prepare_function(key)
    store_function(prepared_key, item)
