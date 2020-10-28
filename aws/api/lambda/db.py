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
item["id"] = "3"
item["tag"] = "dev"
item["name"] = "👩🏻‍💻👩🏻‍💻👩🏻‍💻"
item["ami"] = "ami-005e54dee72cc1d00"
item["instance_type"] = "t2.micro"
item["db_allocated_storage"] = "20"
item["db_instance_class"] = "db.t2.micro"


data = dict_to_item(item)
dynamodb = boto3.client('dynamodb')
dynamodb.put_item(TableName='tag', Item=data["M"])
