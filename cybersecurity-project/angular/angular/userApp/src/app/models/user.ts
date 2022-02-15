import { Course } from "./course"
import { Hobby } from "./hobby"
import { Passport } from "./passport"

export class User{
	ID?:string;
	CreatedBy?:string;
	CreatedAt?:string
	DeletedAt?:string;
    FirstName?:string
	LastName?: string
	Email?:string
	Password?:string
	Passport?:Passport
	Courses?:Course[]
	Hobbies?:Hobby[]
	LevelBIBA?:number;
	LevelBell?:number;
	SecretAnswer?:string;
}