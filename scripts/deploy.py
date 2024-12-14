import os
import boto3
import requests
from dotenv import load_dotenv
import time
import json


def main():
    load_dotenv()  # take environment variables from .env.

    bucketName = 'auth-microservice'
    objectName = 'yandex_cloud.zip'

    uploadArchiveToBucket(bucketName, objectName)

    time.sleep(3)  # Sleep for 3 seconds just in case

    # use this to delete previous versions `yc serverless function version delete <version_id>`
    createNewFnVersionFromBucket(bucketName, objectName)

    # returnVersionToGoMod()


def filePutContents(fl: str, dt: str):
    flHandle = open(fl, 'w')
    flHandle.write(dt)
    flHandle.close()


def fileGetContents(fl: str) -> str:
    flHandle = open(fl, 'r')
    contents = flHandle.read()
    flHandle.close()
    return contents


def returnVersionToGoMod():
    print('returnVersionToGoMod')
    filePutContents(
        'go.mod',
        fileGetContents('go.mod').replace(
            "module github.com/gennadyterekhov/auth-microservice\n\n",
            "module github.com/gennadyterekhov/auth-microservice\n\ngo 1.23.0",
        ),
    )


def uploadArchiveToBucket(bucketName: str, objectName: str):
    print('uploadArchiveToBucket')
    session = boto3.session.Session()
    s3 = session.client(service_name='s3', endpoint_url='https://storage.yandexcloud.net')

    # upload to bucket
    s3.upload_file('artefacts/yandex_cloud.zip', bucketName, objectName)


def createNewFnVersionFromBucket(bucketName: str, objectName: str):
    print('createNewFnVersionFromBucket')
    createFunctionVersionBody = {
        "functionId": os.getenv('YA_CLOUD_FUNCTION_ID', ''),
        "runtime": "golang121",
        "entrypoint": 'cmd/server/ya_cloud.Handler',
        "resources": {
            # 128 MB
            "memory": "134217728"
        },
        "executionTimeout": "1s",
        "serviceAccountId": os.getenv('YA_CLOUD_SERVICE_ACCOUNT_ID', ''),
        "package": {
            "bucketName": bucketName,
            "objectName": objectName
        },
    }

    IAM_TOKEN = os.getenv('YA_CLOUD_IAM_TOKEN', '')
    headers = {'Authorization': f'Bearer {IAM_TOKEN}'}
    url = 'https://serverless-functions.api.cloud.yandex.net/functions/v1/versions'

    # create new fn version
    response = requests.post(url, data=json.dumps(createFunctionVersionBody), headers=headers)

    print(response)
    print('txt', response.text)
    print('reason', response.reason)


if __name__ == '__main__':
    main()
