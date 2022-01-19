import { Passport } from "./passport";
import { Course } from "./course";
import { Hobby } from "./hobby";

export class User{
	ID?:string;
	CreatedBy?:string;
	CreatedAtTime?:string
	DeletedAt?:string;
    FirstName?:string
	LastName?: string
    Address?:string
	Email?:string
	Password?:string
	Passport?:Passport
	Hobbies?:Hobby[]
    Courses?:Course[]
}