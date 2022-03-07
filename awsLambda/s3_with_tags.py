import json
import os
import boto3
from boto3.dynamodb.conditions import Key,Attr
from datetime import datetime
dynamodb = boto3.resource('dynamodb')
print(dynamodb)

table_cust=dynamodb.Table('MediaDetails')
print(table_cust)


def lambda_handler(event, context):
    # TODO implement
    print("My event ",event['Records'][0]['s3']['object']['key'])
    print("Hello chinmay I am called ")
    # resp= table_cust.scan(FilterExpression=Attr("email").eq("cparkhi16"),ProjectionExpression="fname")
    # print("Data ",resp['Items'])
    
    # # resp1=table_cust.query(KeyConditionExpression=Key('email').eq("cparkhi16"))
    # # print("Query ",resp1['Items'])
    
    # # resp2=table_org.query(KeyConditionExpression=Key('email').eq("abc@fp.com") & Key('address').eq("ABC"))
    # # print("Querying Org table ",resp2['Items'])
    
    # resp3= table_cust.get_item(Key={"email":"cparkhi16"})
    # print("get item ",resp3)
    # if "Item" in resp3:
    #     e=resp3['Item']['email']
    #     print("Found !! ",e)
    # else:
    #     print("Not found")
    
    # resp4= table_org.get_item(Key={"email":"abc@fp.com","address":"ABC"})
    # print("get item in org ",resp4)
    # if "Item" in resp4:
    #     en=resp4['Item']['email']
    #     print("Found !! ",en)
    # else:
    #     print("Not found")
    
    
    # # r=table_cust.put_item(Item={"email":"sg@fp.com","fname":"SG"})
    # # print(r)
    
    # re=table_cust.delete_item(Key={"email":"sg@fp.com"})
    # print(re)
    s3 = boto3.client('s3')
    # bucket = 'demochinmaybucket'
    # key = 'test/bluescreen.jpg'
    # response = s3.get_object_tagging(
    #         Bucket=bucket,
    #          Key=key,
    #      )
    # tag_set = response.get("TagSet")
    # print("Tag set =>   -- ",tag_set)
    response = s3.get_object_tagging(
    Bucket='demochinmaybucket',
    Key=event['Records'][0]['s3']['object']['key'],
    )
    e=response.get("TagSet")
    tagemail=e[0]["Value"]
    print(tagemail)
    resp3= table_cust.get_item(Key={"email":tagemail})
    now = datetime.now()
    date_time = now.strftime("%m/%d/%Y, %H:%M:%S")
    
    if "Item" in resp3:
        e=resp3['Item']['email']
        print("Found !! ",e)
        table_cust.update_item(
        Key={"email":e},
        UpdateExpression= "SET uploadedat=:uploadedat,objectname=:objectname",
        ExpressionAttributeValues={
            ":uploadedat":date_time,
            ":objectname":event['Records'][0]['s3']['object']['key']
        }
        )
    else:
        r=table_cust.put_item(Item={"email":tagemail,"uploadedat":date_time,"objectname":event['Records'][0]['s3']['object']['key']})
        print("Not found")
    
    print("now date_time =", date_time)
    
    return {
        'statusCode': 200,
        'body': json.dumps('Hello from Lambda!')
    }

    