import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, CanActivate, Router, RouterStateSnapshot, UrlTree } from '@angular/router';
import { Observable } from 'rxjs';
import { ApiService } from '../api/api.service';

@Injectable({
  providedIn: 'root'
})
export class AuthGuard implements CanActivate {
  constructor(private api:ApiService, public router: Router) { }

  canActivate(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot): Observable<boolean | UrlTree> | Promise<boolean | UrlTree> | boolean | UrlTree {
      console.log("In auth service");
    let token=this.api.decryptData(localStorage.getItem('token')!);
    if(!token){
      this.router.navigate(['login']);
      return false;
    }
    this.api.checkToken(token).subscribe((data)=>{
      console.log(data);
      
      return true;
    },(error)=>{
      console.log(error)
      this.router.navigate(['login']);
      return false;
    });
    return true;
  }
  
}
