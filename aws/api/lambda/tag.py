import json
import os
import boto3
from boto3.dynamodb.conditions import Key
import re

def get_tag(team):
    dynamodb = boto3.resource('dynamodb')
    table = dynamodb.Table('tag')
    response = table.query(
        TableName='tag',
        KeyConditionExpression=Key('tag').eq('{0}'.format(team)),
    )
    return response

res = get_tag("team_abc")

def handler(event, context):
    body = {}
    status_code = 200
    
    if event["headers"]["Authorization"]:
        tok = re.findall('Bearer\s(.*)',event["headers"]["Authorization"])
        if tok[0] == os.getenv('apiTok') :
            body["msg"] = "👍Yep! it is the right api key"
            if event["queryStringParameters"]:
                if 'team' in event["queryStringParameters"]:
                    q = event["queryStringParameters"]["team"]
                    res = get_tag(q)
                    body["tags"]=res["Items"]
            else:
                body["msg"] = "👋 hello from the api endpoint"
        else:
            body["msg"] = "⛔Oops! wrong api key"
            status_code = 403
    else:
        body["msg"] = "😥 Looks like you forget to set the api key as X-Authorization header"
        status_code = 403

        
    return {
            "statusCode": status_code,
            "body": json.dumps(body, ensure_ascii=False),
            "headers": {
                "Content-Type": "application/json"
            }
        }    