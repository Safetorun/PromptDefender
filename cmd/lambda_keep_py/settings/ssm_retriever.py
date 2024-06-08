import boto3


def get_secret(secret_name: str) -> str:
    """
    Retrieve secret from AWS Secrets Manager

    :param secret_name:

    :return:
    """
    session = boto3.session.Session()
    client = session.client(service_name='ssm')
    get_secret_value_response = client.get_parameter(Name=secret_name, WithDecryption=True)
    return get_secret_value_response['Parameter']['Value']
