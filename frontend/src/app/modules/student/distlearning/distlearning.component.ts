import { Component, OnInit } from '@angular/core';
import { DistlearningService } from './distlearning.service';
import { CommonModule } from '@angular/common';
import { Router } from '@angular/router';

@Component({
  selector: 'app-distlearning',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './distlearning.html',
//   styleUrl: './distlearning.css'
  
})
export class DistlearningComponent implements OnInit {

    constructor(
        private _distservice: DistlearningService,
        private _router: Router,
    )
    {
    }
    
    

    ngOnInit(): void {
        console.log('dist learn load')
    }

}
