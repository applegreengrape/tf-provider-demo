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
item["tag"] = "team_abc"
item["name"] = "some_cool_name"
item["hostname"] = "📛some_cool_hostname"
item["owner"] = "🐻some_cool_owner"
item["region"] = "📍some_cool_region"
item["meta"] = "Blablablablabla, blabla bla blablablabla blabla blablablablablabla"

_item = dict_to_item(item)
print(_item["M"])
dynamodb = boto3.client('dynamodb')
dynamodb.put_item(TableName='tag', Item=_item["M"])