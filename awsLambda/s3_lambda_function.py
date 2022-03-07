import json
import os
import boto3
from boto3.dynamodb.conditions import Key,Attr
from datetime import datetime
dynamodb = boto3.resource('dynamodb')
print(dynamodb)

table_cust=dynamodb.Table('Customers')
print(table_cust)

table_org=dynamodb.Table('Org')
print(table_org)
def lambda_handler(event, context):
    # TODO implement
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
    now = datetime.now()
    date_time = now.strftime("%m/%d/%Y, %H:%M:%S")
    print("now date_time =", date_time)
    table_cust.update_item(
        Key={"email":"kp@fp.com"},
        UpdateExpression= "SET uploadedat=:uploadedat",
        ExpressionAttributeValues={
            ":uploadedat":date_time
        }
        )
    return {
        'statusCode': 200,
        'body': json.dumps('Hello from Lambda!')
    }

    