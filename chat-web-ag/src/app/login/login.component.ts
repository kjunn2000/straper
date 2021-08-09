import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { SharedService } from '../services/shared.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})

export class LoginComponent implements OnInit {

  username: string = '';
  password: string = '';
  error = false;
  hide = true;

  constructor(
    private router: Router,
    private service: SharedService
  ) { }

  ngOnInit(): void {
  }

  login() {
    if (this.username === '1' && this.password === '1') {
      this.error = false;
      this.router.navigate(['/home']);
      this.service.setIsLogin();
    } else {
      this.error = true;
    }

  }

}
