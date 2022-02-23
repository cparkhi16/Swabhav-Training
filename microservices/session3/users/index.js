const express=require('express')
const cors=require('cors')
const uuid=require('uuid')
const bodyParser=require('body-parser');
const { default: axios } = require('axios');
const user = require('./usersmodel');
const app = express();
const bcrypt = require('bcrypt');
async function hashIt(password){
//   const salt = await bcrypt.genSalt(6);
//   const hashed = await bcrypt.hash(password, salt);
//   console.log("Hashed pwd ",hashed)
const hashedPassword = bcrypt.hashSync(password, bcrypt.genSaltSync());
  return hashedPassword;
}
 function comparePassword(password,hashedPassword){
    const doesPasswordMatch = bcrypt.compareSync(password, hashedPassword)
    return doesPasswordMatch
  }
app.use(cors())
app.use(bodyParser.json())
//Insert into userdb.users values("j88160f1-8e1f-4847-a963-a40a0c035159","user4","$2b$10$us494jZvonIkJKVxQvdvjuoTaAbHK/GzO9jmgWB0U6GunaJ.I48N2");
let allUsers =[];
app.post('/api/v1/login',async (req,resp)=>{
    const {username,password}=req.body
    console.log("Login called ",username,password)
    console.log("Hashed pwd ",hashIt(password))
    user.findAll(function(err, users) {
        console.log('controller')
        if (err)
        resp.send(err);
        console.log('res', users);
        //resp.send(users);
        allUsers=users;
        //console.log("all users in db here ",allUsers)
        for (let u of allUsers ){
            console.log("user ",u)
            const validPassword =  comparePassword(password,u.password)
            console.log("password ",validPassword)
            if(u.username==username && validPassword){
                console.log("User found !! ")
                resp.send({validUser:true,userData:u})
                return
            }
        }
        resp.send({validUser:false})
      });
})


app.listen(4006,()=>{
    console.log("UserService has started on 4006")
})