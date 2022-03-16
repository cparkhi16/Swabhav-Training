import json
import boto3
import uuid
import base64
dynamodb = boto3.resource('dynamodb')
table_user=dynamodb.Table('Users')
table_ekartitem=dynamodb.Table('EkartItems')
table_ekartcart=dynamodb.Table('EkartCart')
getMethod='GET'
postMethod='POST'
registerPath='/register'
loginPath='/login'
createProductPath='/product'
orderPath='/product/order'
deleteMethod='DELETE'
putMethod='PUT'
from boto3.dynamodb.conditions import Key,Attr
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
        print("get my ordered products called !")
        return {
                'statusCode': 200,
                'body': json.dumps(res)
            }

def getOrderedProducts(event):
    params=event['queryStringParameters']
    email=params['email']
    resp1=table_ekartcart.scan(FilterExpression=Attr("email").eq(email),ProjectionExpression="prod_name,total_amount,quantity")
    print("--== ",resp1['Items'])
    return resp1['Items']
def deleteProduct(event):
    params=event['queryStringParameters']
    email=params['email']
    item_id=params['itemno']
    resp1=table_ekartitem.query(KeyConditionExpression=Key('item_id').eq(item_id))
    if resp1['Items'][0]['supplier_email']==email:
        re=table_ekartitem.delete_item(Key={"item_id":item_id,"supplier_email":email})
        #r=table_ekartcart.delete_item(Key={"item_id":item_id})
        return True,"Deleted successfully"
    else:
        return False,"Not allowed to delete other's products"
        
def updateProduct(event):
    json_data=json.loads(event['body'])
    item_id=json_data['data']['itemno']
    quantity=json_data['data']['quantity']
    email=json_data['data']['email']
    price=json_data['data']['price']
    prod_name=json_data['data']['productname']
    resp3= table_ekartitem.get_item(Key={"item_id":item_id,"supplier_email":email})
    if "Item" not in resp3:
        return False,"Item does not exist with the given item id and email"
    else:
        table_ekartitem.update_item(
                Key={"item_id":item_id,"supplier_email":email},
                UpdateExpression= "SET quantity_available=:quantity_available,price=:price,prod_name=:prod_name",
                ExpressionAttributeValues={
                    ":quantity_available":quantity,
                    ":price":price,
                    ":prod_name":prod_name
                }
                )
        resp1=table_ekartcart.query(KeyConditionExpression=Key('item_id').eq(item_id))
        print("====== ",resp1['Items'])
        for item in resp1['Items']:
            print(item['prod_name'],item['email'])
            table_ekartcart.update_item(
                Key={"item_id":item_id,"email":item['email']},
                UpdateExpression= "SET prod_name=:prod_name",
                ExpressionAttributeValues={
                ":prod_name":prod_name,
                }
            )
        return True,"Product updated"
        # table_ekartcart.update_item(
        #         Key={"item_id":item_id,"email":email},
        #         UpdateExpression= "SET quantity=:quantity,total_amount=:total_amount",
        #         ExpressionAttributeValues={
        #         ":quantity":str(int(resp3['Item']['quantity'])+int(quanRequired)),
        #         ":total_amount":str(int(resp3['Item']['total_amount'])+int(total_amount))
        #         }
        #     )
    
    
def placeOrder(event):
    json_data=json.loads(event['body'])
    item_id=json_data['data']['itemno']
    quanRequired=json_data['data']['quantity']
    email=json_data['data']['email']
    isPresent=checkIfUserExists(email)
    if isPresent==False:
        return False,"No user exists with this email-id.Please register yourself"
    resp1=table_ekartitem.query(KeyConditionExpression=Key('item_id').eq(item_id))
    print("Query ",resp1['Items'])
    if resp1['Items'][0]['supplier_email']==email:
        print("Not allowed to buy your own item")
        return False,"Not allowed to buy your own item"
    elif int(resp1['Items'][0]['quantity_available'])<int(quanRequired):
        return False,"Sorry,required quantity is more than available"
    else:
        resp3= table_ekartcart.get_item(Key={"item_id":item_id,"email":email})
        print("-= =-",resp3)
        total_amount=str(int(resp1['Items'][0]['price'])*int(quanRequired))
        if "Item" in resp3:
            table_ekartcart.update_item(
                Key={"item_id":item_id,"email":email},
                UpdateExpression= "SET quantity=:quantity,total_amount=:total_amount",
                ExpressionAttributeValues={
                ":quantity":str(int(resp3['Item']['quantity'])+int(quanRequired)),
                ":total_amount":str(int(resp3['Item']['total_amount'])+int(total_amount))
                }
            )
        else:
            r=table_ekartcart.put_item(Item={"item_id":item_id,"email":email,"order_status":"ordered","prod_name":resp1['Items'][0]['prod_name'],"quantity":quanRequired,"total_amount":total_amount})
        print(quanRequired,item_id)
        table_ekartitem.update_item(
        Key={"item_id":item_id,"supplier_email":resp1['Items'][0]['supplier_email']},
        UpdateExpression= "SET quantity_available=:quantity_available",
        ExpressionAttributeValues={
            ":quantity_available":str(int(resp1['Items'][0]['quantity_available'])-int(quanRequired))
        }
        )
        return True, "Order placed"
    
def getProducts(event):
    params=event['queryStringParameters']
    isPersonal=params['personal']
    email=params['email']
    isUserPresent=checkIfUserExists(email)
    if isUserPresent==False:
        return False
    if isPersonal=="true":
        resp= table_ekartitem.scan(FilterExpression=Attr("supplier_email").eq(email),ProjectionExpression="prod_name,price,quantity_available")
        print(">>",resp["Items"])
    else:
        resp= table_ekartitem.scan(FilterExpression=Attr("supplier_email").ne(email),ProjectionExpression="prod_name,price,quantity_available")
        print(">>",resp["Items"])
    return resp["Items"]
       
def createProduct(event):
    json_data=json.loads(event['body'])
    supplier_email=json_data['data']['email']
    isUserPresent=checkIfUserExists(supplier_email)
    if isUserPresent==False:
        return False
    else:
        item_id= str(uuid.uuid4())
        price=json_data['data']['price']
        prod_name=json_data['data']['productname']
        quantity_available=json_data['data']['quanAvail']
        r=table_ekartitem.put_item(Item={"item_id":item_id,"supplier_email":supplier_email,"price":price,"prod_name":prod_name,"quantity_available":quantity_available})
        return True
    # print(r)
    
def registerUser(event):
    json_data=json.loads(event['body'])
    print(json_data)
    email=json_data['data']['email']
    isPresent=checkIfUserExists(email)
    if isPresent:
        return False
    else:   
        fname=json_data['data']['fname']
        lname=json_data['data']['lname']
        pwd=json_data['data']['pwd']
        pwd_string_bytes = pwd.encode("ascii")
        base64_bytes = base64.b64encode(pwd_string_bytes)
        base64_string = base64_bytes.decode("ascii")
        print(email,fname,lname,pwd,base64_string)
        r=table_user.put_item(Item={"email":email,"fname":fname,"lname":lname,"pwd":base64_string})
        print(r)
        return True

def checkIfUserExists(email):
    resp3= table_user.get_item(Key={"email":email})
    if "Item" in resp3:
        return True
    else:
        return False

def checkUserCredentials(event):
    json_data=json.loads(event['body'])
    print(json_data)
    email=json_data['data']['email']
    isUserPresent=checkIfUserExists(email)
    if isUserPresent==False:
        return False
    else:
        pwd=json_data['data']['pwd']
        print("User entered pwd ",pwd)
        resp3= table_user.get_item(Key={"email":email})
        print("User details --+ ",resp3['Item'])
        db_user_pwd=resp3['Item']['pwd']
        base64_bytes = db_user_pwd.encode("ascii")
        sample_string_bytes = base64.b64decode(base64_bytes)
        sample_string = sample_string_bytes.decode("ascii")
        if sample_string==pwd:
            return True
        else:
            return False
        
    
    