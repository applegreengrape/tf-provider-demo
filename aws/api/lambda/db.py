import json
import boto3

def dict_to_item(raw):
    if isinstance(raw, dict):
        return {
            'M': {
                key: dict_to_item(value)
                for key, value in raw.items()
            }
        }
    elif isinstance(raw, list):
        return {
            'L': [dict_to_item(value) for value in raw]
        }
    elif isinstance(raw, str) or isinstance(raw, unicode):
        return {'S': raw}
    elif isinstance(raw, int):
        return {'N': str(raw)}
    elif raw is None:
        return {'NULL': True}
      
item = {}
item["id"] = "1"
item["tag"] = "abc"
item["name"] = "🗿"
item["hostname"] = "📛 hostname"
item["owner"] = "🐻 owner"
item["region"] = "📍 home"
item["meta"] = "📝 Blablablablabla, blabla bla blablablabla blabla blablablablablabla"

item2 = {}
item2["id"] = "2"
item2["tag"] = "releaser"
item2["product"] = "🦄App"
item2["version"] = "v0.0.1"
item2["relased date"] = "11/11/2020"

_item = dict_to_item(item)
_item2 = dict_to_item(item2)
dynamodb = boto3.client('dynamodb')
dynamodb.put_item(TableName='tag', Item=_item["M"])
dynamodb.put_item(TableName='tag', Item=_item2["M"])