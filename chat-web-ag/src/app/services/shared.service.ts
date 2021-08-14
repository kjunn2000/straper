import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class SharedService {

  private isLogin: boolean = false;
  constructor() { }

  getIsLogin() {
    return this.isLogin;
  }

  setIsLogin() {
    this.isLogin = true;
  }
}
