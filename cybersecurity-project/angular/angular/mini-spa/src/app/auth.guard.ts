import { Injectable } from "@angular/core";
import { ActivatedRouteSnapshot, CanActivate, Router, RouterStateSnapshot, UrlTree } from "@angular/router";
import { Observable } from "rxjs";


@Injectable()
export class CanActivateTeam implements CanActivate {
  constructor(private router:Router) {}

  canActivate(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot
  ): Observable<boolean|UrlTree>|Promise<boolean|UrlTree>|boolean|UrlTree {
    //return this.permissions.canActivate(this.currentUser, route.params['id']);
    if(parseInt(route.params['id'])%3==0){
        this.router.navigate([''])
        return true;
    }
    this.router.navigate(['notfound'])
    return false;
  }
}