import json
import os
import boto3
from boto3.dynamodb.conditions import Key

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

    if 'apiTok' in event["queryStringParameters"]:
            if event["queryStringParameters"]["apiTok"] == os.getenv('apiTok') :
                body["msg"] = "👍Yep! it is the right api key"
            else:
                body["msg"] = "⛔Oops! wrong api key"
                status_code = 403
    else:
        body["msg"] = "😥 Looks like you forget the api key. (e.g. /v1/tag?apiTok=your_api_key)"
        status_code = 403
    
    if 'team' in event["queryStringParameters"]:
        q = event["queryStringParameters"]["team"]
        res = get_tag(q)
        body["tags"]=res["Items"]
        
    return {
            "statusCode": status_code,
            "body": json.dumps(body, ensure_ascii=False),
            "headers": {
                "Content-Type": "application/json"
            }
        }    