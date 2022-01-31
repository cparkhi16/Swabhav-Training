import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import {ObsService} from '../myservice/obs.service'
import { ActivatedRouteSnapshot, CanActivate, Router, RouterStateSnapshot, UrlTree } from '@angular/router';
@Injectable({
  providedIn: 'root'
})
export class AuthGuardService implements CanActivate  {
  isValid:any
  constructor(private obs:ObsService, public router: Router) { }
  canActivate(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot): Observable<boolean | UrlTree> | Promise<boolean | UrlTree> | boolean | UrlTree {
      console.log("Checking token ");
      this.checkValidity()
      //console.log(this.isValid)
    // if(this.isValid==false){
    //   this.router.navigate(['login'])
    //   return false;
    // }
    return true;
  }
  checkValidity(){
    console.log("Checking validity")
  let token:any=localStorage.getItem('Token');
  let decryptedToken:any =this.obs.decryptData(token)
  this.obs.validateToken(decryptedToken).subscribe((data:any)=>{
    console.log("Valid token" , data);
    this.isValid=data.IsValidToken;
    if(this.isValid==false){
      this.router.navigate(['login'])
    }
  });
  console.log("End of checkvalidity")
}
}
