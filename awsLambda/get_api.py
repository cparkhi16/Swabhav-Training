import json
import os
import boto3
from boto3.dynamodb.conditions import Key,Attr
from datetime import datetime
dynamodb = boto3.resource('dynamodb')
print(dynamodb)

table_cust=dynamodb.Table('Users')
print(table_cust)


def lambda_handler(event, context):
    # TODO implement
    print("My event ",event)
    params=event['queryStringParameters']
    email=params['email']
    fname=params['fname']
    lname=params['lname']
    phone=params['phone']
    print(email,fname,lname,phone)
    isPresent=checkIfUserExists(email)
    print("Is user valid ",isPresent)
    if isPresent==True:
        return {
        'statusCode': 403,
        'body': json.dumps('User already exists !!')
        }
    else:
        r=table_cust.put_item(Item={"email":email,"fname":fname,"lname":lname,"phone":phone})
        return {
        'statusCode': 200,
        'body': json.dumps('Hello from Lambda!'+email)
        }

def checkIfUserExists(email):
    resp3= table_cust.get_item(Key={"email":email})
    if "Item" in resp3:
        return True
    else:
        return False