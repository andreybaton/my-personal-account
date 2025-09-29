import { Component, OnInit } from '@angular/core';
import { ProfileService } from './profile.service';
import { CommonModule } from '@angular/common';
import { Router } from '@angular/router';

@Component({
  selector: 'app-distlearning',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './profile.component.html',
//   styleUrl: './distlearning.css'
  
})
export class ProfileComponent implements OnInit {

    constructor(
        private _profservice: ProfileService,
        private _router: Router,
    )
    {
    }
    
    

    ngOnInit(): void {
        console.log('dist learn load')
    }

}
