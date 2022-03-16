import json
import os
import boto3
from boto3.dynamodb.conditions import Key,Attr
from datetime import datetime
dynamodb = boto3.resource('dynamodb')
print(dynamodb)

table_cust=dynamodb.Table('NewUsers')
print(table_cust)


def lambda_handler(event, context):
    # TODO implement
    print("My event ",event)
    params=event['queryStringParameters']
    email=params['email']
    fname=params['fname']
    # lname=params['lname']
    # phone=params['phone']
    print(email,fname)
    #print("Is user valid ",isPresent)
    resp1=table_cust.query(KeyConditionExpression=Key('email').eq(email))
    print("Query ",resp1['Items'])
    values=resp1['Items']
    return {
        'statusCode': 200,
        'body': json.dumps('Users found')
    }

