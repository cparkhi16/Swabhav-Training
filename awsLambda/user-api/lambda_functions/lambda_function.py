import json
from test import *

getMethod='GET'
postMethod='POST'
registerPath='/register'
loginPath='/login'
createProductPath='/product'
orderPath='/product/order'
deleteMethod='DELETE'
putMethod='PUT'

def lambda_handler(event, context):
    # TODO implement
    httpMethod= event['httpMethod']
    path=event['path']
    if httpMethod==postMethod and path==registerPath:
        res=registerUser(event)
        print("Registered user ...")
        if res==True:
            return {
            'statusCode': 200,
            'body': json.dumps('User registered successfully !')
            }
        else:
            return {
            'statusCode': 403,
            'body': json.dumps('User already exists !')
            }
    elif httpMethod==postMethod and path==loginPath:
        print("Login called !!")
        isUserValid=checkUserCredentials(event)
        if isUserValid==True:
            return {
            'statusCode': 200,
            'body': json.dumps('Login successfull !!')
            }
        else:
            return {
            'statusCode': 403,
            'body': json.dumps('Invalid credentials !!')
            }
    elif httpMethod==postMethod and path==createProductPath:
        res=createProduct(event)
        print("Create product path called !")
        if res==True:
            return {
                'statusCode': 200,
                'body': json.dumps('Created product !!')
            }
        else:
            return {
                'statusCode': 403,
                'body': json.dumps('User does not exist.Please register and then add product !!')
            }
    elif httpMethod==putMethod and path==createProductPath:
        isProductUpdated,res=updateProduct(event)
        print("Update product path called !")
        if isProductUpdated:
            return {
                'statusCode': 200,
                'body': json.dumps('Updated product !!')
            }
        else:
            return {
                'statusCode': 403,
                'body': json.dumps(res)
            }
            
    elif httpMethod==deleteMethod and path==createProductPath:
        print("Delete product called !")
        isProductDeleted,res=deleteProduct(event)
        if isProductDeleted:
            return {
                'statusCode': 200,
                'body': json.dumps('Deleted product !!')
            }
        else:
            return {
                'statusCode': 403,
                'body': json.dumps(res)
            }
    elif httpMethod==getMethod and path==createProductPath:
        p=getProducts(event)
        print("get product  called !",p)
        if p!=False:
            p=json.dumps(p)
            return {
                'statusCode': 200,
                'body': p
            }
        else:
             return {
                'statusCode': 403,
                'body': json.dumps('No user found with this email id')
            }
    elif httpMethod==postMethod and path==orderPath:
        print("create order called !")
        isTransactionCompleted,res=placeOrder(event)
        if isTransactionCompleted:
            return {
            'statusCode': 200,
            'body': json.dumps('Order placed !!')
            }
        else:
            return {
            'statusCode': 403,
            'body': json.dumps(res)
        }
    elif httpMethod==getMethod and path==orderPath:
        # res=createProduct(event)
        res=getOrderedProducts(event)
        print("get my ordered products called !",res)
        if res!=False:
            return {
                'statusCode': 200,
                'body': json.dumps(res)
            }
        else:
            return {
                'statusCode': 403,
                'body': json.dumps('User does not exist with given mail id')
            }
            
    
